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
		var p presenter.QuoteRequest

		if err := bindData(c, &p); err != nil {
			log.Println("[Service] Erro ao processar request")
			respondAccept(c, http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := s.ProcessQuote(c, &p); err != nil {
			log.Println("[Service] Erro ao processar cotaçao")
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
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
