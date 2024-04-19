package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost         = 12
	minFirstNameLen    = 2
	minLastNameLen     = 2
	minPasswordNameLen = 7
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

// user validation
func (params CreateUserParams) Validate() map[string]string {
	var err = map[string]string{}
	if len(params.FirstName) < minFirstNameLen {
		err["firstName"] = fmt.Sprintf("firstName length should be atleast %d characters", minFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen {
		err["lastName"] = fmt.Sprintf("lastName length should be atleast %d characters", minLastNameLen)
	}
	if len(params.Password) < minPasswordNameLen {
		err["password"] = fmt.Sprintf("password length should be atleast %d characters", minPasswordNameLen)
	}
	if !isEmailValid(params.Email) {
		err["email"] = fmt.Sprintf("email is invalid")
	}
	return err
}

func isEmailValid(email string) bool {
	emailRegexp := regexp.MustCompile(`^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$`)
	return emailRegexp.MatchString(email)
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}
