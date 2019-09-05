package models

import (
	"github.com/gin-gonic/gin"
	"github.com/go-bongo/bongo"
)

// User model
type User struct {
	bongo.DocumentBase `bson:",inline" json:"_id"`
	Name               string `json:"name" binding:"required"`
	Gender             string `json:"gender" binding:"required"`
	Age                int    `json:"age" binding:"required"`
}

// Users slice
type Users []User

// Validate the request body with user model
// and retrun 406 response if any error
func (u *User) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(u); err != nil {
		return err
	}
	return nil
}
