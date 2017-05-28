package email

// Email represents a very basic email structure
type Email struct {
	To         string //TODO: []string
	From       string
	Subject    string
	Body       string
	Attachment string // The path to the file to be attached, this is also used as the filename
}
