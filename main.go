package main

import (
	"fmt"
	"log"

	pb "github.com/dillonlpeterson/shippy-user-service/proto/user"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/mdns"
)

func main() {
	// Creates a database connection and handles
	// Closing it again before exit.
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Automatically migrates the user struct into database columns/types etc. This will
	// Check for chages and migrate them each time
	// this service is restarted
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}
	tokenService := &TokenService{repo}

	// Create a new service
	srv := micro.NewService(
		// This name must match the package name given in the protobuf definition
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags
	srv.Init()

	pubsub := srv.Server().Options().Broker
	// Will comment this out now to save having to run this locally
	// publisher := micro.NewPublisher("user.created", srv.Client())

	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService, pubsub})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
