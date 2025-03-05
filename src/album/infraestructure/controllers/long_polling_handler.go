package controllers

import (
	"API_ejemplo/src/album/domain"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LongPollingController struct{
	DBClient *mongo.Client
}

func NewLongPollingController(dbClient *mongo.Client) *LongPollingController{
	return &LongPollingController{DBClient: dbClient}
}

func (lpc *LongPollingController) LongPollingHandler(c *gin.Context) {
	collection := lpc.DBClient.Database("MundyWalk").Collection("albums")

	var lastUpdated time.Time
	if err := c.BindQuery(&lastUpdated); err != nil {
		lastUpdated = time.Time{} 
	}

	timeout := time.After(10 * time.Second) 

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusNoContent, gin.H{"message": "No se detectaron cambios"})
			return

		default:
			var album domain.Album

			filter := bson.M{"LastUpdated": bson.M{"$gt": lastUpdated}}

			err := collection.FindOne(c, filter).Decode(&album)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					time.Sleep(2 * time.Second) 
					continue
				}

				log.Println("Error al obtener álbum:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener álbum"})
				return
			}

			lastUpdated = album.LastUpdated
			c.JSON(http.StatusOK, gin.H{
				"message": "Cambio detectado en el álbum",
				"album":   album.Title, 
				"modified": true, 
			})
			return
		}
	}
}