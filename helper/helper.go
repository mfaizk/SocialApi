package helper

import (
	"context"
	"fmt"
	"log"
	"net/mail"
	"socialapi/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

const connectionString = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.3.1"
const dbName = "Social"
const colName = "usercredential"

func fatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	fatalErr(err)
	fmt.Println("MongoDB Connection sucessfull")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
	// addUserToDB(model.UserModel{})
}

func validateEmail(e string) error {
	_, err := mail.ParseAddress(e)
	if err != nil {
		return err
	}
	return nil
}

func checkDuplicateEmail(email string) bool {
	filter := bson.M{"email": email}
	cur, err := collection.Find(context.Background(), filter)
	fatalErr(err)

	if cur.TryNext(context.Background()) {
		return false
	}
	defer cur.Close(context.Background())
	return true

}
func AddUserToDB(u model.UserModel) string {
	if validateEmail(u.Email) == nil {
		if checkDuplicateEmail(u.Email) {
			passhash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
			fatalErr(err)
			u.Password = string(passhash)
			collection.InsertOne(context.Background(), u)
			return "nil"
		} else {
			return "Duplicate email"
		}

	} else {
		return validateEmail(u.Email).Error()
	}

}

func AuthChecker(u model.UserModel) bool {
	filter := bson.M{"email": u.Email}
	cur, err := collection.Find(context.Background(), filter)
	fatalErr(err)
	var us []model.UserModel
	for cur.Next(context.Background()) {
		var singleUser model.UserModel
		err := cur.Decode(&singleUser)
		fatalErr(err)

		us = append(us, singleUser)
	}
	defer cur.Close(context.Background())

	for _, myuserdb := range us {
		err := bcrypt.CompareHashAndPassword([]byte(myuserdb.Password), []byte(u.Password))
		if err != nil {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println("Welcome to Social Api")
// 	var pwd string = "faiz"
// 	pass := []byte(pwd)

// 	hash, _ := bcrypt.GenerateFromPassword(byteHash, bcrypt.MinCost)
// 	fmt.Println([]byte(pwd))
// 	fmt.Println(hash)
// 	err := bcrypt.CompareHashAndPassword(hash, pass)
// 	if err != nil {
// 		fmt.Println("false")
// 	}
// 	fmt.Println("true")

// }
