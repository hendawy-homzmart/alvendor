package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/El-Hendawy/gograph/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)


type AdminRepository interface {
	SaveAdmin(customer *model.Admin)
	FindAllAdmin() []*model.Admin
}

type SellerRepository interface {
	Save(seller *model.Seller)
	FindAll() []*model.Seller
}
type CustomerRepository interface {
	SaveCustomer(customer *model.Customer)
	FindAllCustomers() []*model.Customer
}

type database struct {
	client *mongo.Client
}

const (

	DATABASE   = "myFirstDatabase"
	
	COLLECTION = "sellers"

	COLLECTIONCUSTOMER = "customers"
	
	COLLECTIONADMIN = "admins"

)




func NewAdmin() AdminRepository {
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
func New() SellerRepository {
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




func NewCustomer() CustomerRepository {
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



func (db *database) Save(seller *model.Seller) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), seller)
	if err != nil {
		log.Fatal(err)
	}
}
func (db *database) FindAll() []*model.Seller {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	options := options.Find()
	options.SetSort(bson.M{"first_name": 1})
	options.SetSkip(0)
	options.SetLimit(9)
	cursor, err := collection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var result []*model.Seller
	for cursor.Next(context.TODO()) {
		var v *model.Seller
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}
	return result
}


func (db *database) SaveCustomer(customer *model.Customer) {
	collection := db.client.Database(DATABASE).Collection(COLLECTIONCUSTOMER)
	_, err := collection.InsertOne(context.TODO(), customer)
	if err != nil {
		log.Fatal(err)
	}
}
func (db *database) FindAllCustomers() []*model.Customer {
	collection := db.client.Database(DATABASE).Collection(COLLECTIONCUSTOMER)
	options := options.Find()
	options.SetSort(bson.M{"first_name": 1})
	options.SetSkip(0)
	options.SetLimit(9)
	cursor, err := collection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var result []*model.Customer
	for cursor.Next(context.TODO()) {
		var v *model.Customer
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}
	return result
}




func (db *database) SaveAdmin(admin *model.Admin) {
	collection := db.client.Database(DATABASE).Collection(COLLECTIONADMIN)
	_, err := collection.InsertOne(context.TODO(), admin)
	if err != nil {
		log.Fatal(err)
	}
}
func (db *database) FindAllAdmin() []*model.Admin {
	collection := db.client.Database(DATABASE).Collection(COLLECTIONADMIN)
	options := options.Find()
	options.SetSort(bson.M{"first_name": 1})
	options.SetSkip(0)
	options.SetLimit(9)
	cursor, err := collection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var result []*model.Admin
	for cursor.Next(context.TODO()) {
		var v *model.Admin
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}
	return result
}