package usecase

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"siji/sms-api/actor"
	"siji/sms-api/util"
	"time"
)

var log logrus.Logger

func init() {
	log = logrus.New()
	log.Formatter = &logrus.TextFormatter{}
}

const (
	SMS_API_V1 string = "1"
)

type (
	MessageRequest struct {
		Broadcast      actor.BroadcastSms
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

	sender, errSender := m.SenderRepo.FindSender(m.Request.SenderId)
	user, errUser := m.UserRepo.FindAuthenticatedUser(m.Request.Username, m.Request.Password)

	if errSender != nil {
		err = errSender
	} else {

		//process sms api v1 request
		if m.Request.ApiVersion == SMS_API_V1 {

			transactionId = m.sendSmsV1(sender)

		} else {

			if errUser != nil {

				err = errUser

			} else {

				transactionId = m.sendSmsV2(sender, user)

			}
		}

	}

	return transactionId, err

}

func (m *Message) sendSmsV1(sender actor.Sender) string {

	var senderValue string

	if sender.SenderName != nil {
		senderValue = sender.SenderName
	} else {
		senderValue = sender.SenderId
	}

	statV1 := actor.UserMessageStatusV1{
		Message:       m.Request.MessageContent,
		Destination:   m.Request.Destination,
		Username:      m.Request.Username,
		Type:          m.Request.Type,
		MessageId:     m.GenerateMessageId(),
		SenderName:    senderValue,
		Status:        "", //initial state of message status
		SendTimeStamp: time.Now(),
	}

	messageLog := actor.UserMessageLog{
		Status:         0,
		MessageId:      statV1.MessageId,
		Type:           statV1.Type,
		MessageContent: statV1.Message,
	}

	m.MessageStatusV1Repo.SaveMessage(statV1)
	m.MessageLogRepo.SaveMessageLog(messageLog)

	//after this push to queue based on type

	return statV1.MessageId
}

func (m *Message) sendSmsV2(sender actor.Sender, user actor.SmsApiUser) string {

	statv2 := actor.UserMessageStatus{
		Message:        m.Request.MessageContent,
		Type:           m.Request.Type,
		MessageId:      m.GenerateMessageId(),
		SenderId:       sender.SenderId,
		Destination:    m.Request.Destination,
		Acknowledged:   nil,
		BroadcastSmsId: m.Request.Broadcast,
		MessageStatus:  "",
		SentTime:       time.Now(),
		StatusTime:     nil,
		UserId:         user.Username,
	}

	m.MessageStatusRepo.SaveMessage(statv2)

	//will be put to queue later

	return statv2.MessageId

}

/*

  Generate Message Id

*/
func (m *Message) GenerateMessageId() string {

	var prefix string
	time := time.Now().Format(util.GetConfig().GetString("date.format"))

	var buffer bytes.Buffer

	if m.Request.Type == actor.BINARY {
		prefix = "2GPI"
	} else {
		if m.Request.ApiVersion == SMS_API_V1 {
			prefix = "0GPI"
		} else {
			prefix = util.GetConfig().GetString("transaction.prefix")
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

/*

get sms count

*/
func (m *Message) GetSmsCount() int {

	if m.Request.Type == actor.BINARY {
		return math.Ceil(float64(len(m.Request.MessageContent))+40) / 140
	}

	//is long sms
	if len(m.Request.MessageContent) > 160 {
		return 1
	}

	// is split sms
	if m.Request.IsSplitSms {
		return math.Ceil(float64(len(m.Request.MessageContent) / 160))
	}

	return math.Ceil(float64(len(m.Request.MessageContent) / 153))
}

func (m *Message) IsValidMessage() bool {

	var isValid bool

	if nil != m.Request.Username && m.Request.Username != "" {

		log.Infof("Username is required")

	} else if nil != m.Request.Password && m.Request.Password != "" {

		log.Infof("Password is required")

	} else if nil != m.Request.SenderId && m.Request.SenderId != "" {

		log.Infof("SenderId is required")

	} else if nil != m.Request.Destination && m.Request.Destination != "" {

		log.Infof("Destination is required")

	} else if nil != m.Request.MessageContent && m.Request.MessageContent != "" {

		log.Infof("Message Content is required")

	} else {
		isValid = true
	}

	return isValid

}
