package mysql

import "siji/sms-api/actor"

type (
	MessageLogRepoImpl struct {
		actor.UserMessageLog
	}
)

func NewMessageLogRepoImpl() MessageLogRepoImpl {

	var messageLog MessageLogRepoImpl
	return messageLog

}

func (m MessageLogRepoImpl) SaveMessageLog(u actor.UserMessageLog) (int, error) {
	return 0, nil
}

func (m MessageLogRepoImpl) UpdateMessageLog(u actor.UserMessageLog) (int, error) {
	return 0, nil
}

func (m MessageLogRepoImpl) DeleteMessageLog(u actor.UserMessageLog) (int, error) {
	return 0, nil
}

func (m MessageLogRepoImpl) IsMessageLogExist(id string) bool {
	return false
}

func (m MessageLogRepoImpl) FindMessageLog(id string) (*actor.UserMessageLog, error) {
	return nil, nil
}

func (m MessageLogRepoImpl) FindMessagesLog(limit int) []*actor.UserMessageLog {
	return nil
}
