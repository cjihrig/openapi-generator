/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// StoreApiController binds http requests to an api service and writes the service results to the http response
type StoreApiController struct {
	service StoreApiServicer
	errorHandler ErrorHandler
}

// StoreApiOption for how the controller is set up.
type StoreApiOption func(*StoreApiController)

// WithStoreApiErrorHandler inject ErrorHandler into controller
func WithStoreApiErrorHandler(h ErrorHandler) StoreApiOption {
	return func(c *StoreApiController) {
		c.errorHandler = h
	}
}

// NewStoreApiController creates a default api controller
func NewStoreApiController(s StoreApiServicer, opts ...StoreApiOption) Router {
	controller := &StoreApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the StoreApiController
func (c *StoreApiController) Routes() Routes {
	return Routes{ 
		{
			"DeleteOrder",
			strings.ToUpper("Delete"),
			"/v2/store/order/{orderId}",
			c.DeleteOrder,
		},
		{
			"GetInventory",
			strings.ToUpper("Get"),
			"/v2/store/inventory",
			c.GetInventory,
		},
		{
			"GetOrderById",
			strings.ToUpper("Get"),
			"/v2/store/order/{orderId}",
			c.GetOrderById,
		},
		{
			"PlaceOrder",
			strings.ToUpper("Post"),
			"/v2/store/order",
			c.PlaceOrder,
		},
	}
}

// DeleteOrder - Delete purchase order by ID
func (c *StoreApiController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "orderId")
	
	result, err := c.service.DeleteOrder(r.Context(), orderId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetInventory - Returns pet inventories by status
func (c *StoreApiController) GetInventory(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetInventory(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetOrderById - Find purchase order by ID
func (c *StoreApiController) GetOrderById(w http.ResponseWriter, r *http.Request) {
	orderId, err := parseInt64Parameter(chi.URLParam(r, "orderId"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetOrderById(r.Context(), orderId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// PlaceOrder - Place an order for a pet
func (c *StoreApiController) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	order := &Order{}
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.PlaceOrder(r.Context(), *order)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}
