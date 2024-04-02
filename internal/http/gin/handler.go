package gin

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucas-dev3/frete-rapido-go.git/config"
	"github.com/lucas-dev3/frete-rapido-go.git/quote"
	"gopkg.in/yaml.v3"
)

// type errorResponse struct {
// 	Error string `json:"error" xml:"error"`
// }

func Handlers(envs *config.Environments, quoteService quote.UseCase) *gin.Engine {
	routers := gin.Default()

	routers.GET("/health", healthHandler)
	v1 := routers.Group("/api/v1")

	pg := v1.Group("/quote")
	MakeQuoteHandlers(pg, quoteService)

	return routers
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func respondAccept(c *gin.Context, status int, data interface{}) {
	fmt.Println(c.GetHeader("Accept"))

	switch c.GetHeader("Accept") {
	case "text/xml", "application/xml":
		c.XML(status, data)
	case "application/json":
		c.JSON(status, data)
	case "application/x-yaml", "text/x-yaml", "text/yaml", "application/yaml":
		yamlData, err := yaml.Marshal(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Data(status, "application/x-yaml", yamlData)
		return
	default:
		c.JSON(status, data)
		return
	}
}

func bindData(c *gin.Context, data interface{}) error {
	switch c.GetHeader("Content-Type") {
	case "application/xml", "text/xml", "application/json":
		if err := c.ShouldBind(data); err != nil {
			return err
		}
	case "application/x-yaml", "text/x-yaml", "text/yaml", "application/yaml":
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(body, data); err != nil {
			return err
		}
	default:
		if err := c.BindJSON(data); err != nil {
			return err
		}
	}
	return nil
}

func IsEmpty(value string) bool {
	return value == "" || value == "  "
}
