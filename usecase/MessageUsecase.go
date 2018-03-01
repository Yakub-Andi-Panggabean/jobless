package usecase

import "sync"

type (
	MessageUsecase interface {
		SendMessage(r MessageRequest) (string, error)
		GenerateMessageId(r MessageRequest) string
	}

	message struct {
		MessageStatusV1Repo MessageStatusv1Repository
		MessageStatusRepo   MessageStatusRepository
		MessageLogRepo      MessageLogRepository
		UserRepo            UserRepository
		SenderRepo          SenderRepository
	}
)

var (
	MessageUsecaseInstance MessageUsecase
	messageInstaceOnce     sync.Once
)

/**

create new instance of MessageUsecase interface

*/
func (f *factory) NewMessageUsecase() MessageUsecase {

	messageInstaceOnce.Do(func() {

		MessageUsecaseInstance = message{
			MessageStatusRepo:   f.NewMessageStatusRepository(),
			UserRepo:            f.NewUserRepository(),
			SenderRepo:          f.NewSenderRepository(),
			MessageLogRepo:      f.NewMessageLogRepository(),
			MessageStatusV1Repo: f.NewMessageStatusV1Repository(),
		}

	})

	return MessageUsecaseInstance

}
