package actor

import (
	"github.com/sirupsen/logrus"
	"math"
	"time"
	"siji/sms-api/usecase"
)

var log logrus.Logger

func init() {

	log = logrus.New()

}

type MessageType int8

const (
	LATIN   MessageType = 0
	UNICODE MessageType = 1
	BINARY  MessageType = 2
)

type (


	UserMessageStatus struct {
		Actor
		MessageId        string
		Destination      string
		MessageStatus    string
		SentTime         time.Time
		StatusTime       time.Time
		Message          string
		Type             MessageType
		Acknowledged     bool
		BroadcastSmsId   int
		UserId           string
		SenderId         string
		ErrorCode        string
		ErrorDescription string
		Index            int
	}

	UserMessageStatusV1 struct {
		Actor
		MessageId        string
		Username         string
		SenderName       string
		Destination      string
		SendTimeStamp    time.Time
		Status           string
		StatusTimeStamp  time.Time
		Type             MessageType
		Message          string
		ErrorCode        string
		ErrorDescription string
	}

	UserMessageLog struct {
		Id             int
		MessageId      string
		MessageContent string
		Type           MessageType
		Status         int
	}

	BroadcastSms struct {
		BroadcastSmsId int
		TimeStamp      time.Time
		SmsCount       int
	}

	PushStatusFlag struct {
		UserMessageStatusId int
		ProceedTimeStamp    time.Time
	}

	OperatorDialPrefix struct {
		OpCountryCode    string
		OpId             string
		OpDialRangeLower string
		OpDialRangeUpper string
		LongRangeLower   int64
		LongRangeUpper   int64
	}
)

/*

get sms count

*/
func (m usecase.MessageRequest) GetSmsCount() int {

	if m.Type == BINARY {
		return math.Ceil(float64(len(m.MessageContent))+40) / 140
	}

	//is long sms
	if len(m.MessageContent) > 160 {
		return 1
	}

	// is split sms
	if m.IsSplitSms {
		return math.Ceil(float64(len(m.MessageContent) / 160))
	}

	return math.Ceil(float64(len(m.MessageContent) / 153))
}

func (m usecase.MessageRequest) IsValidMessage() bool {

	var isValid bool

	if nil != m.Username && m.Username != "" {

		log.Infof("Username is required")

	} else if nil != m.Password && m.Password != "" {

		log.Infof("Password is required")

	} else if nil != m.SenderId && m.SenderId != "" {

		log.Infof("SenderId is required")

	} else if nil != m.Destination && m.Destination != "" {

		log.Infof("Destination is required")

	} else if nil != m.MessageContent && m.MessageContent != "" {

		log.Infof("Message Content is required")

	} else {
		isValid = true
	}

	return isValid

}
