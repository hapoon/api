// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Status
 *
 * ステータスを管理するAPI定義です。
 *
 * API version: 0.1.0
 */

package api




type StatusGet200Response struct {

	Status string `json:"status,omitempty"`
}

// AssertStatusGet200ResponseRequired checks if the required fields are not zero-ed
func AssertStatusGet200ResponseRequired(obj StatusGet200Response) error {
	return nil
}

// AssertStatusGet200ResponseConstraints checks if the values respects the defined constraints
func AssertStatusGet200ResponseConstraints(obj StatusGet200Response) error {
	return nil
}
