package gin

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucas-dev3/frete-rapido-go.git/internal/http/presenter"
	"github.com/lucas-dev3/frete-rapido-go.git/quote"
)

func ProcessQuote(s quote.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("[Service] Processo de cotaçao iniciado")
		log.Println("estou no handler gin")
		var p presenter.QuoteRequest

		if err := c.BindJSON(&p); err != nil {
			log.Println("[Service] Erro ao fazer bind do JSON")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		eq := p.ToQuoteEntity()

		if err := s.ProcessQuote(c, eq); err != nil {
			log.Println("[Service] Erro ao processar cotaçao")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}

}

func MetricsQuote(s quote.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		lastQuote := c.Query("last_quotes")

		fmt.Println("lastQuote: ", lastQuote)
	}

}

func MakeQuoteHandlers(r *gin.RouterGroup, s quote.UseCase) {
	r.Handle("POST", "/quote", ProcessQuote(s))
	r.Handle("GET", "/quote", MetricsQuote(s))
}
