package mysql

import "siji/sms-api/actor"

type (
	MessageStatusRepoImpl struct {
		actor.UserMessageStatus
	}
)

func NewMessageStatusRepoImpl() MessageStatusRepoImpl {

	var messageStatus MessageStatusRepoImpl
	return messageStatus

}

func (m MessageStatusRepoImpl) SaveMessage(u actor.UserMessageStatus) (int, error) {
	return 0, nil
}

func (m MessageStatusRepoImpl) UpdateMessage(u actor.UserMessageStatus) (int, error) {
	return 0, nil
}

func (m MessageStatusRepoImpl) IsMessageExist(id string) bool {
	return false
}

func (m MessageStatusRepoImpl) FindMessage(id string) (*actor.UserMessageStatus, error) {
	return nil, nil
}

func (m MessageStatusRepoImpl) FindMessages(limit int) []*actor.UserMessageStatus {
	return nil
}
