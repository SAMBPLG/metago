package metago

import (
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

func checkForError(resp *resty.Response, err error, errMessage string) error {
	if err != nil {
		return &APIError{
			Code:    0,
			Message: errors.Wrap(err, errMessage).Error(),
			Type:    ParseAPIErrType(err),
		}
	}

	if resp == nil {
		return &APIError{
			Message: "empty response",
			Type:    ParseAPIErrType(err),
		}
	}

	if resp.IsError() {
		// var msg string

		if resp.StatusCode() == http.StatusUnauthorized {
			var err error
			if e, ok := resp.Error().(*MetabaseErr); ok {
				err = e
			}
			return NewInvalidCredentialError("Unauthorized", err)
		}

		// if e, ok := resp.Error().(*HTTPErrorResponse); ok && e.NotEmpty() {
		// 	log.Println("Kesini juga")
		// 	msg = fmt.Sprintf("%s: %s", resp.Status(), e)
		// } else {
		// 	log.Println("Malah kesini juga")
		// 	msg = resp.Status()
		// }

		// return &APIError{
		// 	Code:    resp.StatusCode(),
		// 	Message: msg,
		// 	Type:    ParseAPIErrType(err),
		// }
	}

	return nil
}

// ParseAPIErrType is a convenience method for returning strongly
// typed API errors.
func ParseAPIErrType(err error) APIErrType {
	if err == nil {
		return APIErrTypeUnknown
	}
	switch {
	default:
		return APIErrTypeUnknown
	}
}

func JoinPath(separator string, path ...string) string {
	return strings.Join(path, separator)
}
