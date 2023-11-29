package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type UserDetails struct {
	Name     string   `json: username`
	Password string   `json: password`
	UserId   int      `json: userID`
	Email    string   `json: email`
	ChatIds  []string `json: chatIDs`
}

type ChatDetails struct {
	message    string `json: message`
	SenderID   string `json: senderid`
	ReceiverID string `json: receiverid`
	messageID  string `json: messageid`
}

// Check Login details GET with request
// Fetch details from mongoDb
func Login(w http.ResponseWriter, r *http.Request) {

	// TODO: Connection to DB will move to init method in main.go
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://Cluster60322:fHpLfGVpWmZJ@cluster60322.jx4fmn1.mongodb.net/?appName=mongosh+2.1.0"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	userCollection := client.Database("myChat").Collection("Users")

	//Decode request body
	// Decode the JSON request body
	var req UserDetails
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", req.Name}},
				bson.D{{"password", req.Password}},
			},
		},
	}

	cursor, err := userCollection.Find(context.TODO(), filter)

	if err != nil {
		fmt.Println("Login details not found.. Please signup")
		//Redirect to error page
	} else {
		if cursor.RemainingBatchLength() > 0 {
			fmt.Print("Login successed")
		}
		//Redirect to main page that have chats
	}

}

// POST method that add entries to mongoDB
func SignUp(w http.ResponseWriter, r *http.Request) {
	// TODO: Connection to DB will move to init method in main.go
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://Cluster60322:fHpLfGVpWmZJ@cluster60322.jx4fmn1.mongodb.net/?appName=mongosh+2.1.0"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	userCollection := client.Database("myChat").Collection("Users")

	//Decode request body
	// Decode the JSON request body
	var req UserDetails
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Prepare UserData and enter
	user := bson.D{{"Name", req.Name}, {"Password", req.Password}, {"EmailID", req.Email}, {"UserID", rand.Intn(24000-14000) + 14000}}

	_, err = userCollection.InsertOne(context.TODO(), user)

	if err != nil {
		panic(err)
	}
}

// Get one on one chats
func Chat(w http.ResponseWriter, r *http.Request) {

}

// Get all chats
func AllChat(w http.ResponseWriter, r *http.Request) {
	// TODO: Connection to DB will move to init method in main.go
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://Cluster60322:fHpLfGVpWmZJ@cluster60322.jx4fmn1.mongodb.net/?appName=mongosh+2.1.0"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	userCollection := client.Database("myChat").Collection("Users")

	//Decode request body
	// Decode the JSON request body
	var req UserDetails
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Prepare UserData and enter
	filterUser := bson.D{{"Name", req.Name}}

	cursor, err := userCollection.Find(context.TODO(), filterUser)

	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, res := range results {
		fmt.Println(res)

	}

	if err != nil {
		fmt.Println("No chats found")
		//Redirect to error page
	} else {
		// TODO: Connection to DB will move to init method in main.go
		// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://Cluster60322:fHpLfGVpWmZJ@cluster60322.jx4fmn1.mongodb.net/?appName=mongosh+2.1.0"))
		// if err != nil {
		// 	panic(err)
		// }
		// if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		// 	panic(err)
		// }

		// userCollection := client.Database("myChat").Collection("Chats")

		// // Filter all chats with chatIds
		// filterchat := bson.D{{"ChatID", results.chatIds}}

		// chat, err := userCollection.Find(context.TODO(), filterchat)

		// //Print all chats
		// var results []bson.M

		// if err = chat.All(context.TODO(), &results); err != nil {
		// 	panic(err)
		// }

		// for _, res := range results {
		// 	res, _ := json.Marshal(res)
		// 	fmt.Println(res)
		// }

	}

}

// Real time chat
func TalkChat(w http.ResponseWriter, r *http.Request) {

}
