package handler

import (
	"context"
	"fmt"
	"net/http"
	"privex/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

// List Messages
func (e *Env) ListMessages(c *gin.Context) {
	messages, err := e.Db.ListMessages(context.Background())
	fmt.Println(messages, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, messages)
}

func (e *Env) GetMessage(c *gin.Context) {
	id := c.Params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	message, err := e.Db.GetMessage(context.Background(), int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, message)
}

func (e *Env) PostMessage(c *gin.Context) {
	var json struct {
		UserID      int32  `json:"UserID" binding:"required"`
		MessageText string `json:"MessageText" binding:"required"`
		MessageType string `json:"MessageType" binding:"required"`
	}

	err := c.Bind(&json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	fmt.Println(json)
	message, err := e.Db.CreateMessage(context.Background(), database.CreateMessageParams{UserID: json.UserID, MessageText: json.MessageText, MessageType: json.MessageType})
	fmt.Println(message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Println(message)
	c.JSON(http.StatusOK, message)
	return

}

func (e *Env) DeleteMessage(c *gin.Context) {
	id := c.Params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = e.Db.DeleteMessage(context.Background(), int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "message deleted"})
}
