package mysql

import (
	"siji/sms-api/actor"
	"siji/sms-api/usecase"
)

type (
	UserRepoImpl struct {
		usecase.UserRepository
	}
)

func NewUserRepoImpl() UserRepoImpl {

	var userRepo UserRepoImpl
	return userRepo
}

func (u UserRepoImpl) FindUser(id string) (*actor.SmsApiUser, error) {
	return nil, nil
}

func (u UserRepoImpl) FindAuthenticatedUser(username string, password string) (*actor.SmsApiUser, error) {
	return nil, nil
}

func (u UserRepoImpl) FindUsers(limit int) []*actor.SmsApiUser {
	return nil
}
