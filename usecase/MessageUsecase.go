package usecase

import ()
import "sync"

type (
	MessageUsecase interface {
		SendMessage(m MessageRequest) (string, error)
		GenerateMessageId() string
	}

	Message struct {
		Request             MessageRequest
		MessageStatusV1Repo MessageStatusv1Repository
		MessageStatusRepo   MessageStatusRepository
		MessageLogRepo      MessageLogRepository
		UserRepo            UserRepository
		SenderRepo          SenderRepository
	}
)

var (
	messageInstance    Message
	messageInstaceOnce sync.Once
)

func (f *factory) NewMessageStatusRepository() MessageStatusRepository {

	messageInstaceOnce.Do(func() {

		//messageInstance:={

		//}

	})

}
