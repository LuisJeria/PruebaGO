package main

import (
	"context"
	"fmt"
	"hola/internal/controller"
	"hola/internal/repository"
	"hola/internal/service"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	loadEnviromentVariables()
	ctx := context.Background()
	client, _ := getFirestoreClient(ctx)
	personajeRepository := repository.NewPersonajeRepository(client)
	personajeService := service.NewPersonajeService(personajeRepository)

	fmt.Println("starting firestore client project id : " + os.Getenv("GOOGLE_PROJECT_ID"))
	controller.NewPersonajeController(personajeService)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

}

func loadEnviromentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func getFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	c, err := firestore.NewClient(ctx, os.Getenv("GOOGLE_PROJECT_ID"), option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		log.Fatal("Error getting firestore client")
	}
	return c, nil
}
