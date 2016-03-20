package codederror

// 100-199 reserved for general validation errors
const (
	// CodeErrInvalidField is a generic error for an invalid JSON field.
	CodeErrInvalidField = iota + 100

	// CodeErrFieldRequired is an error code for missing JSON fields that are requied.
	CodeErrFieldRequired

	// CodeErrFieldEmpty is an error code for empty JSON fields.
	CodeErrFieldEmpty

	// CodeErrInvalidJWTToken is a generic error for an invalid JWT token.
	CodeErrInvalidJWTToken
)

// 200-299 reserved for homes-api service

// 300-399 reserved for fortress service
const (
	// CodeErrInvalidCredentials is an error code for when the users email and password do not match.
	CodeErrInvalidCredentials = iota + 300

	// CodeErrUnknownEmail is an error code for when the email is not for a registered User.
	CodeErrUnknownEmail

	// CodeErrUnknownID is an error code for when a User's ID is not found in the database.
	CodeErrUnknownID
)

// 400-499 reserved for data-import service

const (
	// CodeErrGeneric is a generic error code used when wrapping standard errors.
	CodeErrGeneric = iota + 900

	// CodeErrUnableToParseRequest is an error code for when a request body cannot be unmarshalled into JSON.
	CodeErrUnableToParseRequest
)
