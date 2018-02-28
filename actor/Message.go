package actor

import (
	"time"
)

func init() {

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
		MessageId      string    `json:"messageId"`
		Destination    string    `json:"destination"`
		MessageStatus  string    `json:"messageStatus"`
		SentTime       time.Time `json:"sendDateTime"`
		StatusTime     time.Time `json:"statusDateTime"`
		Message        string    `json:"message"`
		Acknowledged   bool      `json:"acknowledged"`
		BroadcastSmsId int       `json:"broadcastSmsId"`
		UserId         string    `json:"userId"`
		SenderId       string    `json:"senderId"`
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
