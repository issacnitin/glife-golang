package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"../../common"
	"../../common/mongodb"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
)

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

	var profileInDB User
	err = mongodb.Profile.FindOne(context.TODO(), filter).Decode(&profileInDB)

	if err == nil {
		http.Error(w, "Profile already exist", 500)
	}

	var result User
	result.Country = ""
	result.Email = resp2.Email
	result.Name = resp2.Name
	result.Phone = ""
	result.ProfileId = GenerateProfileId()
	result.Username = _GenerateUsername()

	insertResult, err := mongodb.Profile.InsertOne(context.TODO(), result)

	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	fmt.Println("Inserted document from google auth: ", insertResult)

	claims := jwt.MapClaims{
		"profileid": result.ProfileId,
	}

	_, token, err := tokenAuth.Encode(claims)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var response common.Token
	response.Token = token
	render.JSON(w, r, response)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var req User

	json.Unmarshal(b, &req)
	req.ProfileId = GenerateProfileId()

	// BIG TODO: Hash Password
	// TODO: Assuming single email, that need not be the case, user can have multiple emails linked to same account
	// For example, registration with a non google email and trying to register later with a google email
	insertResult, err := mongodb.Profile.InsertOne(context.TODO(), req)
	if err != nil {
		http.Error(w, "Registration failed, MongoDB unavailable at the moment", 500)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	claims := jwt.MapClaims{
		"profileid": req.ProfileId,
	}

	_, token, err := tokenAuth.Encode(claims)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("%s", token)

	var response common.Token
	response.Token = token
	render.JSON(w, r, response)
}
