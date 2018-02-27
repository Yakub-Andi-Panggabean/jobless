package usecase

import (
	"bytes"
	"siji/sms-api/actor"
	"math/rand"
	"time"
)

const (
	SMS_API_V1       string = "1"
	DATE_TIME_FORMAT        = "2006-01-02 15:04:05.00"
)

type (
	MessageRequest struct {
		Username       string
		Password       string
		SenderId       string
		Destination    string
		ApiVersion     string //to differentiate sms api v2 and v1
		IPOrigin       string // only for sms api v1
		Id             string
		Type           actor.MessageType
		IsSplitSms     bool
		MessageContent string
	}
)

func (m *Message) SendMessage() (string, error) {

	var err error
	var transactionId string

	if m.Request.ApiVersion == SMS_API_V1 {

		statV1 := actor.UserMessageStatusV1{
			Message:     m.Request.MessageContent,
			Destination: m.Request.Destination,
			Username:    m.Request.Username,
			Type:        m.Request.Type,
			MessageId:m.GenerateMessageId(),
		}

		messageLog := actor.UserMessageLog{}

		m.MessageStatusV1Repo.SaveMessage(statV1)
		m.MessageLogRepo.SaveMessageLog(messageLog)

	} else {

	}

	return transactionId, err

}

/*

  Generate Message Id

*/
func (m *Message) GenerateMessageId() string {

	var prefix string
	time := time.Now().Format(DATE_TIME_FORMAT)

	var buffer bytes.Buffer

	if m.Request.Type == actor.BINARY {
		prefix = "2GPI"
	} else {
		if m.Request.ApiVersion == SMS_API_V1 {
			prefix = "0GPI"
		} else {
			prefix = "5GPI"
		}
	}

	buffer.Write(prefix)
	buffer.Write(time)
	buffer.Write(".")
	buffer.Write(randSeq(5))

	return buffer.String()

}

/*

   generate random alphanumeric

*/
func randSeq(n int) string {

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)

	for i := range b {

		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
