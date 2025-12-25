package school

import "go_poc/pkg/generic"

type SchoolHandler struct {
	*generic.BaseHandler[School]
}

func NewSchoolHandler(service *SchoolService) *SchoolHandler {
	baseHandl := generic.NewBaseHandle(service.BaseService)
	return &SchoolHandler{BaseHandler: baseHandl}
}
