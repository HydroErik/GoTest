package main

import (
	//"encoding/base64"
	"fmt"
	//"encoding/json"
	"context"
	"log"
	"time"
	"os"
	
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func encryptPass(old_pass string) (string, error) {
	//encoded := base64.StdEncoding.EncodeToString([]byte(old_pass))
	encoded, err := bcrypt.GenerateFromPassword([]byte(old_pass), bcrypt.DefaultCost)
	return string(encoded), err
}



func main() {


	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found")
	}

	mongUrl := os.Getenv("MONGOSTRING")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongUrl))
	if err != nil { log.Fatal(err) }

	dbs, err := client.ListDatabaseNames(ctx, bson.D{})

	if err != nil {log.Fatal(err)}

	for _, val := range dbs {
		fmt.Println(val)
	}
	
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	
	user_collection := client.Database("GoUsers").Collection("gocredentials")
	
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//opt := options.Find().SetProjection(bson.D{{"esundblad", 1}})
	cur, err := user_collection.Find(ctx, bson.D{{"username", "esundblad"}})
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		fmt.Println(result["email"])
		
		pass := result["password"]
		encrypted, _ := encryptPass(pass.(string))
		err = bcrypt.CompareHashAndPassword([]byte(encrypted), []byte("SunWood4117!"))
		if err != nil{
			fmt.Println(err) 
		}
		

		fmt.Printf("Pass raw: %s\nEncrypted: %s", pass, encrypted)

		update := bson.D{{"$set", bson.D{{"password", encrypted}}}}
		opts := options.FindOneAndUpdate().SetUpsert(true)
		var updatedDocument bson.M 
		err = user_collection.FindOneAndUpdate(
			context.TODO(), 
			bson.D{{"username", "esundblad"}}, 
			update, 
			opts,
		).Decode(&updatedDocument)
		 
		if err != nil {
			fmt.Printf("Update Password failed with error: %s", err)
		}

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}





}

