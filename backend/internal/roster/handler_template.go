package roster

import (
	"GEWIS-Rooster/internal/models"
	"GEWIS-Rooster/internal/platform/httputil"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (h *Handler) registerTemplateRoutes(organ *gin.RouterGroup, template *gin.RouterGroup, db *gorm.DB) {
	organ.POST("/templates", requireOrganRoleParam(db, "organId", models.RoleAdmin), h.CreateRosterTemplate)
	organ.GET("/templates", h.GetRosterTemplates)

	template.GET("", h.GetRosterTemplate)
	template.PUT("", requireTemplateOrganRoleParam(db, "templateId", models.RoleAdmin), h.UpdateRosterTemplate)
	template.DELETE("", requireTemplateOrganRoleParam(db, "templateId", models.RoleAdmin), h.DeleteRosterTemplate)

	template.PATCH("/shifts/:shiftId",
		requireBelongsTo(db, &models.RosterTemplateShift{}, "shiftId", "template_id", "templateId"),
		h.UpdateRosterTemplateShift)

	template.POST("/preferences", h.CreateRosterTemplateShiftPreference)
	template.GET("/preferences", h.GetRosterTemplateShiftPreferences)
	template.PATCH("/preferences/:preferenceId",
		requireBelongsTo(db, &models.RosterTemplateShiftPreference{}, "preferenceId", "roster_template_shifts.template_id", "templateId",
			"JOIN roster_template_shifts ON roster_template_shifts.id = roster_template_shift_preferences.roster_template_shift_id"),
		h.UpdateRosterTemplateShiftPreference)
}

// CreateRosterTemplate
//
//	@Summary	Creates a template of a roster by defining the name of the shifts
//	@Security	BearerAuth
//	@Tags		Roster Template
//	@Accept		json
//	@Produce	json
//	@Param		organId	path		int						true	"Organ ID"
//	@Param		params	body		TemplateCreateRequest					false	"Template Params"
//	@Success	200	{array}		models.RosterTemplate	"Created Template"
//	@Failure	400	{string}	string				"Invalid request"
//	@ID			createRosterTemplate
//	@Router		/organs/{organId}/templates [post]
func (h *Handler) CreateRosterTemplate(c *gin.Context) {
	organID, ok := httputil.MustPathUint(c, "organId")
	if !ok {
		return
	}

	var param TemplateCreateRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request" + err.Error()})
		return
	}

	createdTemplate, err := h.rosterService.CreateRosterTemplate(organID, &param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTemplate)
}

// GetRosterTemplates
//
//	@Summary	Get all roster templates of an organ
//	@Security	BearerAuth
//	@Tags		Roster Template
//	@Accept		json
//	@Produce	json
//	@Param		organId	path		int	true	"Organ ID"
//	@Success	200		{array}		models.RosterTemplate
//	@Failure	400		{string}	string
//	@ID			getRosterTemplates
//	@Router		/organs/{organId}/templates [get]
func (h *Handler) GetRosterTemplates(c *gin.Context) {
	organID, ok := httputil.MustPathUint(c, "organId")
	if !ok {
		return
	}

	rosterTemplates, err := h.rosterService.GetRosterTemplates(organID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, rosterTemplates)
}

// GetRosterTemplate
//
//	@Summary	Get a roster template by ID
//	@Security	BearerAuth
//	@Tags		Roster Template
//	@Accept		json
//	@Produce	json
//	@Param		templateId	path		int	true	"Template ID"
//	@Success	200		{object}		models.RosterTemplate
//	@Failure 	404 	{string} 	string
//	@Failure	400		{string}	string
//	@ID			getRosterTemplate
//	@Router		/templates/{templateId} [get]
func (h *Handler) GetRosterTemplate(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "templateId")
	if !ok {
		return
	}

	rosterTemplate, err := h.rosterService.GetRosterTemplate(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, rosterTemplate)
}

// UpdateRosterTemplate
//
//	@Summary	Updates a roster template by ID
//	@Security	BearerAuth
//	@Tags		Roster Template
//	@Accept		json
//	@Produce	json
//	@Param		templateId	path		int	true	"Template ID"
//	@Param		params	body		TemplateUpdateParams	false "Update params"
//	@Success	200		{object}		models.RosterTemplate
//	@Failure	400		{string}	string
//	@Failure 	404 	{string} 	string
//	@ID			updateRosterTemplate
//	@Router		/templates/{templateId} [put]
func (h *Handler) UpdateRosterTemplate(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "templateId")
	if !ok {
		return
	}

	var updateParams TemplateUpdateParams
	if err := c.ShouldBindJSON(&updateParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	updatedTemplate, err := h.rosterService.UpdateRosterTemplate(id, &updateParams)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Roster template not found"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTemplate)
}

// DeleteRosterTemplate
//
//	@Summary	Deletes a roster template by ID
//	@Security	BearerAuth
//	@Tags		Roster Template
//	@Accept		json
//	@Produce	json
//	@Param		templateId	path		int	true	"Template ID"
//	@Success	200		{string}	string
//	@Failure 	404 	{string} 	string
//	@Failure	400		{string}	string
//	@ID			deleteRosterTemplate
//	@Router		/templates/{templateId} [delete]
func (h *Handler) DeleteRosterTemplate(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "templateId")
	if !ok {
		return
	}

	err := h.rosterService.DeleteRosterTemplate(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

// UpdateRosterTemplateShift
//
//	@Summary   Updates a roster template shift by ID
//	@Security  BearerAuth
//	@Tags      Roster Template
//	@Accept    json
//	@Produce   json
//	@Param     templateId path   int    true   "Template ID"
//	@Param     shiftId    path   int    true   "Template Shift ID"
//	@Param     params body      TemplateShiftUpdateRequest  true "Update params"
//	@Success   200    {object}   models.RosterTemplateShift
//	@Failure   400    {string}   string
//	@Failure   404    {string}   string
//	@ID        updateRosterTemplateShift
//	@Router    /templates/{templateId}/shifts/{shiftId} [patch]
func (h *Handler) UpdateRosterTemplateShift(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "shiftId")
	if !ok {
		return
	}

	var updateParams TemplateShiftUpdateRequest
	if err := c.ShouldBindJSON(&updateParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	shift, err := h.rosterService.UpdateRosterTemplateShift(id, &updateParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shift)
}

// CreateRosterTemplateShiftPreference
//
//	@Summary   Creates a roster template shift preference
//	@Security  BearerAuth
//	@Tags      Roster Template
//	@Accept    json
//	@Produce   json
//	@Param     templateId path   int    true   "Template ID"
//	@Param     params body      TemplateShiftPreferenceCreateRequest  true "Creation params"
//	@Success   201       {object}      models.RosterTemplateShiftPreference
//	@Failure   400       {string}   string
//	@ID        createRosterTemplateShiftPreference
//	@Router    /templates/{templateId}/preferences [post]
func (h *Handler) CreateRosterTemplateShiftPreference(c *gin.Context) {
	templateID, ok := httputil.MustPathUint(c, "templateId")
	if !ok {
		return
	}

	var params TemplateShiftPreferenceCreateRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	preference, err := h.rosterService.CreateRosterTemplateShiftPreference(templateID, params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, preference)
}

// GetRosterTemplateShiftPreferences
//
//	@Summary   Gets shift preferences of a template filtered by user
//	@Security  BearerAuth
//	@Tags      Roster Template
//	@Accept    json
//	@Produce   json
//	@Param     templateId  path     int     true  "Template ID"
//	@Param     userId      query    int     true  "User ID"
//	@Success   200         {array}  models.RosterTemplateShiftPreference
//	@Failure   400         {string} string
//	@ID        getRosterTemplateShiftPreferences
//	@Router    /templates/{templateId}/preferences [get]
func (h *Handler) GetRosterTemplateShiftPreferences(c *gin.Context) {
	templateID, ok := httputil.MustPathUint(c, "templateId")
	if !ok {
		return
	}

	var params TemplateShiftPreferenceFilterParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters: " + err.Error()})
		return
	}

	preferences, err := h.rosterService.GetRosterTemplateShiftPreferences(templateID, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch preferences: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, preferences)
}

// UpdateRosterTemplateShiftPreference
//
//	@Summary   Updates a roster template shift preference by ID
//	@Security  BearerAuth
//	@Tags      Roster Template
//	@Accept    json
//	@Produce   json
//	@Param     templateId   path   int    true   "Template ID"
//	@Param     preferenceId path   int    true   "Preference ID"
//	@Param     params body      TemplateShiftPreferenceUpdateRequest  true "Update params"
//	@Success   200       {object}      models.RosterTemplateShiftPreference
//	@Failure   400       {string}   string
//	@Failure   404       {string}   string
//	@ID        updateRosterTemplateShiftPreference
//	@Router    /templates/{templateId}/preferences/{preferenceId} [patch]
func (h *Handler) UpdateRosterTemplateShiftPreference(c *gin.Context) {
	id, ok := httputil.MustPathUint(c, "preferenceId")
	if !ok {
		return
	}

	var params TemplateShiftPreferenceUpdateRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	updatedPreference, err := h.rosterService.UpdateRosterTemplateShiftPreference(id, params)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Shift preference not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedPreference)
}
