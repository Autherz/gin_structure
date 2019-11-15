package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
)

type Todo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
}

func (b *Todo) TableName() string {
	return "todo"
}

// func GetAllTodos(todo *Todo) (err error) {

// 	// value1 := Todo{ID: 1, Title: "Article 1", Description: "Article 1 body"}
// 	// value2 := Todo{ID: 2, Title: "Article 2", Description: "Article 2 body"}
// 	todo.ID = 1
// 	todo.Title = "test"
// 	todo.Description = "test-desc"
// 	return nil
// }

func GetAllTodos(todo *[]Todo) (err error) {

	// value1 := Todo{ID: 1, Title: "Article 1", Description: "Article 1 body"}
	// value2 := Todo{ID: 2, Title: "Article 2", Description: "Article 2 body"}
	*todo = append(*todo, Todo{Title: "Article 1", Description: "Article 1 body"})
	*todo = append(*todo, Todo{Title: "Article 2", Description: "Article 2 body"})
	// todo.ID = 1
	// todo.Title = "test"
	// todo.Description = "test-desc"
	return nil
}

func TestTemp(todo *Todo) error {
	todo.Title = "test2"
	return nil
}

func InsertTodo(todo *Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := ConnectClient(ctx)
	db := client.Database("dev_db")
	collection := db.Collection("todo")
	defer client.Disconnect(ctx)

	collection.InsertOne(ctx, todo)
	return err
}

func GetAllTodo(todos *[]Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := ConnectClient(ctx)
	db := client.Database("dev_db")
	collection := db.Collection("todo")
	defer client.Disconnect(ctx)

	cur, _ := collection.Find(ctx, bson.D{})
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var todo Todo
		// fmt.Printf("Before Decode Todo %s\n", todo.Title)
		cur.Decode(&todo)
		// fmt.Printf("After Decode Todo %s\n", todo.Title)
		*todos = append(*todos, todo)
	}

	return nil
}

func GetTodo(todo *Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := ConnectClient(ctx)
	db := client.Database("dev_db")
	collection := db.Collection("todo")
	defer client.Disconnect(ctx)

	// filter := bson.M{"$or": []interface{}{
	// 	bson.M{"username": u.Username},
	// 	bson.M{"profile.email": u.Profile.Email},
	// },
	
	filter := bson.M{"_id": todo.ID}
	collection.FindOne(ctx, filter).Decode(&todo)
	return nil
}

func PutTodo(todo *Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := ConnectClient(ctx)
	db := client.Database("dev_db")
	collection := db.Collection("todo")
	defer client.Disconnect(ctx)

	filter := bson.M{"_id": todo.ID}
	update := bson.M{"$set": todo}
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func DeleteTodo(todo *Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := ConnectClient(ctx)
	db := client.Database("dev_db")
	collection := db.Collection("todo")
	defer client.Disconnect(ctx)

	filter := bson.M{"_id": todo.ID}
	deleteResult, err := collection.DeleteOne(ctx, filter)
	fmt.Printf("Delete count is %d", deleteResult.DeletedCount)
	return err
}
