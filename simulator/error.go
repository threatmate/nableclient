package simulator

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/tekkamanendless/httperror"
	"github.com/threatmate/restfulwrapper"
)

// APIError is an error that represents an API error.
type APIError struct {
	code    int
	message string
}

var _ error = (*APIError)(nil)
var _ restfulwrapper.ErrorWriter = (*APIError)(nil)

func (e *APIError) Error() string {
	return e.message
}

func (e *APIError) WriteError(resp *restful.Response) {
	output := ErrorResponse{
		Status:  e.code,
		Message: e.message,
	}
	resp.WriteHeaderAndEntity(e.code, output)
}

func (e *APIError) Unwrap() []error {
	return []error{httperror.ErrorFromStatus(e.code)}
}
