package mongodb

import (
    "errors"
    "gopkg.in/mgo.v2/bson"
)

func ObjectIDFromHexString(s string) (id bson.ObjectId, err error) {
    if bson.IsObjectIdHex(s) {
        id = bson.ObjectIdHex(s)
        err = nil
    } else {
        err = errors.New("invalid object id provided")
    }

    return id, err
} 