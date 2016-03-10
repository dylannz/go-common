package codederror

// Interface is an basic error that also provided a code and potential additional info.
type Interface interface {
	Error() string
	ErrorCode() int
	AdditionalInfo() string
}
