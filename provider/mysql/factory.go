package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/modules/db/app"
	"siji/sms-api/usecase"
	"siji/sms-api/util"
	"strconv"
	"sync"
)

type (
	storageFactory struct {
		db *sql.DB
	}
)

var (
	storage                           storageFactory
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

	driver := util.GetConfig().GetString("db.driver")
	username := util.GetConfig().GetString("db.user")
	password := util.GetConfig().GetString("db.password")
	dbname := util.GetConfig().GetString("db.name")
	dbport := util.GetConfig().GetInt("db.port")
	dbaddress := util.GetConfig().GetString("db.address")

	db, err := sql.Open(driver, username+":"+password+"@tcp("+dbaddress+":"+strconv.Itoa(dbport)+")/"+dbname)

	if err != nil {

		panic("an error when trying to connect to database :" + err.Error())

	}

	storage.db = db

	return storage

}

func (s storageFactory) NewMessageStatusRepository() usecase.MessageStatusRepository {

	messageStatusRepoOnce.Do(func() {
		messageStatusRepositoryInstance = NewMessageStatusRepoImpl()
	})

	return messageStatusRepositoryInstance
}

func (s storageFactory) NewUserRepository() usecase.UserRepository {

	userRepositoryOnce.Do(func() {
		userRepositoryInstance = NewUserRepoImpl(db.Db)
	})

	return userRepositoryInstance
}

func (s storageFactory) NewSenderRepository() usecase.SenderRepository {

	senderRepositoryOnce.Do(func() {
		senderRepositoryInstance = NewSenderRepoImpl()
	})

	return senderRepositoryInstance
}

func (s storageFactory) NewMessageStatusV1Repository() usecase.MessageStatusv1Repository {

	messageStatusv1RepositoryOnce.Do(func() {
		messageStatusv1RepositoryInstance = NewMessageStatusV1RepoImpl()
	})

	return messageStatusv1RepositoryInstance
}

func (s storageFactory) NewMessageLogRepository() usecase.MessageLogRepository {

	messageLogRepositoryOnce.Do(func() {
		messageLogRepositoryInstance = NewMessageLogRepoImpl()
	})

	return messageLogRepositoryInstance
}
