package mongooperations

import (
	"context"
	"log"
	"time"

	"github.com/dantdj/RecipeBook/pkg/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeDb struct {
	Client *mongo.Client
}

func (db *RecipeDb) Insert(userId, name, ingredients, method string) (string, error) {
	collection := db.Client.Database("recipestore").Collection("recipes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := uuid.NewString()

	_, err := collection.InsertOne(ctx, bson.M{"_id": id, "name": name, "ingredients": ingredients, "method": method, "userId": userId})
	if err != nil {
		return "", err
	}

	return id, nil
}

func (db *RecipeDb) Get(id string) (*models.Recipe, error) {
	collection := db.Client.Database("recipestore").Collection("recipes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	recipe := models.Recipe{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&recipe)
	if err != nil {
		log.Fatal(err)
	}

	return &recipe, nil
}
