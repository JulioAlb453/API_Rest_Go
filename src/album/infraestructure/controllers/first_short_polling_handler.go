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

type ShortPollingStockController struct {
	DBClient *mongo.Client
}

func NewShortPollingStockController(client *mongo.Client) *ShortPollingStockController {
	return &ShortPollingStockController{DBClient: client}
}

// ShortPollingStockHandler maneja las solicitudes de short polling
func (s *ShortPollingStockController) ShortPollingStockHandler(c *gin.Context) {
	collection := s.DBClient.Database("MundyWalk").Collection("albums")

	// Recuperar el valor actual de "LastUpdated" del cliente si lo envía
	var lastUpdated time.Time
	if err := c.BindQuery(&lastUpdated); err != nil {
		lastUpdated = time.Time{} // Si no se envía, establecer como fecha inicial (epoch)
	}

	// Canal para manejar el tiempo de espera del short polling
	timeout := time.After(10 * time.Second) // Máximo tiempo de espera de 10 segundos

	for {
		select {
		case <-timeout:
			// Si se alcanza el tiempo de espera, responder sin cambios
			c.JSON(http.StatusNoContent, gin.H{"message": "No se detectaron cambios en el stock"})
			return

		default:
			var album domain.Album

			// Filtro para buscar álbumes con un stock mayor a 0 y actualizaciones después de "lastUpdated"
			filter := bson.M{
				"Stock": bson.M{"$gt": 0},
				"LastUpdated": bson.M{"$gt": lastUpdated},
			}

			// Buscar el álbum con cambios en el stock
			err := collection.FindOne(c, filter).Decode(&album)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					// Si no hay documentos, continuar esperando
					time.Sleep(2 * time.Second) // Esperar antes de volver a intentar
					continue
				}

				// Si ocurre un error, responder con el error
				log.Println("Error al obtener álbum:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener álbum"})
				return
			}

			// Si se encuentra un álbum actualizado, responder con los datos
			lastUpdated = album.LastUpdated
			c.JSON(http.StatusOK, gin.H{
				"message": "Cambio en el stock detectado",
				"album":   album.Title, // Solo el título del álbum
				"stock":   album.Stock, // Nuevo valor de stock
				"modified": true,      // Indica que el stock fue modificado
			})
			return
		}
	}
}