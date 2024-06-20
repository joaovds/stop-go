package helpers

import (
	"net/http"

	"github.com/joaovds/stop-go/internal/presentation/rest/response"
)

// ----- Success -----

func NewNoContent() *response.HttpRes {
	return response.NewHttpResData(http.StatusNoContent, nil)
}

func NewOk(data any) *response.HttpRes {
	return response.NewHttpResData(http.StatusOK, data)
}

func NewCreated(data any) *response.HttpRes {
	return response.NewHttpResData(http.StatusCreated, data)
}

// ----- Error -----

func NewBadRequestError() *response.HttpRes {
	return response.NewHttpResError(http.StatusBadRequest, "bad request")
}

func NewUnauthorizedError() *response.HttpRes {
	return response.NewHttpResError(http.StatusUnauthorized, "unauthorized")
}

func NewForbiddenError() *response.HttpRes {
	return response.NewHttpResError(http.StatusForbidden, "forbidden")
}

func NewNotFoundError() *response.HttpRes {
	return response.NewHttpResError(http.StatusNotFound, "not found")
}

func NewInternalServerError() *response.HttpRes {
	return response.NewHttpResError(http.StatusInternalServerError, "internal server error")
}
