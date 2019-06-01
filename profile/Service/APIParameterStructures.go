package profile

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetUserRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
}

type GetUserResponse struct {
	ProfileId primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Phone     string             `json:"phone"`
	Country   string             `json:"country"`
}

type RegisterUserRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
}

type RegisterUserResponse struct {
	ProfileId primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Phone     string             `json:"phone"`
	Country   string             `json:"country"`
}
