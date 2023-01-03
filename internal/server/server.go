package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/multimoml/stats/docs"
	"github.com/multimoml/stats/internal/config"
	"github.com/multimoml/stats/internal/model"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run(ctx context.Context) {
	// Load environment variables
	cfg := config.LoadConfig()

	// Set up router
	router := gin.Default()

	// Endpoints
	router.GET("/stats/live", Liveness)
	router.GET("/stats/ready", Readiness)
	router.GET("/stats/openapi", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/stats/openapi/index.html")
	})

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/stats/openapi/index.html")
	})

	router.GET("/openapi/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1
	v1 := router.Group("/stats")
	{
		v1.GET("/all", Stats)
	}

	// Start HTTP server
	log.Fatal(router.Run(fmt.Sprintf(":%s", cfg.Port)))
}

// Liveness is a simple endpoint to check if the server is alive
// @Summary Get liveness status of the microservice
// @Description Get liveness status of the microservice
// @Tags Kubernetes
// @Success 200 {string} string
// @Router /live [get]
func Liveness(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "alive"})
}

// Readiness is a simple endpoint to check if the server is ready
// @Summary Get readiness status of the microservice
// @Description Get readiness status of the microservice
// @Tags Kubernetes
// @Success 200 {string} string
// @Failure 503 {string} string
// @Router /ready [get]
func Readiness(c *gin.Context) {
	dispatcher := "http://dispatcher:6001"

	// if using dev environment access local tracker
	if os.Getenv("ACTIVE_ENV") == "dev" {
		dispatcher = "http://localhost:6001"
	}

	_, err := http.Get(dispatcher + "/products/ready")
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusServiceUnavailable, gin.H{"status": "not ready"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "ready"})

}

// Stats returns statistics about products
// @Summary Get all statistics
// @Description Get all statistics
// @Tags Statistics
// @Produce json
// @Success 200 {object} object
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /all [get]
func Stats(c *gin.Context) {
	dispatcher := "http://dispatcher:6001"

	// if using dev environment access local stats
	if os.Getenv("ACTIVE_ENV") == "dev" {
		dispatcher = "http://localhost:6001"
	}

	res, err := http.Get(dispatcher + "/products/v1/all")
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// decode JSON response into products
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// decode body into products
	var products []model.Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// get all category names
	var categoryNames []string
	for _, product := range products {
		categoryNames = append(categoryNames, product.CategoryName)
	}
	if len(categoryNames) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no category names found"})
		return
	}

	// number of all products by categoryName in key value pair
	var productsByCategoryName = make(map[string]int)
	for _, product := range products {
		productsByCategoryName[product.CategoryName]++
	}
	if len(productsByCategoryName) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no products found"})
		return
	}
	c.IndentedJSON(http.StatusOK, productsByCategoryName)

	// number of all products by Brand in key value pair
	var productsByBrand = make(map[string]int)
	for _, product := range products {
		productsByBrand[product.Brand]++
	}
	if len(productsByBrand) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no products found"})
		return
	}
	c.IndentedJSON(http.StatusOK, productsByBrand)

	// number of all products by PriceInTime IsOnPromotion true by brand in key value pair
	var productsOnPromotionByBrand = make(map[string]int)
	for _, product := range products {
		if product.PriceInTime[0].IsOnPromotion {
			productsOnPromotionByBrand[product.Brand]++
		}
	}
	if len(productsOnPromotionByBrand) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no products found"})
		return
	}
	c.IndentedJSON(http.StatusOK, productsOnPromotionByBrand)

}
