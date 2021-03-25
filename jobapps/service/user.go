package jobapps

import (
	"github.com/nikhilmalhotra123/apps/jobapps/config/db"
	"github.com/nikhilmalhotra123/apps/jobapps/model"
  "go.mongodb.org/mongo-driver/bson"
)

// GetUserByUsername function
func GetUserByUsername(username string) (*model.User, error) {
	ctx := getContext()

	collection, err := db.GetUserDB()
  if err != nil {
    return nil, err
  }

	var result model.User
	if err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// InsertUser into db
func InsertUser(user *model.User) (error) {
	ctx := getContext()

	collection, err := db.GetUserDB()
  if err != nil {
    return err
  }

	_, err = collection.InsertOne(ctx, *user)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser into db
func UpdateUser(user *model.User) (error) {
	ctx := getContext()

	collection, err := db.GetUserDB()
  if err != nil {
    return err
  }

	_, err = collection.UpdateOne(ctx, bson.M{"username": user.Username}, user)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserByUsername from db
func DeleteUserByUsername(username string) (error) {
	ctx := getContext()

	collection, err := db.GetUserDB()
  if err != nil {
    return err
  }

	_, err = collection.DeleteOne(ctx, bson.M{"username": username});
	if err != nil {
		return err
	}

	return nil
}
