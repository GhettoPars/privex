package handler

import (
	"context"
	"net/http"
	"privex/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e *Env) GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	user, err := e.Db.GetUser(context.Background(), int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (e *Env) PostUser(c *gin.Context) {
	var json struct {
		UserName string `json:"UserName" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}

	err := c.Bind(&json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	user, err := e.Db.CreateUser(context.Background(), database.CreateUserParams{UserName: json.UserName, UserRole: "user", Email: "example@mail.com", Password: json.Password})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}
