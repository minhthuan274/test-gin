package users

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/minhthuan274/test-gin/models"
)

const (
	CollectionUser = "users"
)

func fetchUser(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	oID := bson.ObjectIdHex(c.Param("id"))
	var res models.User
	_ = db.C(CollectionUser).FindId(oID).One(&res)
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func fetchAllUsers(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var users []models.User
	_ = db.C(CollectionUser).Find(nil).All(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
