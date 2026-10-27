package roster

import (
	"GEWIS-Rooster/internal/models"
	"GEWIS-Rooster/internal/platform/httputil"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (h *Handler) registerRosterRoutes(organ *gin.RouterGroup, roster *gin.RouterGroup, db *gorm.DB) {
	organ.POST("/rosters", requireOrganRoleParam(db, "organId", models.RoleAdmin), h.CreateRoster)
	organ.GET("/rosters", requireOrganRoleParam(db, "organId", models.RoleMember), h.GetRosters)
	roster.GET("", RequireRosterOrganRoleParam(db, "rosterId", models.RoleMember), h.GetRoster)
	roster.PATCH("", RequireRosterOrganRoleParam(db, "rosterId", models.RoleAdmin), h.UpdateRoster)
	roster.DELETE("", RequireRosterOrganRoleParam(db, "rosterId", models.RoleAdmin), h.DeleteRoster)
}

// CreateRoster
//
//	@Summary	CreateRoster a new roster
//	@Security	BearerAuth
//	@Tags		Roster
//	@Accept		json
//	@Produce	json
//	@Param		organId			path		int					true	"Organ ID"
//	@Param		createParams	body		RosterCreateRequest	true	"Roster input"
//	@Success	200				{object}	models.Roster
//	@Failure	400				{string}	string
//	@ID			createRoster
//	@Router		/organs/{organId}/rosters [post]
func (h *Handler) CreateRoster(c *gin.Context) {
	organID, ok := httputil.MustPathUint(c, "organId")
	if !ok {
		return
	}

	var param *CreateRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	createdRoster, err := h.rosterService.CreateRoster(organID, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdRoster)
}

// GetRosters
//
//		@Summary	Get all rosters of an organ, optionally filtered by date or archived state
//		@Security	BearerAuth
//		@Tags		Roster
//		@Accept		json
//		@Produce	json
//	 @Param     organId path      int                 true   "Organ ID"
//	 @Param     filter  query     RosterFilterParams  false  "Filter parameters"
//		@Success	200		{array}		models.Roster
//		@Failure	400		{string}	string
//		@ID			getRosters
//		@Router		/organs/{organId}/rosters [get]
func (h *Handler) GetRosters(c *gin.Context) {
	organID, ok := httputil.MustPathUint(c, "organId")
	if !ok {
		return
	}

	var params FilterParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	rosters, err := h.rosterService.GetRosters(&organID, &params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, rosters)
}

// GetRoster
//
//	@Summary	Get a specific roster by id
//	@Security	BearerAuth
//	@Tags		Roster
//	@Accept		json
//	@Produce	json
//
//	@Param		rosterId	path		uint	true	"Roster ID"
//
//	@Success	200	{object}	models.Roster
//	@Failure	400	{string}	string
//	@Failure	404	{string}	string
//	@ID			getRoster
//	@Router		/rosters/{rosterId} [get]
func (h *Handler) GetRoster(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "rosterId")
	if !ok {
		return
	}

	params := &FilterParams{
		ID: &id,
	}
	roster, err := h.rosterService.GetRosters(nil, params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roster)
}

// UpdateRoster
//
//	@Summary	Update a roster
//	@Security	BearerAuth
//	@Tags		Roster
//	@Accept		json
//	@Produce	json
//	@Param		rosterId		path		uint						true	"Roster ID"
//	@Param		updateParams	body		RosterUpdateRequest	true	"Roster input"
//	@Success	200				{object}	models.Roster
//	@Failure	400				{string}	string
//	@ID			updateRoster
//	@Router		/rosters/{rosterId} [patch]
func (h *Handler) UpdateRoster(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "rosterId")
	if !ok {
		return
	}

	var updateParams UpdateRequest
	if err := c.ShouldBindJSON(&updateParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedRoster, err := h.rosterService.UpdateRoster(id, &updateParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedRoster)
}

// DeleteRoster
//
//	@Summary	DeleteRoster a roster
//	@Security	BearerAuth
//	@Tags		Roster
//	@Accept		json
//	@Produce	json
//	@Param		rosterId	path		int	true	"Roster ID"
//	@Success	200	{string}	string
//	@Failure	400	{string}	string
//	@Failure	404	{string}	string
//	@ID			deleteRoster
//	@Router		/rosters/{rosterId} [delete]
func (h *Handler) DeleteRoster(c *gin.Context) {
	rosterID, ok := httputil.MustPathUint(c, "rosterId")
	if !ok {
		return
	}

	if err := h.rosterService.DeleteRoster(rosterID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Roster deleted",
	})
}
