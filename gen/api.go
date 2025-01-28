// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Status
 *
 * ステータスを管理するAPI定義です。
 *
 * API version: 0.1.0
 */

package api

import (
	"context"
	"net/http"
)



// StatusAPIRouter defines the required methods for binding the api requests to a responses for the StatusAPI
// The StatusAPIRouter implementation should parse necessary information from the http request,
// pass the data to a StatusAPIServicer to perform the required actions, then write the service results to the http response.
type StatusAPIRouter interface { 
	StatusGet(http.ResponseWriter, *http.Request)
}


// StatusAPIServicer defines the api actions for the StatusAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type StatusAPIServicer interface { 
	StatusGet(context.Context) (ImplResponse, error)
}
