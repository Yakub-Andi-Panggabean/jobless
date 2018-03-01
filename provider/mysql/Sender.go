package mysql

import (
	"siji/sms-api/actor"
	"siji/sms-api/usecase"
)

type (
	SenderRepoImpl struct {
		usecase.SenderRepository
	}
)

func NewSenderRepoImpl() SenderRepoImpl {

	var senderRepo SenderRepoImpl
	return senderRepo

}

func (u SenderRepoImpl) UpdateSender(s actor.Sender) (int, error) {
	return 0, nil
}
func (u SenderRepoImpl) FindSender(id string) (*actor.Sender, error) {
	return nil
}
func (u SenderRepoImpl) FindSenders(limit int) []*actor.Sender {
	return nil
}
