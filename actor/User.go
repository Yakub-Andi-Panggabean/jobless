package actor

import (
	//"strconv"
	"github.com/sirupsen/logrus"
	"time"
)

var log logrus.Logger

func init() {

	log = logrus.New()

}

type (
	Sender struct {
		SenderId      string
		SenderEnabled bool
		UserId        string
		SenderName    string
		RangeStart    int64
		RangeEnd      int64
		Cobrander     string
	}

	SmsApiUser struct {
		Actor
		UserId            int
		Version           int
		Username          string
		Password          string
		Active            bool
		Counter           int
		LastAccess        time.Time
		SenderIds         []string
		AuthorizedIPs     []string
		VirtualNumber     []string
		Cobrander         []string
		DeliveryStatusUrl string
		UrlInvalidCount   int
		UrlActive         bool
		UrlLastRetry      time.Time
		IsUseBlackList    bool
		IsPostPaidUser    bool
		InactiveReason    bool
		TryCount          int
		DateTimeTry       time.Time
	}
)

/*

 check valid ip

*/
func (u SmsApiUser) IsValidIp(ip string) bool {

	var isExist bool = false

	for index, value := range u.AuthorizedIPs {

		if value == ip {

			log.Debug("valid ip found at index :", index)
			isExist = true

		}

	}

	return isExist

}

/*

 get sender id

*/
/*func (u SmsApiUser) GetSenderId(sender string) SenderId {

	var convertedSenderId int64
	var result SenderId

	converted, err := strconv.ParseInt(sender, 10, 64)

	if nil == err {

		convertedSenderId = converted

	} else {
		convertedSenderId = 0
	}

	for _, value := range u.SenderIds {

		if sender == value {

			if convertedSenderId > 0 {

				value.SenderId = "+" + convertedSenderId

			}

			result = value

		} else if convertedSenderId > 0 && nil != value.RangeStart && nil != value.RangeEnd && convertedSenderId >= value.RangeStart && convertedSenderId <= value.RangeEnd {

			value.SenderId = "+" + convertedSenderId

			result = value

		}

	}

	return result

}*/
