package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/keithkfield/go-messenger/messengerpb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection

// string userName = 1;
// string message = 2;
// string date = 3;
type item struct {
	ID       primitive.ObjectID
	Username string `bson: "username"`
	Message  string `bson: "message"`
	Date     string `bson: "date"`
}
type server struct{}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	port := "0.0.0.0:8080"
	fmt.Printf("Starting Server on Port: %v", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Unable to COnnect to port")
	}
	opt := []grpc.ServerOption{}
	s := grpc.NewServer(opt...)
	messengerpb.RegisterMessageServiceServer(s, &server{})
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://keithkfield:Icecat12!@goconnect-glenv.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatalf("Error connecting to Mongo Database: %v", err)
	}
	collection = client.Database("messages").Collection("messages")
	// END

	go func() {
		fmt.Printf("Listening on port: %v", port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve %v", err)
		}
		fmt.Print("working")
	}()
	// Wait for control c to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	// Block untill signal recieved
	<-ch
	fmt.Println("Stopping Server")
	s.Stop()
	fmt.Println("CLoesing Listner")
	lis.Close()
	fmt.Println("Cloesing mongo COnnection")
	client.Disconnect(context.TODO())
	fmt.Println("Cloesing Program")
}

func (*server) CreateMessage(ctx context.Context, req *messengerpb.CreateMessageRequest) (*messengerpb.CreateMessageResponse, error) {
	message := req.GetMessage()
	data := item{
		Username: message.GetUserName(),
		Message:  message.GetMessage(),
		Date:     strconv.Itoa(time.Year()),
	}
	return nil, nil
}
