package user

import (
	_ "GEWIS-Rooster/internal/models"
	"GEWIS-Rooster/internal/platform/httputil"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	userService Service
}

func NewUserHandler(rg *gin.RouterGroup, userService Service) *Handler {
	h := &Handler{userService: userService}

	g := rg.Group("/users")

	g.POST("", h.Create)
	g.GET("", h.GetAllUsers)
	g.GET("/:userId", h.GetUserByID)
	g.DELETE("/:userId", h.Delete)

	return h
}

// Create
//
//	@Summary		CreateRoster a new user
//	@Security		BearerAuth
//	@Description	create user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			createParams	body		UserCreateRequest	true	"User input"
//	@Success		200				{object}	models.User
//	@Failure		400				{string}	string
//	@ID				createUser
//	@Router			/users [post]
func (h *Handler) Create(c *gin.Context) {
	var param *CreateRequest

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	createdUser, err := h.userService.Create(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": createdUser,
	})
}

// GetAllUsers
//
//		@Summary      Get all users with optional filtering
//		@Description  Retrieve a list of users with optional query parameter filtering
//		@Security     BearerAuth
//		@Tags         User
//		@Accept       json
//		@Produce      json
//	 @Param        organId    query     uint    false  "Organ ID"
//	 @Param        gewisId    query     uint    false  "GEWIS ID"
//		@Success      200         {array}   models.User
//		@Failure      400         {object}  map[string]string
//		@ID           getUsers
//		@Router       /users [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	f := FilterParams{}
	if err := c.ShouldBindQuery(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := h.userService.Get(&f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID
//
//	@Summary      Get user by ID
//	@Description  Retrieve a specific user by their unique ID
//	@Security     BearerAuth
//	@Tags         User
//	@Accept       json
//	@Produce      json
//	@Param        userId      path      uint    true   "User ID"
//	@Success      200         {object}  models.User
//	@Failure      400         {object}  map[string]string
//	@Failure      404         {object}  map[string]string
//	@ID           getUser
//	@Router       /users/{userId} [get]
func (h *Handler) GetUserByID(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "userId")
	if !ok {
		return
	}

	filters := &FilterParams{ID: &id}

	users, err := h.userService.Get(filters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Delete
//
//	@Summary	DeleteRoster a user
//	@Security	BearerAuth
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		userId	path		int	true	"User ID"
//	@Success	200	{string}	string
//	@Failure	400	{string}	string
//	@Failure	404	{string}	string
//	@ID			deleteUser
//	@Router		/users/{userId} [delete]
func (h *Handler) Delete(c *gin.Context) {
	userID, ok := httputil.MustPathUint(c, "userId")
	if !ok {
		return
	}

	if err := h.userService.Delete(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}
