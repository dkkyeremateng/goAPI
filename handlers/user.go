package handlers

import (
	"userAPI/models"

	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

// Handlers handles all user routes
type Handlers struct {
	conn *bongo.Connection
}

// List handles the GET /users request
func (h *Handlers) List(c *gin.Context) {
	user := &models.User{}
	userSlice := models.Users{}

	users := h.conn.Collection("users").Find(bson.M{})
	for users.Next(user) {
		userSlice = append(userSlice, *user)
	}

	c.JSON(200, userSlice)
}

// Show handles the GET /users/:id request
func (h *Handlers) Show(c *gin.Context) {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		c.Status(404)
		return
	}

	user := &models.User{}
	if err := h.conn.Collection("users").FindById(bson.ObjectIdHex(id), user); err != nil {
		c.Status(404)
		return
	}
	c.JSON(200, user)
}

// Post handles the POST /users request
func (h *Handlers) Post(c *gin.Context) {
	user := &models.User{}

	if err := user.Validate(c); err != nil {
		c.JSON(406, gin.H{"error": err.Error()})
		return
	}

	if err := h.conn.Collection("users").Save(user); err != nil {
		c.Status(400)
		return
	}
	c.Status(201)
}

// Delete handles the DELETE /users/:id request
func (h *Handlers) Delete(c *gin.Context) {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		c.Status(404)
		return
	}

	user := &models.User{}
	if err := h.conn.Collection("users").FindById(bson.ObjectIdHex(id), user); err != nil {
		c.Status(404)
		return
	}

	if err := h.conn.Collection("users").DeleteDocument(user); err != nil {
		c.Status(400)
		return
	}
	c.Status(204)
}

// Update handles the PUT /users/:id request
func (h *Handlers) Update(c *gin.Context) {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		c.Status(404)
		return
	}

	user := &models.User{}
	if err := user.Validate(c); err != nil {
		c.JSON(406, gin.H{"error": err.Error()})
		return
	}

	userFromDB := &models.User{}
	if err := h.conn.Collection("users").FindById(bson.ObjectIdHex(id), userFromDB); err != nil {
		c.Status(404)
		return
	}

	userFromDB.Age = user.Age
	userFromDB.Gender = user.Gender
	userFromDB.Name = user.Name

	if err := h.conn.Collection("users").Save(userFromDB); err != nil {
		c.Status(400)
		return
	}
	c.Status(201)
}

// SetupRoutes setup routers for user
func (h *Handlers) SetupRoutes(r *gin.Engine) {
	r.GET("/api/users", h.List)
	r.GET("/api/users/:id", h.Show)
	r.POST("/api/users", h.Post)
	r.PUT("/api/users/:id", h.Update)
	r.DELETE("/api/users/:id", h.Delete)
}

// NewHandlers returns new handler
func NewHandlers(c *bongo.Connection) *Handlers {
	return &Handlers{
		conn: c,
	}
}
