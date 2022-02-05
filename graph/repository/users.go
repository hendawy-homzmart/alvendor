package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/El-Hendawy/gograph/graph/model"
	"github.com/El-Hendawy/gograph/jwt"
	"golang.org/x/crypto/bcrypt"

	//	"github.com/El-Hendawy/gograph/jwt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type AdminRepository interface {
	SaveAdmin(admin *model.Admin)
	FindAllAdmin() []*model.Admin
}

type ProductRepository interface {
	SaveProduct(product *model.Product)
	//FindAllPro() []*model.Product
}

type SellerRepository interface {
	Save(seller *model.Seller)
	FindAll() []*model.Seller
	AuthenticateSeller(login *model.Login) (*model.UserData, error)
}
type CustomerRepository interface {
	SaveCustomer(customer *model.Customer)
	FindAllCustomers() []*model.Customer
}

type database struct {
	client *mongo.Client
}

const (
	DATABASE = "myFirstDatabase"

	COLLECTION = "sellers"

	COLLECTIONCUSTOMER = "customers"

	COLLECTIONADMIN = "admins"

	COLLECTIONPRODUCT = "products"
)

func NewProduct() ProductRepository {
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

func (db *database) AuthenticateSeller(login *model.Login) (*model.UserData, error) {

	collection := db.client.Database(DATABASE).Collection(COLLECTION)

	filter := bson.M{"userinfo.email": login.Email}

	var result *model.UserData

	var respone *model.Seller

	err := collection.FindOne(context.TODO(), filter).Decode(&respone)

	if err == mongo.ErrNoDocuments {

		fmt.Println("record does not exist")

	} else if err != nil {

		log.Fatal(err)

	}

	result = &model.UserData{

		UserData: respone.UserInfo,

		Token: "",
	}

	check := CheckPasswordHash(login.Password, respone.UserInfo.Password)

	if check {
		token, err := jwt.GenerateToken(result.UserData)

		if err != nil {

			log.Fatal(err)

		}

		result.Token = token

	}
	// } else {
	// 	result = &model.UserData{
	// 		UserData: &model.User{},
	// 		Token:    "",
	// 	}
	// }
	return result, err

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (db *database) SaveProduct(product *model.Product) {
	collection := db.client.Database(DATABASE).Collection(COLLECTIONPRODUCT)
	_, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Fatal(err)
	}
}
