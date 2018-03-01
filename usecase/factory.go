package usecase

type (
	QueuePublisher interface {
		Publish(message string, exchange string, exchangeType string) error
	}

	Factory interface {
		NewMessageUsecase() MessageUsecase
	}

	factory struct {
		StorageFactory
		Publisher QueuePublisher
	}
)

func New(q QueuePublisher, sf StorageFactory) Factory {

	return &factory{
		Publisher:      q,
		StorageFactory: sf,
	}

}
