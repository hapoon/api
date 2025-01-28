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
)

// StatusAPIService is a service that implements the logic for the StatusAPIServicer
// This service should implement the business logic for every endpoint for the StatusAPI API.
// Include any external packages or services that will be required by this service.
type StatusAPIService struct {
}

// NewStatusAPIService creates a default api service
func NewStatusAPIService() *StatusAPIService {
	return &StatusAPIService{}
}

// StatusGet - ステータスを取得します
func (s *StatusAPIService) StatusGet(ctx context.Context) (ImplResponse, error) {
	// TODO - update StatusGet with the required logic for this service method.
	// Add api_status_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, StatusGet200Response{}) or use other options such as http.Ok ...
	return Response(200, StatusGet200Response{Status: "仕事中"}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	//return Response(http.StatusNotImplemented, nil), errors.New("StatusGet method not implemented")
}
