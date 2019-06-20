package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"../../common/mongodb"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tokenAuth *jwtauth.JWTAuth

type Claims struct {
	jwt.Claims
	Email string
}

func (m *Claims) Valid() error {
	return nil
}

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("dasjkhfiuadufasdfasf832742389r923rc325235c7n6235"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)

	// Protected routes
	router.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)
		r.Get("/api/v1/profile/search/{searchstring}", FindUser)
		r.Get("/api/v1/profile/getuser", GetUser)
	})

	// Public routes
	router.Group(func(r chi.Router) {
		r.Post("/api/v1/profile/register", RegisterUser)
		r.Post("/api/v1/profile/login", LoginUser)
		r.Post("/api/v1/profile/registerWithGoogle", RegisterUserWithGoogle)
		r.Post("/api/v1/profile/loginWithGoogle", LoginUserWithGoogle)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome anonymous"))
		})
	})

	return router
}

func NewProfileId() primitive.ObjectID {
	return primitive.NewObjectID()
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req LoginUserRequest
	json.Unmarshal(b, &req)

	filter := bson.D{
		{"email", req.Email},
	}

	var profileInDB RegisterUserRequest

	err = mongodb.Profile.FindOne(context.TODO(), filter).Decode(&profileInDB)

	// BIG TODO: Use JWT Token - this is hackable and unsecure
	if err != nil {
		http.Error(w, "Login failed, user not found", http.StatusBadRequest)
		return
	}

	if profileInDB.Password != req.Password {
		http.Error(w, "Login failed due to password mismatch", http.StatusBadRequest)
		return
	}

	claim := jwt.MapClaims{"email": req.Email}
	jwtauth.SetExpiryIn(claim, (1 * time.Minute))

	_, token, err := tokenAuth.Encode(claim)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", token)

	render.JSON(w, r, token)
}

func LoginUserWithGoogle(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var req LoginUserWithGoogleRequest

	json.Unmarshal(b, &req)

	// Verify google auth token is valid for our client
	resp, err := http.Get(fmt.Sprintf("https://oauth2.googleapis.com/tokeninfo?id_token=%s", req.Token))

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var resp2 GoogleAuthVerification
	json.Unmarshal(body, &resp2)

	fmt.Printf("%s", resp2.Aud)
	if resp2.Aud != "249369235819-11cfia1ht584n1kmk6gh6kbba8ab429u.apps.googleusercontent.com" || resp2.Email != req.Email {
		http.Error(w, "Authentication failure because of, you know, reasons", 500)
	}

	filter := bson.D{
		{"email", resp2.Email},
	}

	var profileInDB GetUserResponse
	err = mongodb.Profile.FindOne(context.TODO(), filter).Decode(&profileInDB)

	if err != nil {
		http.Error(w, "Profile donot exist", 500)
	}

	// BIG TODO: Use JWT Token - this is hackable just by tampering response, and unsecure

	claims := &Claims{
		Email: resp2.Email,
	}

	_, token, err := tokenAuth.Encode(claims)

	if err != nil {
		log.Fatal(err)
	}

	render.JSON(w, r, token)
}

func RegisterUserWithGoogle(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var req RegisterUserWithGoogleRequest

	json.Unmarshal(b, &req)

	// Verify google auth token is valid for our client
	resp, err := http.Get(fmt.Sprintf("https://oauth2.googleapis.com/tokeninfo?id_token=%s", req.Token))

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var resp2 GoogleAuthVerification
	json.Unmarshal(body, &resp2)

	if resp2.Aud != "249369235819-11cfia1ht584n1kmk6gh6kbba8ab429u.apps.googleusercontent.com" {
		http.Error(w, err.Error(), 500)
	}

	// Check if user exist in database
	filter := bson.D{
		{"email", resp2.Email},
	}

	var profileInDB GetUserResponse
	err = mongodb.Profile.FindOne(context.TODO(), filter).Decode(&profileInDB)

	if err == nil {
		http.Error(w, "Profile already exist", 500)
	}

	var result RegisterUserResponse
	result.Country = ""
	result.Email = resp2.Email
	result.Name = resp2.Name
	result.Phone = ""
	result.ProfileId = primitive.NewObjectID()

	insertResult, err := mongodb.Profile.InsertOne(context.TODO(), result)

	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	fmt.Println("Inserted document from google auth: ", insertResult)

	claims := &Claims{
		Email: result.Email,
	}

	_, token, err := tokenAuth.Encode(claims)

	if err != nil {
		log.Fatal(err)
	}

	render.JSON(w, r, token)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var req RegisterUserRequest

	json.Unmarshal(b, &req)

	var result RegisterUserResponse
	result.Country = req.Country
	result.Email = req.Email
	result.Name = req.Name
	result.Phone = req.Phone
	result.ProfileId = primitive.NewObjectID()

	// BIG TODO: Hash Password
	// TODO: Assuming single email, that need not be the case, user can have multiple emails linked to same account
	// For example, registration with a non google email and trying to register later with a google email
	insertResult, err := mongodb.Profile.InsertOne(context.TODO(), req)
	if err != nil {
		http.Error(w, "Registration failed, MongoDB unavailable at the moment", 500)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	claims := &Claims{
		Email: result.Email,
	}

	_, token, err := tokenAuth.Encode(claims)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", token)

	render.JSON(w, r, token)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	token, claims, err2 := jwtauth.FromContext(r.Context())

	if err2 != nil {
		fmt.Printf("%s", err2.Error())
		return
	}

	fmt.Printf("%s", token)
	fmt.Printf("%s", claims)
	// If struct not initialzed, inner variables don't exist
	profileId := fmt.Sprintf("%s", "nitin.i.joy@gmail.com")

	filter := bson.D{{
		"_id",
		bson.D{{
			"$in",
			bson.A{profileId},
		}},
	}}

	var result GetUserResponse

	err := mongodb.Profile.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusGone)
	}

	render.JSON(w, r, result)
}

func FindUser(w http.ResponseWriter, r *http.Request) {

	fmt.Sprintf("FindUser called")
	ss := fmt.Sprintf("%s", chi.URLParam(r, "searchstring"))
	ss = fmt.Sprintf(".*%s.*", ss)

	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"name", primitive.Regex{ss, ""}}},
			bson.D{{"country", ss}},
			bson.D{{"phone", ss}},
		}}}
	findOptions := options.Find()

	var result []GetUserResponse

	cur, err := mongodb.Profile.Find(context.TODO(), filter, findOptions)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	var x GetUserResponse
	for cur.Next(context.TODO()) {
		err := cur.Decode(&x)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Printf("%s", x)
		result = append(result, x)
	}

	render.JSON(w, r, result)

}
