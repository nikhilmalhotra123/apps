package jobapps

import (
  "go.mongodb.org/mongo-driver/bson"
	"context"
	"github.com/nikhilmalhotra123/apps/jobapps/model"
	"github.com/nikhilmalhotra123/apps/jobapps/config/db"
)

// GetApplicationByID function
func GetApplicationByID(id string) (*model.Application, error) {
	ctx := getContext()

	collection, err := db.GetApplicationDB()
  if err != nil {
    return nil, err
  }

	var result model.Application
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAllApplicationsByUsername function
func GetAllApplicationsByUsername(username string) (*[]model.Application, error) {
	ctx := getContext()

	collection, err := db.GetApplicationDB()
  if err != nil {
    return nil, err
  }

	cur, err := collection.Find(ctx, bson.M{"username": username})
	if err != nil {
		return nil, err
	}

	var results []model.Application
	if err = cur.All(context.Background(), &results); err != nil {
  	return nil, err
	}
	return &results, nil
}

// InsertApplication into db
func InsertApplication(app *model.Application) (error) {
	ctx := getContext()

	collection, err := db.GetApplicationDB()
  if err != nil {
    return err
  }

	_, err = collection.InsertOne(ctx, *app)
	if err != nil {
		return err
	}

	return nil
}

// UpdateApplication into db
func UpdateApplication(app *model.Application) (error) {
	ctx := getContext()

	collection, err := db.GetApplicationDB()
  if err != nil {
    return err
  }

	appToUpdate := bson.M{
        "$set": app,
    }

	_, err = collection.UpdateOne(ctx, bson.M{"_id": app.ID}, appToUpdate)
	if err != nil {
		return err
	}

	return nil
}

// DeleteApplicationByID from db
func DeleteApplicationByID(id string) (error) {
	ctx := getContext()

	collection, err := db.GetApplicationDB()
  if err != nil {
    return err
  }

	_, err = collection.DeleteOne(ctx, bson.M{"_id": id});
	if err != nil {
		return err
	}

	return nil
}
