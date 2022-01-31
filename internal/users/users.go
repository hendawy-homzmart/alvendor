package users

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/El-Hendawy/gograph/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

const (
	DATABASE = "myFirstDatabase"

	COLLECTION = "sellers"

	COLLECTIONCUSTOMER = "customers"

	COLLECTIONADMIN = "admins"
)

type database struct {
	client *mongo.Client
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsersRepository interface {
	GetUserIdByUsername(username string) (string, error)
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func (db *database) GetUserIdByUsername(email string) (string, error) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)

	filter := bson.M{"userinfo.email": email}

	//	var result *model.UserData

	var respone *model.Seller

	err := collection.FindOne(context.TODO(), filter).Decode(&respone)

	if err != nil {
		println(err)
	}

	var Id string
	Id = respone.UserInfo.ID

	return Id, nil
}
func NewValidate() UsersRepository {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://hendawy:no1canbmeh@cluster0.lfvaq.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	clientOptions = clientOptions.SetMaxPoolSize(55)
	userClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = userClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return &database{
		client: userClient,
	}
}
