package mysql

import "siji/sms-api/actor"

type (
	MessageStatusV1RepoImpl struct {
		actor.UserMessageStatusV1
	}
)

func NewMessageStatusV1RepoImpl() MessageStatusV1RepoImpl {

	var messageStatus MessageStatusV1RepoImpl
	return messageStatus

}

func (m MessageStatusV1RepoImpl) SaveMessage(u actor.UserMessageStatusV1) (int, error) {
	return 0, nil
}
func (m MessageStatusV1RepoImpl) UpdateMessage(u actor.UserMessageStatusV1) (int, error) {
	return 0, nil
}
func (m MessageStatusV1RepoImpl) IsMessageExist(id string) bool {
	return false
}
func (m MessageStatusV1RepoImpl) FindMessage(id string) (*actor.UserMessageStatusV1, error) {
	return nil, nil
}
func (m MessageStatusV1RepoImpl) FindMessages(limit int) []*actor.UserMessageStatusV1 {
	return nil
}
