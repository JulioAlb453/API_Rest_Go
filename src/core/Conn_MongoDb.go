
package core

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func Connect() *mongo.Client {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error al cargar el archivo .env: %s", err)
    }

    MONGO_URI := os.Getenv("MONGO_URI")
    clientOptions := options.Client().ApplyURI(MONGO_URI)

    client, err := mongo.NewClient(clientOptions)
    if err != nil{
        log.Fatal("Error al crear el cliente de MongoDB: %v", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil{
        log.Fatal("Error al hacer la conexion con MongoDB: %v", err)
    }

    err = client.Ping(ctx, nil)
    if err != nil{
        log.Fatalf("Error al verificar la conexion con MongoDB", err)
    }
 
    log.Println("Conexion a MongoDB exitosa")
    return client
}

