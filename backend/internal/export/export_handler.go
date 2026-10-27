package export

import (
	"GEWIS-Rooster/internal/models"
	"GEWIS-Rooster/internal/platform/httputil"
	"GEWIS-Rooster/internal/roster"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	exportService Service
}

func NewExportHandler(exportService Service, rg *gin.RouterGroup, db *gorm.DB) *Handler {
	h := &Handler{exportService: exportService}

	rg.GET("/rosters/:rosterId/export", roster.RequireRosterOrganRoleParam(db, "rosterId", models.RoleMember), h.AssignmentToPng)

	return h
}

// AssignmentToPng
//
// @Summary      Export roster assignments as PNG
// @Description  Generates and downloads a PNG image containing the shift assignments for a specific roster.
// @Security     BearerAuth
// @Tags         Export
// @Produce      image/png
// @Param        rosterId   path      uint  true  "Roster ID"
// @Success      200  {file}    binary
// @Failure      400  {object}  map[string]string "Invalid ID format"
// @Failure      500  {object}  map[string]string "Internal server error"
// @ID           exportRoster
// @Router       /rosters/{rosterId}/export [get]
func (h *Handler) AssignmentToPng(c *gin.Context) {
	rosterID, ok := httputil.MustPathUint(c, "rosterId")
	if !ok {
		return
	}

	imgBytes, err := h.exportService.AssignmentsToPng(rosterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileName := fmt.Sprintf("roster-%d.png", rosterID)

	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, "image/png", imgBytes)
}
