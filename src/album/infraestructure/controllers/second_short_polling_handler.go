package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"API_ejemplo/src/album/domain"
)

type ShortPollingPriceController struct {
	DBClient *mongo.Client
}

func NewShortPollingPriceController(client *mongo.Client) *ShortPollingPriceController {
	return &ShortPollingPriceController{DBClient: client}
}

func (s *ShortPollingPriceController) ShortPollingPriceHandler(c *gin.Context) {
	collection := s.DBClient.Database("MundyWalk").Collection("albums")

	var lastUpdated time.Time
	if err := c.BindQuery(&lastUpdated); err != nil {
		lastUpdated = time.Time{} 
	}

	timeout := time.After(10 * time.Second) 

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusNoContent, gin.H{"message": "No se detectaron cambios en el precio"})
			return

		default:
			var album domain.Album
			filter := bson.M{
				"Price": bson.M{"$gt": 0},
				"LastUpdated": bson.M{"$gt": lastUpdated},
			}

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
				"message": "Cambio en el precio detectado",
				"album":   album.Title, 
				"price":   album.Price,
				"modified": true,    
			})
			return
		}
	}
}
