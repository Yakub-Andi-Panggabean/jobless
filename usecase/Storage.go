package usecase

import "firswap/sms-api/actor"

type (
	MessageStatusRepository interface {
		SaveMessage(m actor.UserMessageStatus) (int, error)
		UpdateMessage(m actor.UserMessageStatus) (int, error)
		DeleteMessage(m actor.UserMessageStatus) (int, error)
		IsMessageExist(id string) bool
		FindMessage(id string) (*actor.UserMessageStatus, error)
		FindMessages(limit int) []*actor.UserMessageStatus
	}

	UserRepository interface {
		SaveUser(u actor.SmsApiUser) (int, error)
		UpdateUser(u actor.SmsApiUser) (int, error)
		DeleteUser(u actor.SmsApiUser) (int, error)
		FindUser(id string) (*actor.SmsApiUser, error)
		FindUsers(limit int) []*actor.SmsApiUser
	}

	SenderRepository interface {
		SaveSender(s actor.Sender) (int, error)
		UpdateSender(s actor.Sender) (int, error)
		DeleteSender(s actor.Sender) (int, error)
		FindSender(id string) (*actor.Sender, error)
		FindSenders(limit int) []*actor.Sender
	}

	MessageStatusv1Repository interface {
		SaveMessage(m actor.UserMessageStatusV1) (int, error)
		UpdateMessage(m actor.UserMessageStatusV1) (int, error)
		DeleteMessage(m actor.UserMessageStatusV1) (int, error)
		IsMessageExist(id string) bool
		FindMessage(id string) (*actor.UserMessageStatusV1, error)
		FindMessages(limit int) []*actor.UserMessageStatusV1
	}

	MessageLogRepository interface {
		SaveMessageLog(m actor.UserMessageLog) (int, error)
		UpdateMessageLog(m actor.UserMessageLog) (int, error)
		DeleteMessageLog(m actor.UserMessageLog) (int, error)
		IsMessageLogExist(id string) bool
		FindMessageLog(id string) (*actor.UserMessageLog, error)
		FindMessagesLog(limit int) []*actor.UserMessageLog
	}
)