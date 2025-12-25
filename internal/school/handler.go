package school

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SchoolHandler struct {
	service *SchoolService
}

func NewSchoolHandler(service *SchoolService) *SchoolHandler {
	return &SchoolHandler{service: service}
}

// POST /schools
func (h *SchoolHandler) Create(c echo.Context) error {

	req := new(School)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	school, err := h.service.Create(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, school)
}

// Get /schools
func (h *SchoolHandler) List(c echo.Context) error {
	schools, err := h.service.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, schools)
}

// Get /schools/:id
func (h *SchoolHandler) Get(c echo.Context) error {
	idParam := strings.Trim(c.Param("id"), "/")
	id, convErr := strconv.Atoi(idParam)
	if convErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid school id"})
	}
	school, err := h.service.Get(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, school)
}

// PUT /schools/:id
func (h *SchoolHandler) Update(c echo.Context) error {
	idParam := strings.Trim(c.Param("id"), "/")
	id, convErr := strconv.Atoi(idParam)
	if convErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid school id"})
	}
	req := make(map[string]any)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	req["id"] = id
	updatedSchool, err := h.service.Update(uint(id), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedSchool)
}

// Delete /schools/:id
func (h *SchoolHandler) Delete(c echo.Context) error {
	idParam := strings.Trim(c.Param("id"), "/")
	id, convErr := strconv.Atoi(idParam)
	if convErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid school id"})
	}
	err := h.service.Delete(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "School not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
