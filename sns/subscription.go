package sns

type Subscription struct {
	arn string
}

func (s Subscription) Arn() string {
	return s.arn
}
