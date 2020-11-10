// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ComposeLogs defines model for ComposeLogs.
type ComposeLogs struct {
	ImageLogs      []interface{} `json:"image_logs"`
	KojiImportLogs interface{}   `json:"koji_import_logs"`
	KojiInitLogs   interface{}   `json:"koji_init_logs"`
}

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Distribution  string         `json:"distribution"`
	ImageRequests []ImageRequest `json:"image_requests"`
	Koji          Koji           `json:"koji"`
	Name          string         `json:"name"`
	Release       string         `json:"release"`
	Version       string         `json:"version"`
}

// ComposeResponse defines model for ComposeResponse.
type ComposeResponse struct {
	Id          string `json:"id"`
	KojiBuildId int    `json:"koji_build_id"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	ImageStatuses []ImageStatus `json:"image_statuses"`
	KojiBuildId   *int          `json:"koji_build_id,omitempty"`
	KojiTaskId    int           `json:"koji_task_id"`
	Status        string        `json:"status"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {
	Architecture string       `json:"architecture"`
	ImageType    string       `json:"image_type"`
	Repositories []Repository `json:"repositories"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Status string `json:"status"`
}

// Koji defines model for Koji.
type Koji struct {
	Server string `json:"server"`
	TaskId int    `json:"task_id"`
}

// Repository defines model for Repository.
type Repository struct {
	Baseurl string  `json:"baseurl"`
	Gpgkey  *string `json:"gpgkey,omitempty"`
}

// Status defines model for Status.
type Status struct {
	Status string `json:"status"`
}

// PostComposeJSONBody defines parameters for PostCompose.
type PostComposeJSONBody ComposeRequest

// PostComposeRequestBody defines body for PostCompose for application/json ContentType.
type PostComposeJSONRequestBody PostComposeJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create compose
	// (POST /compose)
	PostCompose(ctx echo.Context) error
	// The status of a compose
	// (GET /compose/{id})
	GetComposeId(ctx echo.Context, id string) error
	// Get logs for a compose.
	// (GET /compose/{id}/logs)
	GetComposeIdLogs(ctx echo.Context, id string) error
	// status
	// (GET /status)
	GetStatus(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostCompose converts echo context to params.
func (w *ServerInterfaceWrapper) PostCompose(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostCompose(ctx)
	return err
}

// GetComposeId converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeId(ctx, id)
	return err
}

// GetComposeIdLogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeIdLogs(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeIdLogs(ctx, id)
	return err
}

// GetStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetStatus(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStatus(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST("/compose", wrapper.PostCompose)
	router.GET("/compose/:id", wrapper.GetComposeId)
	router.GET("/compose/:id/logs", wrapper.GetComposeIdLogs)
	router.GET("/status", wrapper.GetStatus)

}
