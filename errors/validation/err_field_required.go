package validation

import (
	"fmt"

	"github.com/HomesNZ/go-common/errors/codederror"
)

// ErrFieldRequired is raised when a required field is missing or empty.
func ErrFieldRequired(field string) codederror.Interface {
	return codederror.New(fmt.Sprintf("%s is a required field", field), codederror.CodeErrFieldRequired)
}
