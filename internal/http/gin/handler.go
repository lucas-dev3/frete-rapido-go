package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucas-dev3/frete-rapido-go.git/config"
	"github.com/lucas-dev3/frete-rapido-go.git/quote"
)

// type errorResponse struct {
// 	Error string `json:"error" xml:"error"`
// }

func Handlers(envs *config.Environments, quoteService quote.UseCase) *gin.Engine {
	routers := gin.Default()

	routers.GET("/health", healthHandler)
	v1 := routers.Group("/api/v1")

	MakeQuoteHandlers(v1, quoteService)

	return routers
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func IsEmpty(value string) bool {
	return value == "" || value == "  "
}
