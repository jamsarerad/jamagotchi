package db

import (
  "fmt"
  "context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson"
  "golang.org/x/crypto/bcrypt"
)

func ConnectToUsers() *mongo.Collection {
  clientOptions := options.Client().ApplyURI("mongodb+srv://ae3whitlock:TKOfk68cRnEY5Iya@jamagotchi.oyrvaiw.mongodb.net/")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("jamagotchi").Collection("users")
	return collection
}

func RegisterUser(collection *mongo.Collection, username, email, password, leetCodeUsername string) (bool, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password: ", err)
    return false, err
	}

  //TODO: Preload the user's leetcode and authenticate it here
	newUser := bson.M{
    "username":             username,
    "email":                email,
    "hashed_password":      hashedPassword,
    "leetcode_username":    leetCodeUsername,
    "credits":              100,
    "leetcode_count": bson.M{
      "easy":               0,
      "medium":             0,
      "hard":               0,
    },
	}

	insertResult, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
    return false, err
	}

	fmt.Printf("Registered user with ID: %v\n", insertResult.InsertedID)

  return true, nil
}
