package models

// OperationResponse Serves as the response for the controllers in case of errors or any significative message that does not involves a GET method.
type OperationResponse struct {
	Code        int
	Description interface{}
	Errors      []string
}
