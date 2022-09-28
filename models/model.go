package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Class       string             `json:"class,omitempty" bson:"class,omitempty"`
	DateOfBirth string             `json:"dateOfBirth,omitempty" bson:"dateOfBirth,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	FullName    string             `json:"fullName,omitempty" bson:"fullName,omitempty"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	PhoneNumber float64            `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	StudentID   int32              `json:"studentID,omitempty" bson:"studentID,omitempty"`
}
