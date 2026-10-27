package roster

import (
	"GEWIS-Rooster/internal/models"
	"GEWIS-Rooster/internal/platform/httputil"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (h *Handler) registerShiftRoutes(roster *gin.RouterGroup, db *gorm.DB) {
	roster.POST("/shifts", RequireRosterOrganRoleParam(db, "rosterId", models.RoleAdmin), h.CreateRosterShift)
	roster.PATCH("/shifts/:shiftId",
		RequireRosterOrganRoleParam(db, "rosterId", models.RoleAdmin),
		requireBelongsTo(db, &models.RosterShift{}, "shiftId", "roster_id", "rosterId"),
		h.UpdateRosterShift)
	roster.DELETE("/shifts/:shiftId",
		RequireRosterOrganRoleParam(db, "rosterId", models.RoleAdmin),
		requireBelongsTo(db, &models.RosterShift{}, "shiftId", "roster_id", "rosterId"),
		h.DeleteRosterShift)

	roster.POST("/answers", RequireRosterOrganRoleParam(db, "rosterId", models.RoleMember), h.CreateRosterAnswer)
	roster.PATCH("/answers/:answerId",
		RequireRosterOrganRoleParam(db, "rosterId", models.RoleMember),
		requireBelongsTo(db, &models.RosterAnswer{}, "answerId", "roster_id", "rosterId"),
		h.UpdateRosterAnswer)
}

// CreateRosterShift
//
//	@Summary	Create a new roster shift
//	@Security	BearerAuth
//	@Tags		Roster Shift
//	@Accept		json
//	@Produce	json
//	@Param		rosterId		path		int					true	"Roster ID"
//	@Param		createParams	body		ShiftCreateRequest	true	"Roster shift input"
//	@Success	200				{object}	models.RosterShift
//	@Failure	400				{string}	string
//	@ID			createRosterShift
//	@Router		/rosters/{rosterId}/shifts [post]
func (h *Handler) CreateRosterShift(c *gin.Context) {
	rosterID, ok := httputil.MustPathUint(c, "rosterId")
	if !ok {
		return
	}

	var param *ShiftCreateRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	createdRosterShift, err := h.rosterService.CreateRosterShift(rosterID, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdRosterShift)
}

// UpdateRosterShift
//
//		@Summary   Update a roster shift
//		@Security  BearerAuth
//		@Tags      Roster Shift
//		@Accept    json
//		@Produce   json
//		@Param     rosterId       path      int                               true   "Roster ID"
//		@Param     shiftId        path      int                               true   "Roster Shift ID"
//		@Param     updateParams   body      ShiftUpdateRequest    true   "Update input"
//		@Success   200            {object}  models.RosterShift
//		@Failure   400            {string}  string
//	 	@Failure 404 {string} string
//		@ID        updateRosterShift
//		@Router    /rosters/{rosterId}/shifts/{shiftId} [patch]
func (h *Handler) UpdateRosterShift(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "shiftId")
	if !ok {
		return
	}

	var param ShiftUpdateRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedShift, err := h.rosterService.UpdateRosterShift(id, &param)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Roster shift not found"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedShift)
}

// DeleteRosterShift
//
//	@Summary	Deletes a roster shift
//	@Security	BearerAuth
//	@Tags		Roster Shift
//	@Accept		json
//	@Produce	json
//	@Param		rosterId	path		int	true	"Roster ID"
//	@Param		shiftId		path		int	true	"Roster Shift ID"
//	@Success	200	{string}	string
//	@Failure	400	{string}	string
//	@ID			deleteRosterShift
//	@Router		/rosters/{rosterId}/shifts/{shiftId} [delete]
func (h *Handler) DeleteRosterShift(c *gin.Context) {
	shiftID, ok := httputil.MustPathUint(c, "shiftId")
	if !ok {
		return
	}

	if err := h.rosterService.DeleteRosterShift(shiftID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Roster deleted",
	})
}

// CreateRosterAnswer
//
//	@Summary	Create a new roster shift answer
//	@Security	BearerAuth
//	@Tags		Roster Answer
//	@Accept		json
//	@Produce	json
//	@Param		rosterId		path		int					true	"Roster ID"
//	@Param		createParams	body		AnswerCreateRequest	true	"Roster answer input"
//	@Success	200				{object}	models.RosterAnswer
//	@Failure	400				{string}	string
//	@ID			createRosterAnswer
//	@Router		/rosters/{rosterId}/answers [post]
func (h *Handler) CreateRosterAnswer(c *gin.Context) {
	rosterID, ok := httputil.MustPathUint(c, "rosterId")
	if !ok {
		return
	}

	var param *AnswerCreateRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request " + err.Error()})
		return
	}

	createdAnswer, err := h.rosterService.CreateRosterAnswer(rosterID, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdAnswer)
}

// UpdateRosterAnswer
//
//	@Summary	Updates a roster answer with the new value
//	@Security	BearerAuth
//	@Tags		Roster Answer
//	@Accept		json
//	@Produce	json
//	@Param		rosterId		path		int									true	"Roster ID"
//	@Param		answerId		path		int									true	"Roster Answer ID"
//	@Param		updateParams	body		AnswerUpdateRequest	true	"New answer value"
//	@Success	200				{object}	models.RosterAnswer
//	@Failure	400				{string}	json
//	@Failure	404				{string}	json
//	@ID			updateRosterAnswer
//	@Router		/rosters/{rosterId}/answers/{answerId} [patch]
func (h *Handler) UpdateRosterAnswer(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "answerId")
	if !ok {
		return
	}

	var updateParams AnswerUpdateRequest
	if err := c.ShouldBindJSON(&updateParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedAnswer, err := h.rosterService.UpdateRosterAnswer(id, &updateParams)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAnswer)
}
