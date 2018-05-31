package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhthuan274/test-gin/db"
	"github.com/minhthuan274/test-gin/middleware"
	"github.com/minhthuan274/test-gin/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func init() {
	db.Connect()
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.Use(middleware.Connect)
	r.Use(middleware.ErrorHandler)

	v3 := r.Group("/api/v3")
	v3.Use(middleware.Auth())
	{
		v3.POST("/reviews", postReview)
		v3.GET("/home", getHome)
	}

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func getHome(c *gin.Context) {
	userID := c.GetString("userID")
	c.JSON(http.StatusOK, userID)
}

func postReview(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	userID := c.GetString("userID")
	oUserID := bson.ObjectIdHex(userID)
	ID := bson.NewObjectId()
	var json models.ReviewJson
	c.BindJSON(&json)
	err := db.C(models.CollectionReview).Insert(models.Review{
		ID,
		oUserID,
		bson.ObjectIdHex(json.Merchant),
		json.Feedback,
		json.Point,
	})

	if err != nil {
		c.JSON(http.StatusNotAcceptable, err)
		return
	}

	reviews := models.Review{}
	_ = db.C(models.CollectionReview).Find(bson.M{"_id": ID}).One(&reviews)

	c.JSON(http.StatusCreated, reviews)
}
