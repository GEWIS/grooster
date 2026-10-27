package roster

import (
	"GEWIS-Rooster/internal/models"
	"GEWIS-Rooster/internal/platform/httputil"
	"GEWIS-Rooster/internal/platform/middleware"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (h *Handler) registerCommentRoutes(roster *gin.RouterGroup, db *gorm.DB) {
	roster.POST("/comments", RequireRosterOrganRoleParam(db, "rosterId", models.RoleMember), h.CreateRosterComment)
	roster.GET("/comments", RequireRosterOrganRoleParam(db, "rosterId", models.RoleMember), h.GetRosterComments)
}

// CreateRosterComment
//
//	@Summary	Create a new roster comment
//	@Security	BearerAuth
//	@Tags		Roster Comment
//	@Accept		json
//	@Produce	json
//	@Param		rosterId		path		int						true	"Roster ID"
//	@Param		createParams	body		CommentCreateRequest	true	"Roster comment input"
//	@Success	201				{object}	models.RosterComment
//	@Failure	400				{string}	string
//	@Failure	403				{string}	string
//	@Failure	404				{string}	string
//	@ID			createRosterComment
//	@Router		/rosters/{rosterId}/comments [post]
func (h *Handler) CreateRosterComment(c *gin.Context) {
	rosterID, ok := httputil.MustPathUint(c, "rosterId")
	if !ok {
		return
	}

	var param *CommentCreateRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	callerID, exists := c.Get(middleware.UserIDKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if callerID.(uint) != param.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only comment on your own roster entry"})
		return
	}

	createdComment, err := h.rosterService.CreateRosterComment(rosterID, param)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Roster not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdComment)
}

// GetRosterComments
//
//	@Summary	Get all comments for a roster
//	@Security	BearerAuth
//	@Tags		Roster Comment
//	@Accept		json
//	@Produce	json
//	@Param		rosterId	path		int	true	"Roster ID"
//	@Success	200			{array}		models.RosterComment
//	@Failure	400			{string}	string
//	@Failure	404			{string}	string
//	@ID			getRosterComments
//	@Router		/rosters/{rosterId}/comments [get]
func (h *Handler) GetRosterComments(c *gin.Context) {
	rosterID, ok := httputil.MustPathUint(c, "rosterId")
	if !ok {
		return
	}

	comments, err := h.rosterService.GetRosterComments(rosterID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}
