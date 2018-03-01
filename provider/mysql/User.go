package mysql

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"siji/sms-api/actor"
	"siji/sms-api/usecase"
)

type (
	userRepoImpl struct {
		usecase.UserRepository
		connection *sql.DB
	}
)

var log *logrus.Logger

func init() {
	log = logrus.New()
}

func NewUserRepoImpl(db *sql.DB) usecase.UserRepository {

	var userRepo userRepoImpl
	userRepo.connection = db
	return &userRepo

}

func (r userRepoImpl) FindUser(id string) (*actor.SmsApiUser, error) {

	row, err := r.connection.Query("select USER_ID,VERSION,USER_NAME,PASSWORD,ACTIVE,COUNTER,LAST_ACCESS,CREATED_DATE,"+
		"UPDATED_DATE,CREATED_BY,UPDATED_BY,COBRANDER_ID,DELIVERY_STATUS_URL,URL_INVALID_COUNT,"+
		"URL_ACTIVE,URL_LAST_RETRY,USE_BLACKLIST,IS_POSTPAID,TRY_COUNT,INACTIVE_REASON,DATETIME_TRY from SMS_API_USER where USER_ID = ?", id)

	var u actor.SmsApiUser
	var errMessage error

	if err != nil {

		log.Error("an error occured with message ", err.Error())
		errMessage = err

	} else {

		err := row.Scan(&u.UserId, &u.Version, &u.Username, &u.Password,
			&u.Active, &u.Counter, &u.LastAccess, &u.CreatedDate,
			&u.UpdatedDate, &u.CreatedBy, &u.UpdatedBy, &u.Cobrander,
			&u.DeliveryStatusUrl, &u.UrlInvalidCount, &u.UrlActive,
			&u.UrlLastRetry, &u.IsUseBlackList, &u.IsPostPaidUser,
			&u.TryCount, &u.InactiveReason, &u.DateTimeTry)

		if err != nil {
			log.Error("an error occured with message ", err.Error())
			errMessage = err
		}

	}

	return &u, errMessage
}

func (r userRepoImpl) FindAuthenticatedUser(username string, password string) (*actor.SmsApiUser, error) {

	row, err := r.connection.Query("select USER_ID,VERSION,USER_NAME,PASSWORD,ACTIVE,COUNTER,LAST_ACCESS,CREATED_DATE,"+
		"UPDATED_DATE,CREATED_BY,UPDATED_BY,COBRANDER_ID,DELIVERY_STATUS_URL,URL_INVALID_COUNT,"+
		"URL_ACTIVE,URL_LAST_RETRY,USE_BLACKLIST,IS_POSTPAID,TRY_COUNT,INACTIVE_REASON,DATETIME_TRY "+
		"from SMS_API_USER where USER_NAME = ? AND PASSWORD=?", username, password)

	var u actor.SmsApiUser
	var errMessage error

	if err != nil {

		log.Error("an error occured with message ", err.Error())
		errMessage = err

	} else {

		err := row.Scan(&u.UserId, &u.Version, &u.Username, &u.Password,
			&u.Active, &u.Counter, &u.LastAccess, &u.CreatedDate,
			&u.UpdatedDate, &u.CreatedBy, &u.UpdatedBy, &u.Cobrander,
			&u.DeliveryStatusUrl, &u.UrlInvalidCount, &u.UrlActive,
			&u.UrlLastRetry, &u.IsUseBlackList, &u.IsPostPaidUser,
			&u.TryCount, &u.InactiveReason, &u.DateTimeTry)

		if err != nil {
			log.Error("an error occured with message ", err.Error())
			errMessage = err
		}

	}

	return &u, errMessage
}

func (u userRepoImpl) FindUsers(limit int) []*actor.SmsApiUser {
	return nil
}
