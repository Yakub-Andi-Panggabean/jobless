package mysql

import (
	"database/sql"
	"siji/sms-api/actor"
	"siji/sms-api/usecase"
)

type (
	senderRepoImpl struct {
		usecase.SenderRepository
		connection *sql.DB
	}
)

func NewSenderRepoImpl(db *sql.DB) usecase.SenderRepository {

	var senderRepo senderRepoImpl
	senderRepo.connection = db
	return senderRepo

}

func (u senderRepoImpl) UpdateSender(s actor.Sender) (int, error) {
	return 0, nil
}
func (u senderRepoImpl) FindSender(id string) (*actor.Sender, error) {
	return nil, nil
}
func (u senderRepoImpl) FindSenders(limit int) []*actor.Sender {
	return nil
}
