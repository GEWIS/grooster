package roster

import (
	"GEWIS-Rooster/internal/models"
	_ "GEWIS-Rooster/internal/models"
	"GEWIS-Rooster/internal/platform/httputil"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	rosterService Service
}

func NewRosterHandler(rosterService Service, rg *gin.RouterGroup, db *gorm.DB) *Handler {
	h := &Handler{rosterService: rosterService}

	// Gin requires wildcards at the same position to share a name, so every
	// route below these prefixes must use the same parameter names.
	organ := rg.Group("/organs/:organId")
	roster := rg.Group("/rosters/:rosterId")
	template := rg.Group("/templates/:templateId")
	shiftGroup := rg.Group("/shift-groups/:shiftGroupId")

	h.registerRosterRoutes(organ, roster, db)
	h.registerShiftRoutes(roster, db)
	h.registerTemplateRoutes(organ, template, db)
	h.registerCommentRoutes(roster, db)

	roster.GET("/responsibles", RequireRosterOrganRoleParam(db, "rosterId", models.RoleMember), h.GetRosterResponsibles)
	roster.POST("/responsibles", RequireRosterOrganRoleParam(db, "rosterId", models.RoleAdmin), h.CreateRosterResponsible)
	roster.DELETE("/responsibles/:userId", RequireRosterOrganRoleParam(db, "rosterId", models.RoleAdmin), h.DeleteRosterResponsible)

	roster.POST("/fill", RequireRosterOrganRoleParam(db, "rosterId", models.RoleAdmin), h.FillRosterPreferences)

	roster.POST("/lock", RequireRosterOrganRoleParam(db, "rosterId", models.RoleAdmin), h.LockRoster)
	roster.GET("/saved-shifts", h.GetSavedShifts)
	roster.PATCH("/saved-shifts/:savedShiftId", requireBelongsTo(db, &models.SavedShift{}, "savedShiftId", "roster_id", "rosterId"), h.UpdateSavedShift)

	organ.POST("/shift-groups", requireOrganRoleParam(db, "organId", models.RoleAdmin), h.CreateShiftGroup)
	organ.GET("/shift-groups", h.GetShiftGroups)
	shiftGroup.GET("", h.GetShiftGroup)

	shiftGroup.GET("/priorities", requireShiftGroupOrganRoleParam(db, "shiftGroupId", models.RoleAdmin), h.GetShiftGroupPriorities)
	shiftGroup.PUT("/priorities", requireShiftGroupOrganRoleParam(db, "shiftGroupId", models.RoleAdmin), h.UpdateShiftGroupPriority)
	shiftGroup.POST("/push-to-bottom", requireShiftGroupOrganRoleParam(db, "shiftGroupId", models.RoleAdmin), h.PushUserToBottom)

	return h
}

// FillRosterPreferences
//
//	@Summary	Fills a roster with the linked user template preferences
//	@Security	BearerAuth
//	@Tags		Roster
//	@Accept		json
//	@Produce	json
//	@Param		rosterId	path		int	true	"Roster ID"
//	@Success	200	{array}	models.RosterAnswer
//	@Failure	400	{string}	string
//	@Failure 404 {string} string
//	@ID			fillRoster
//	@Router		/rosters/{rosterId}/fill [post]
func (h *Handler) FillRosterPreferences(c *gin.Context) {
	rosterID, ok := httputil.MustPathUint(c, "rosterId")
	if !ok {
		return
	}

	answers, err := h.rosterService.FillRosterPreferences(rosterID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, answers)
}

// LockRoster
//
//	@Summary	Lock a specific roster, storing its current assignments as saved shifts
//	@Security	BearerAuth
//	@Tags		Saved Shift
//	@Accept		json
//	@Produce	json
//	@Param		rosterId	path		int	true	"Roster ID"
//	@Success	200	{string}	json
//	@Failure	400	{string}	json
//	@Failure	404	{string}	json
//	@ID			lockRoster
//	@Router		/rosters/{rosterId}/lock [post]
func (h *Handler) LockRoster(c *gin.Context) {
	rosterID, ok := httputil.MustPathUint(c, "rosterId")
	if !ok {
		return
	}

	err := h.rosterService.SaveRoster(rosterID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "SavedShift not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shift"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// UpdateSavedShift
//
//	@Summary	Update a specific saved shift
//	@Security	BearerAuth
//	@Tags		Saved Shift
//	@Accept		json
//	@Produce	json
//	@Param		rosterId		path		int								true	"Roster ID"
//	@Param		savedShiftId	path		int								true	"SavedShift ID"
//	@Param		updateParams	body		SavedShiftUpdateRequest	true	"Update data"
//	@Success	200				{object}	models.SavedShift
//	@Failure	400				{string}	string	"Invalid request"
//	@Failure	404				{string}	string	"SavedShift not found"
//	@ID			updateSavedShift
//	@Router		/rosters/{rosterId}/saved-shifts/{savedShiftId} [patch]
func (h *Handler) UpdateSavedShift(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "savedShiftId")
	if !ok {
		return
	}

	var updateParams SavedShiftUpdateRequest
	if err := c.ShouldBindJSON(&updateParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	updatedShift, err := h.rosterService.UpdateSavedShift(id, &updateParams)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "SavedShift not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shift"})
		}
		return
	}

	c.JSON(http.StatusOK, updatedShift)
}

// GetSavedShifts
//
//	@Summary	Get all saved shifts for a specific roster
//	@Security	BearerAuth
//	@Tags		Saved Shift
//	@Accept		json
//	@Produce	json
//	@Param		rosterId	path		int					true	"Roster ID"
//	@Success	200	{object}		SavedShiftResponse	"Saved Shift Response"
//	@Failure	400	{string}	string				"Invalid request"
//	@Failure	404	{string}	string				"SavedShift not found"
//	@ID			getSavedShifts
//	@Router		/rosters/{rosterId}/saved-shifts [get]
func (h *Handler) GetSavedShifts(c *gin.Context) {
	rosterID, ok := httputil.MustPathUint(c, "rosterId")
	if !ok {
		return
	}

	savedShifts, savedShiftOrdering, err := h.rosterService.GetSavedRoster(rosterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := SavedShiftResponse{
		SavedShifts:        savedShifts,
		SavedShiftOrdering: savedShiftOrdering,
	}
	// Log the entire struct as a field called "response"
	log.Debug().Interface("response", response).Msg("Sending saved roster response")

	c.JSON(http.StatusOK, response)
}

// PushUserToBottom
//
//	@Summary	Manually push a user to the bottom of a shift group's ordering, as if just assigned
//	@Security	BearerAuth
//	@Tags		ShiftGroup
//	@Accept		json
//	@Produce	json
//	@Param		shiftGroupId	path		int					true	"Shift Group ID"
//	@Param		params	body		PushToBottomRequest	true	"User to push to the bottom"
//	@Success	200	{string}	string
//	@Failure	400	{string}	string				"Invalid request"
//	@ID			pushUserToBottom
//	@Router		/shift-groups/{shiftGroupId}/push-to-bottom [post]
func (h *Handler) PushUserToBottom(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "shiftGroupId")
	if !ok {
		return
	}

	var params PushToBottomRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	if err := h.rosterService.PushUserToBottom(id, params.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// CreateShiftGroup
//
//	@Summary   Create a new shift group
//	@Security  BearerAuth
//	@Tags      ShiftGroup
//	@Accept    json
//	@Produce   json
//	@Param     organId path     int                      true "Organ ID"
//	@Param     params body      ShiftGroupCreateRequest  true "Shift Group Details"
//	@Success   201    {object}   models.ShiftGroup
//	@Failure   400    {string}   string
//	@ID        createShiftGroup
//	@Router    /organs/{organId}/shift-groups [post]
func (h *Handler) CreateShiftGroup(c *gin.Context) {
	organID, ok := httputil.MustPathUint(c, "organId")
	if !ok {
		return
	}

	var params ShiftGroupCreateRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group, err := h.rosterService.CreateShiftGroup(organID, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, group)
}

// GetShiftGroups
//
//	@Summary   Get all shift groups for an organ
//	@Security  BearerAuth
//	@Tags      ShiftGroup
//	@Produce   json
//	@Param     organId path     int  true "Organ ID"
//	@Success   200      {array}   models.ShiftGroup
//	@Failure   400      {string}  string
//	@ID        getShiftGroups
//	@Router    /organs/{organId}/shift-groups [get]
func (h *Handler) GetShiftGroups(c *gin.Context) {
	organID, ok := httputil.MustPathUint(c, "organId")
	if !ok {
		return
	}

	groups, err := h.rosterService.GetShiftGroups(organID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// GetShiftGroup
//
//	@Summary   Get a specific shift group by ID
//	@Security  BearerAuth
//	@Tags      ShiftGroup
//	@Produce   json
//	@Param     shiftGroupId path      int  true "Shift Group ID"
//	@Success   200 {object} models.ShiftGroup
//	@Failure   404 {string}   string
//	@ID        getShiftGroup
//	@Router    /shift-groups/{shiftGroupId} [get]
func (h *Handler) GetShiftGroup(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "shiftGroupId")
	if !ok {
		return
	}

	group, err := h.rosterService.GetShiftGroup(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shift group not found"})
		return
	}

	c.JSON(http.StatusOK, group)
}

// GetShiftGroupPriorities
//
//	@Summary	Get a shift group priorities for a shift group
//	@Security	BearerAuth
//	@Tags		ShiftGroup
//	@Accept		json
//	@Produce	json
//	@Param		shiftGroupId	path		int								true	"ShiftGroup ID"
//	@Success	200				{array}	models.ShiftGroupPriority
//	@Failure	400				{string}	string	"Invalid request"
//	@ID			getShiftGroupPriorities
//	@Router    /shift-groups/{shiftGroupId}/priorities [get]
func (h *Handler) GetShiftGroupPriorities(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "shiftGroupId")
	if !ok {
		return
	}

	shiftPriorities, err := h.rosterService.GetShiftGroupPriorities(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shiftPriorities)
}

// UpdateShiftGroupPriority
//
//	@Summary	Update a shift group priority
//	@Security	BearerAuth
//	@Tags		ShiftGroup
//	@Accept		json
//	@Produce	json
//	@Param		shiftGroupId	path		int								true	"ShiftGroup ID"
//	@Param		updateParams	body		GroupPriorityUpdateParam	true	"Update parameters"
//	@Success	200				{object}	models.ShiftGroupPriority
//	@Failure	400				{string}	string	"Invalid request"
//	@Failure	404				{string}	string	"SavedShift not found"
//	@ID			updateShiftGroupPriority
//	@Router    /shift-groups/{shiftGroupId}/priorities [put]
func (h *Handler) UpdateShiftGroupPriority(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "shiftGroupId")
	if !ok {
		return
	}

	var params GroupPriorityUpdateParam
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupPriority, err := h.rosterService.UpdateShiftGroupPriority(id, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groupPriority)
}

// GetRosterResponsibles
//
//	@Summary		Get roster responsibles
//	@Description	Returns the responsibles for a given roster. Callable by any member of the roster's organ.
//	@Security	BearerAuth
//	@Tags		Roster
//	@Accept		json
//	@Produce	json
//	@Param		rosterId	path		int	true	"Roster ID"
//	@Success	200	{array}		models.RosterResponsible
//	@Failure	400	{string}	string	"Invalid roster ID"
//	@Failure	403	{string}	string	"Insufficient organ permissions"
//	@Failure	404	{string}	string	"Roster not found"
//	@ID			getRosterResponsibles
//	@Router		/rosters/{rosterId}/responsibles [get]
func (h *Handler) GetRosterResponsibles(c *gin.Context) {
	rosterID, ok := httputil.MustPathUint(c, "rosterId")
	if !ok {
		return
	}

	responsibles, err := h.rosterService.GetRosterResponsibles(rosterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responsibles)
}

// CreateRosterResponsible
//
//	@Summary		Create a roster responsible
//	@Description	Assigns a user as responsible for a given roster. Requires admin level or higher in the roster's organ.
//	@Security	BearerAuth
//	@Tags		Roster
//	@Accept		json
//	@Produce	json
//	@Param		rosterId	path		int								true	"Roster ID"
//	@Param		params	body		RosterResponsibleCreateRequest	true	"User to assign as responsible"
//	@Success	201		{object}	models.RosterResponsible
//	@Failure	400		{string}	string	"Invalid request"
//	@Failure	403		{string}	string	"Insufficient organ permissions"
//	@Failure	404		{string}	string	"Roster not found"
//	@ID			createRosterResponsible
//	@Router		/rosters/{rosterId}/responsibles [post]
func (h *Handler) CreateRosterResponsible(c *gin.Context) {
	rosterID, ok := httputil.MustPathUint(c, "rosterId")
	if !ok {
		return
	}

	var params RosterResponsibleCreateRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	responsible, err := h.rosterService.CreateRosterResponsible(rosterID, params.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, responsible)
}

// DeleteRosterResponsible
//
//	@Summary		Delete a roster responsible
//	@Description	Removes a user's responsible assignment for a given roster. Requires admin level or higher in the roster's organ.
//	@Security	BearerAuth
//	@Tags		Roster
//	@Accept		json
//	@Produce	json
//	@Param		rosterId	path		int	true	"Roster ID"
//	@Param		userId	path		int	true	"User ID"
//	@Success	200		{object}	map[string]string
//	@Failure	400		{string}	string	"Invalid request"
//	@Failure	403		{string}	string	"Insufficient organ permissions"
//	@Failure	404		{string}	string	"Responsible not found"
//	@ID			deleteRosterResponsible
//	@Router		/rosters/{rosterId}/responsibles/{userId} [delete]
func (h *Handler) DeleteRosterResponsible(c *gin.Context) {
	rosterID, ok := httputil.MustPathUint(c, "rosterId")
	if !ok {
		return
	}

	userID, ok := httputil.MustPathUint(c, "userId")
	if !ok {
		return
	}

	err := h.rosterService.DeleteRosterResponsible(rosterID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Responsible not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Responsible deleted successfully"})
}
