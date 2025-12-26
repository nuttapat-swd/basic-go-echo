package generic

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BaseHandler[T any] struct {
	service *BaseService[T]
}

func NewBaseHandler[T any](service *BaseService[T]) *BaseHandler[T] {
	return &BaseHandler[T]{service: service}
}

func (h *BaseHandler[T]) GetTypeName() string {
	name := reflect.TypeOf((*T)(nil)).Elem().Name()
	return name
}

func (h *BaseHandler[T]) Create(c echo.Context) error {

	req := new(T)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	entity, err := h.service.Create(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, entity)
}

func (h *BaseHandler[T]) List(c echo.Context) error {
	entities, err := h.service.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, entities)
}

func (h *BaseHandler[T]) Get(c echo.Context) error {
	idParan := strings.Trim(c.Param("id"), "/")
	id, convErr := strconv.Atoi(idParan)
	if convErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid " + strings.ToLower(h.GetTypeName()) + " id"})
	}
	entity, err := h.service.Get(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, entity)
}

func (h *BaseHandler[T]) Update(c echo.Context) error {
	idParan := strings.Trim(c.Param("id"), "/")
	id, convErr := strconv.Atoi(idParan)
	if convErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid " + strings.ToLower(h.GetTypeName()) + " id"})
	}

	req := make(map[string]any)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	req["id"] = id
	entity, err := h.service.Update(uint(id), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, entity)
}

func (h *BaseHandler[T]) Delete(c echo.Context) error {
	idParan := strings.Trim(c.Param("id"), "/")
	id, convErr := strconv.Atoi(idParan)
	if convErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid " + strings.ToLower(h.GetTypeName()) + " id"})
	}

	err := h.service.Delete(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": strings.Title(strings.ToLower(h.GetTypeName())) + " not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
