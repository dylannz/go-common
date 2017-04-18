package sns

import (
	"encoding/json"

	"github.com/HomesNZ/go-common/env"
	"github.com/aws/aws-sdk-go/service/sns"
)

func NewTopic(name string) (*Topic, error) {
	topic := &Topic{
		name: name,
	}
	err := topic.Create()
	if err != nil {
		return nil, err
	}
	return topic, nil
}

// Topic represents an SNS topic
type Topic struct {
	name string
	arn  string
}

// Name returns the name of the topic including environment
func (t Topic) Name() *string {
	suffix := env.Env()
	if suffix == "" {
		suffix = "development"
	}
	name := t.name + "_" + suffix
	return &name
}

// Create attempts to create the SNS topic. CreateTopic() SNS function is
// idempotent, if the topic exists then the existing ARN will be returned.
func (t *Topic) Create() error {
	if !snsEnabled() {
		return nil
	}
	input := &sns.CreateTopicInput{
		Name: t.Name(),
	}
	createTopicOutput, err := Conn().CreateTopic(input)
	if err != nil {
		return err
	}
	t.arn = *createTopicOutput.TopicArn
	return nil
}

func (t *Topic) AddSQSSubscription(sqsArn string) (*Subscription, error) {
	if !snsEnabled() {
		return nil, nil
	}
	protocol := "sqs"
	arn := t.arn
	subscribeOutput, err := Conn().Subscribe(&sns.SubscribeInput{
		Endpoint: &sqsArn,
		Protocol: &protocol,
		TopicArn: &arn,
	})
	if err != nil {
		return nil, err
	}
	return &Subscription{arn: *subscribeOutput.SubscriptionArn}, nil
}

// PushMessage encodes a messageObj into JSON and pushes it to the SNS topic.
func (t Topic) PushMessage(messageObj interface{}) (string, error) {
	if !snsEnabled() {
		return "", nil
	}
	messageObjBytes, err := json.Marshal(messageObj)
	if err != nil {
		return "", err
	}
	messageWrapper := Message{string(messageObjBytes)}
	messageBytes, err := json.Marshal(messageWrapper)
	if err != nil {
		return "", err
	}
	messageStructure := "json"
	message := string(messageBytes)
	topicArn := t.arn
	publishInput := &sns.PublishInput{
		MessageStructure: &messageStructure,
		Message:          &message,
		TopicArn:         &topicArn,
		// MessageAttributes: ,
		// Subject: ,
		// TargetArn: ,
	}
	contextLogger.Debug("publishInput:", *publishInput)
	publishOutput, err := Conn().Publish(publishInput)
	if err != nil {
		return "", err
	}
	contextLogger.Debug("publishOutput:", *publishOutput)
	return *publishOutput.MessageId, nil
}
