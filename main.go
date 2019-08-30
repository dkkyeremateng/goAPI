package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"userAPI/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	users := models.Users{
		{
			Name:   "Daniel",
			Gender: "M",
			Age:    25,
			ID:     1,
		},
	}

	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		for _, u := range users {
			if u.ID == id {
				c.JSON(200, u)
				return
			}
		}
		c.Status(404)
	})

	r.POST("/users", func(c *gin.Context) {
		user := &models.User{}

		if err := user.Validate(c); err != nil {
			c.JSON(406, gin.H{"error": err.Error()})
		}

		user.ID = (len(users) + 1)

		users = append(users, *user)

		c.Status(201)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id, err := checkID(c.Param("id"))
		if err != nil {
			c.Status(400)
			return
		}

		for i, u := range users {
			if u.ID == id {
				users = append(users[:i], users[1:]...)
				c.Status(204)
				return
			}
		}

		c.Status(404)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		id, err := checkID(c.Param("id"))
		if err != nil {
			c.Status(400)
			return
		}

		user := &models.User{}

		if err := user.Validate(c); err != nil {
			c.JSON(406, gin.H{"error": err.Error()})
		}

		for i, u := range users {
			if u.ID == id {
				users = append(users[:i], users[1:]...)
				users = append(users, *user)
				c.JSON(204, user)
				return
			}
		}
		c.Status(404)
	})

	if err := r.Run(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

func checkID(s string) (interface{}, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return nil, errors.New("Unable to convent id to int")
	}
	return id, nil
}
