package organ

import (
	"GEWIS-Rooster/internal/models"
	_ "GEWIS-Rooster/internal/models"
	"GEWIS-Rooster/internal/platform/httputil"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	organService Service
}

func NewOrganHandler(rg *gin.RouterGroup, organService Service, db *gorm.DB) *Handler {
	h := &Handler{organService: organService}

	g := rg.Group("/organs/:organId")

	g.GET("/members", h.GetMembersSettings)
	g.GET("/members/:userId", h.GetMemberSettings)
	g.PATCH("/members/:userId", h.UpdateMemberSettings)

	g.PATCH("/members/:userId/role", requireOrganRoleParam(db, "organId", models.RoleAdmin), h.UpdateMemberRole)

	return h
}

// GetMembersSettings
//
//	@Summary      Get settings for all members within an organ
//	@Security     BearerAuth
//	@Description  Get organ-specific settings like nickname/username for all its members
//	@Tags         Organ
//	@Accept       json
//	@Produce      json
//	@Param        organId        path      uint                                true  "Organ ID"
//	@Success      200            {array}  models.UserOrgan
//	@Failure      400            {string}  string
//	@Failure 404 {string} string
//	@ID	getMembersSettings
//	@Router       /organs/{organId}/members [get]
func (o *Handler) GetMembersSettings(c *gin.Context) {
	organID, ok := httputil.MustPathUint(c, "organId")
	if !ok {
		return
	}

	settings, err := o.organService.GetMembersSettings(organID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Could not find Organ"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, settings)
}

// GetMemberSettings
//
//	@Summary      Get settings for a user within an organ
//	@Security     BearerAuth
//	@Description  Get organ-specific settings like nickname/username
//	@Tags         Organ
//	@Accept       json
//	@Produce      json
//	@Param        organId        path      uint                                true  "Organ ID"
//	@Param        userId         path      uint                                true  "User ID"
//	@Success      200            {object}  models.UserOrgan
//	@Failure      400            {string}  string
//	@Failure 404 {string} string
//	@ID           getMemberSettings
//	@Router       /organs/{organId}/members/{userId} [get]
func (o *Handler) GetMemberSettings(c *gin.Context) {
	organID, ok := httputil.MustPathUint(c, "organId")
	if !ok {
		return
	}

	userID, ok := httputil.MustPathUint(c, "userId")
	if !ok {
		return
	}

	settings, err := o.organService.GetMemberSettings(organID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Could not find Organ or User"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, settings)
}

// UpdateMemberSettings
//
//	@Summary      Update settings for a user within an organ
//	@Security     BearerAuth
//	@Description  Update organ-specific settings like nickname/username
//	@Tags         Organ
//	@Accept       json
//	@Produce      json
//	@Param        organId        path      uint                                true  "Organ ID"
//	@Param        userId         path      uint                                true  "User ID"
//	@Param        updateParams   body      organ.UpdateMemberSettingsParams   true  "Settings input"
//	@Success      200            {object}  models.UserOrgan
//	@Failure      400            {string}  string
//	@ID           updateMemberSettings
//	@Router       /organs/{organId}/members/{userId} [patch]
func (o *Handler) UpdateMemberSettings(c *gin.Context) {
	organID, ok := httputil.MustPathUint(c, "organId")
	if !ok {
		return
	}

	userID, ok := httputil.MustPathUint(c, "userId")
	if !ok {
		return
	}

	var params UpdateMemberSettingsParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := o.organService.UpdateMemberSettings(organID, userID, &params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

// UpdateMemberRole
//
//	@Summary      Update the role of a user within a specific organ
//	@Security     BearerAuth
//	@Description  Update a users role within a specific organ
//	@Tags         Organ
//	@Accept       json
//	@Produce      json
//	@Param        organId        path      uint                                true  "Organ ID"
//	@Param        userId         path      uint                                true  "User ID"
//	@Param        updateParams   body      organ.UpdateMemberRoleParams   true  "Settings input"
//	@Success      200            {object}  models.UserOrgan
//	@Failure      400            {string}  string
//	@Failure 	  404			 {string}  string
//	@ID           updateMemberRole
//	@Router       /organs/{organId}/members/{userId}/role [patch]
func (o *Handler) UpdateMemberRole(c *gin.Context) {
	organID, ok := httputil.MustPathUint(c, "organId")
	if !ok {
		return
	}

	userID, ok := httputil.MustPathUint(c, "userId")
	if !ok {
		return
	}

	var params UpdateMemberRoleParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := o.organService.UpdateMemberRole(organID, userID, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
