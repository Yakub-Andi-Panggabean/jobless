package mysql

import (
	"siji/sms-api/usecase"
	"sync"
)

type (
	storageFactory struct {
	}
)

var (
	messageStatusRepositoryInstance   usecase.MessageStatusRepository
	messageStatusRepoOnce             sync.Once
	userRepositoryInstance            usecase.UserRepository
	userRepositoryOnce                sync.Once
	senderRepositoryInstance          usecase.SenderRepository
	senderRepositoryOnce              sync.Once
	messageStatusv1RepositoryInstance usecase.MessageStatusv1Repository
	messageStatusv1RepositoryOnce     sync.Once
	messageLogRepositoryInstance      usecase.MessageLogRepository
	messageLogRepositoryOnce          sync.Once
)

func NewStorage() usecase.StorageFactory {

	return storageFactory{}

}

func (s *storageFactory) NewMessageStatusRepository() usecase.MessageStatusRepository {

	messageStatusRepoOnce.Do(func() {
		messageStatusRepositoryInstance = NewMessageStatusRepoImpl()
	})

	return messageStatusRepositoryInstance
}

func (s *storageFactory) NewUserRepository() usecase.UserRepository {

	userRepositoryOnce.Do(func() {
		userRepositoryInstance = NewUserRepoImpl()
	})

	return userRepositoryInstance
}

func (s *storageFactory) NewSenderRepository() usecase.SenderRepository {

	senderRepositoryOnce.Do(func() {
		senderRepositoryOnce = NewSenderRepoImpl()
	})

	return senderRepositoryInstance
}

func (s *storageFactory) NewMessageStatusV1Repository() usecase.MessageStatusv1Repository {

	messageStatusv1RepositoryOnce.Do(func() {
		messageStatusv1RepositoryOnce = NewMessageStatusV1RepoImpl()
	})

	return messageStatusv1RepositoryInstance
}

func (s *storageFactory) NewMessageLogRepository() usecase.MessageLogRepository {

	messageLogRepositoryOnce.Do(func() {
		messageLogRepositoryOnce = NewMessageLogRepoImpl()
	})

	return messageLogRepositoryInstance
}
