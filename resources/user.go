package resources

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

type omit *struct{}

type User struct {
	ID        		bson.ObjectId 	`json:"id" bson:"_id,omitempty"`
	Name      		string    		`json:"name" bson:"name"`
	Password 		string			`json:"password" bson:"-"`
	HashedPassword	[]byte			`json:"hashedPassword" bson:"password"`
}

type UserAlias User

type PublicUser struct {
	*UserAlias
	Password omit `json:"password,omitempty"`
}

func (user User) prepareMarshal() *PublicUser {
	return &PublicUser{
		UserAlias: (*UserAlias)(&user),
	}
}

func (user User) MarshalJSON() ([]byte, error) {
	return json.Marshal(user.prepareMarshal())
}

type Users []User

func (users Users) prepareMarshal() []*PublicUser {
	publicUsers := make([]*PublicUser, len(users))
	for i := range users {
		publicUsers[i] = users[i].prepareMarshal()
	}
	return publicUsers
}

func (users Users) MarshalJSON() ([]byte, error) {
	return json.Marshal(users.prepareMarshal())
}