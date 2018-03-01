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

var log *logrus.Logger

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

func (m message) SendMessage(r MessageRequest) (string, error) {

	var err error
	var transactionId string

	user, errUser := m.UserRepo.FindAuthenticatedUser(r.Username, r.Password)
	sender, errSender := m.SenderRepo.FindSender(r.SenderId)

	if errSender != nil {
		err = errSender
	} else {

		//process sms api v1 request
		if r.ApiVersion == SMS_API_V1 {

			transactionId = m.sendSmsV1(*sender, r)

		} else {

			if errUser != nil {

				err = errUser

			} else {

				transactionId = m.sendSmsV2(*sender, *user, r)

			}
		}

	}

	return transactionId, err

}

func (m message) sendSmsV1(sender actor.Sender, r MessageRequest) string {

	var senderValue string

	if sender.SenderName != "" {
		senderValue = sender.SenderName
	} else {
		senderValue = sender.SenderId
	}

	statV1 := actor.UserMessageStatusV1{
		Message:       r.MessageContent,
		Destination:   r.Destination,
		Username:      r.Username,
		Type:          r.Type,
		MessageId:     m.GenerateMessageId(r),
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

func (m message) sendSmsV2(sender actor.Sender, user actor.SmsApiUser, r MessageRequest) string {

	statv2 := actor.UserMessageStatus{
		Message:        r.MessageContent,
		Type:           r.Type,
		MessageId:      m.GenerateMessageId(r),
		SenderId:       sender.SenderId,
		Destination:    r.Destination,
		Acknowledged:   false,
		BroadcastSmsId: r.Broadcast.BroadcastSmsId,
		MessageStatus:  "",
		SentTime:       time.Now(),
		UserId:         user.Username,
	}

	m.MessageStatusRepo.SaveMessage(statv2)

	//will be put to queue later

	return statv2.MessageId

}

/*

  Generate Message Id

*/
func (m message) GenerateMessageId(r MessageRequest) string {

	var prefix string
	time := time.Now().Format(util.GetConfig().GetString("date.format"))

	var buffer bytes.Buffer

	if r.Type == actor.BINARY {
		prefix = "2GPI"
	} else {
		if r.ApiVersion == SMS_API_V1 {
			prefix = "0GPI"
		} else {
			prefix = util.GetConfig().GetString("transaction.prefix")
		}
	}

	buffer.WriteString(prefix)
	buffer.WriteString(time)
	buffer.WriteString(".")
	buffer.WriteString(randSeq(5))

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
func (m message) GetSmsCount(r MessageRequest) int {

	if r.Type == actor.BINARY {
		return int(math.Ceil(float64(len(r.MessageContent))+40) / 140)
	} else if len(r.MessageContent) > 160 { //is long sms
		return 1
	} else if r.IsSplitSms { // is split sms
		return int(math.Ceil(float64(len(r.MessageContent) / 160)))
	} else {
		return int(math.Ceil(float64(len(r.MessageContent) / 153)))
	}

}

func (m message) IsValidMessage(r MessageRequest) bool {

	var isValid bool

	if r.Username != "" {

		log.Infof("Username is required")

	} else if r.Password != "" {

		log.Infof("Password is required")

	} else if r.SenderId != "" {

		log.Infof("SenderId is required")

	} else if r.Destination != "" {

		log.Infof("Destination is required")

	} else if r.MessageContent != "" {

		log.Infof("Message Content is required")

	} else {
		isValid = true
	}

	return isValid

}
