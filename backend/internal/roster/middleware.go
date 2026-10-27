package roster

import (
	"GEWIS-Rooster/internal/models"
	"GEWIS-Rooster/internal/platform/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

func checkAccess(c *gin.Context, db *gorm.DB, organID interface{}, minRole models.OrganRole) {
	userID, exists := c.Get(middleware.UserIDKey)
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var userOrgan models.UserOrgan

	if err := db.Where("user_id = ? AND organ_id = ?", userID, organID).First(&userOrgan).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You are not a member of this organ"})
		return
	}

	if models.RoleWeights[userOrgan.Role] < models.RoleWeights[minRole] {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient organ permissions"})
		return
	}

	c.Next()
}

// requireOrganRoleParam ensures the current user has the required minimum
// role within the organ referenced by the given path parameter.
func requireOrganRoleParam(db *gorm.DB, paramStr string, minRole models.OrganRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		organID := c.Param(paramStr)
		if organID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": paramStr + " is required"})
			return
		}

		checkAccess(c, db, organID, minRole)
	}
}

// RequireRosterOrganRoleParam validates the existence of a roster by its ID
// from the URL parameters and ensures the current user has the required
// minimum role within that roster's organization.
func RequireRosterOrganRoleParam(db *gorm.DB, paramStr string, minRole models.OrganRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		rosterID := c.Param(paramStr)
		if rosterID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": paramStr + " is required"})
			return
		}

		var roster models.Roster
		if err := db.First(&roster, "id = ?", rosterID).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Roster not found"})
			return
		}

		checkAccess(c, db, roster.OrganID, minRole)
	}
}

// requireTemplateOrganRoleParam validates the existence of a roster template
// by its ID from the URL parameters and ensures the current user has the
// required minimum role within that template's organization.
func requireTemplateOrganRoleParam(db *gorm.DB, paramStr string, minRole models.OrganRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		templateID := c.Param(paramStr)
		if templateID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": paramStr + " is required"})
			return
		}

		var template models.RosterTemplate
		if err := db.First(&template, "id = ?", templateID).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Roster template not found"})
			return
		}

		checkAccess(c, db, template.OrganID, minRole)
	}
}

// requireShiftGroupOrganRoleParam validates the existence of a shift group by
// its ID from the URL parameters and ensures the current user has the
// required minimum role within that shift group's organization.
func requireShiftGroupOrganRoleParam(db *gorm.DB, paramStr string, minRole models.OrganRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupID := c.Param(paramStr)
		if groupID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": paramStr + " is required"})
			return
		}

		var shiftGroup models.ShiftGroup
		if err := db.First(&shiftGroup, "id = ?", groupID).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Shift group not found"})
			return
		}

		checkAccess(c, db, shiftGroup.OrganID, minRole)
	}
}

// requireBelongsTo ensures the resource referenced by childParam belongs to
// the parent referenced by parentParam, so nested routes cannot be used to
// reach resources of another parent. parentColumn is compared against the
// parent ID; joins can be given when that column lives on a related table.
func requireBelongsTo(db *gorm.DB, model any, childParam, parentColumn, parentParam string, joins ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := db.Model(model).
			Where(clause.Eq{Column: clause.Column{Table: clause.CurrentTable, Name: "id"}, Value: c.Param(childParam)}).
			Where(parentColumn+" = ?", c.Param(parentParam))
		for _, join := range joins {
			query = query.Joins(join)
		}

		var count int64
		if err := query.Count(&count).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
			return
		}

		c.Next()
	}
}
