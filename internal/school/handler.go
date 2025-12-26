package school

import (
	"go_poc/pkg/generic"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SchoolHandler struct {
	*generic.BaseHandler[School]
}

func NewSchoolHandler(service *SchoolService) *SchoolHandler {
	baseHandler := generic.NewBaseHandler(service.BaseService)
	return &SchoolHandler{BaseHandler: baseHandler}
}

type ClassroomHandler struct {
	srv         *ClassroomService
	baseHandler *generic.BaseHandler[Classroom]
}

func NewClassroomHandler(service *ClassroomService) *ClassroomHandler {
	baseHandler := generic.NewBaseHandler(service.BaseService)
	return &ClassroomHandler{srv: service, baseHandler: baseHandler}
}

func (h *ClassroomHandler) Create(c echo.Context) error {
	req := new(Classroom)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	entity, err := h.srv.Create(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, entity)
}

func (h *ClassroomHandler) List(c echo.Context) error {
	schools, err := h.srv.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, schools)
}

func (h *ClassroomHandler) Get(c echo.Context) error {
	idParam := c.Param("id")
	id, convErr := strconv.Atoi(idParam)
	if convErr != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "Invalid classroom id"},
		)
	}
	entity, err := h.srv.Get(uint(id))
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
	}
	return c.JSON(http.StatusOK, entity)
}

func (h *ClassroomHandler) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, convErr := strconv.Atoi(idParam)
	if convErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid classroom id"})
	}

	req := make(map[string]any)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	entity, err := h.srv.Update(uint(id), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, entity)
}

func (h *ClassroomHandler) Delete(c echo.Context) error {
	return h.baseHandler.Delete(c)
}
