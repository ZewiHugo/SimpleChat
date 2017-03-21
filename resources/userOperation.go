package resources

import (
	"log"
	"SimpleChat/mongodb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	index := mgo.Index{
		Key: []string{"name"},
		Unique: true,
	}

	_, collection := getUsersCollection()
	err := collection.EnsureIndex(index)
	if err != nil {
		log.Fatalf("Can not ensure index on user collection, error: %s", err.Error())
	}
}

func getUsersCollection() (*mgo.Session, *mgo.Collection) {
	session := mongodb.MongoDBSession.Copy()
	return session, session.DB("SimpleChat").C("users")
}

func CreateUser(user *User) error {
	session, c := getUsersCollection()
	defer session.Close()

	bytePassword := []byte(user.Password)
	hasedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error encrypting data")
		return err
	}
	user.HashedPassword = hasedPassword

	user.ID = bson.NewObjectId()
	err = c.Insert(&user)
	if err != nil {
		log.Printf("Error creating data")
		return  err
	}
	return err
}

func GetUserByID(id string) (*User, error) {
	objectID, err := mongodb.ObjectIDFromHexString(id)
	if err != nil {
		log.Printf("Error transfering objectID to hex string!")
		return nil, err
	}

	return GetOneUser(bson.M{"_id": objectID})
}

func GetOneUser(query interface{}) (*User, error) {
	session, c := getUsersCollection()
	defer session.Close()

	var user User
	err := c.Find(query).One(&user)
	return &user, err
}

func GetAllUsers() (*Users, error) {
	session, c := getUsersCollection()
	defer session.Close()

	var users Users
	err := c.Find(nil).Limit(100).All(&users)
	return &users, err
}

func DeleteUserByID(id string) error {
	session, c := getUsersCollection()
	defer session.Close()

	objectID, err := mongodb.ObjectIDFromHexString(id)
	if err != nil {
		log.Printf("Error transfering objectID to hex string!")
		return err
	}

	err = c.RemoveId(objectID)
	return err
}

func DeleteAllUser() error {
	session, c := getUsersCollection()
	defer session.Close()

	_, err := c.RemoveAll(bson.M{})
	return err
}

func VerifyUser(user *User) error {
	userFromDB, err := GetOneUser(bson.M{"name": user.Name})
	if err != nil {
		log.Printf("error searching database")
		return err
	}

	err = bcrypt.CompareHashAndPassword(userFromDB.HashedPassword, []byte(user.Password))
	return err
}