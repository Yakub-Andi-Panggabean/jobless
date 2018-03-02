package mysql

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"siji/sms-api/actor"
	"siji/sms-api/usecase"
	"siji/sms-api/util"
)

type (
	userRepoImpl struct {
		usecase.UserRepository
		connection *sql.DB
	}

	userPersistence struct {
		UserId            sql.NullString
		Version           sql.NullString
		Username          sql.NullString
		Password          sql.NullString
		Active            sql.NullString
		Counter           sql.NullString
		LastAccess        sql.NullString
		SenderIds         sql.NullString
		AuthorizedIPs     sql.NullString
		VirtualNumber     sql.NullString
		Cobrander         sql.NullString
		DeliveryStatusUrl sql.NullString
		UrlInvalidCount   sql.NullString
		UrlActive         sql.NullString
		UrlLastRetry      sql.NullString
		IsUseBlackList    sql.NullString
		IsPostPaidUser    sql.NullString
		InactiveReason    sql.NullString
		TryCount          sql.NullString
		DateTimeTry       sql.NullString
		CreatedDate       sql.NullString
		UpdatedDate       sql.NullString
		CreatedBy         sql.NullString
		UpdatedBy         sql.NullString
	}
)

var log *logrus.Logger

func init() {

	log = logrus.New()

}

func NewUserRepoImpl(db *sql.DB) usecase.UserRepository {

	if db == nil {
		panic("db connection nil detected")
	}

	var userRepo userRepoImpl
	userRepo.connection = db
	return &userRepo

}

func (r userRepoImpl) FindUser(id string) (*actor.SmsApiUser, error) {

	var u actor.SmsApiUser
	var errMessage error

	if r.connection == nil {
		log.Error("nil connection")
	} else {

		row, err := r.connection.Query("select USER_ID,VERSION,USER_NAME,PASSWORD,ACTIVE,COUNTER,LAST_ACCESS,CREATED_DATE,"+
			"UPDATED_DATE,CREATED_BY,UPDATED_BY,COBRANDER_ID,DELIVERY_STATUS_URL,URL_INVALID_COUNT,"+
			"URL_ACTIVE,URL_LAST_RETRY,USE_BLACKLIST,IS_POSTPAID,TRY_COUNT,INACTIVE_REASON,DATETIME_TRY from USER where USER_ID = ?", id)

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

	}

	return &u, errMessage
}

func (r userRepoImpl) FindAuthenticatedUser(username string, password string) (*actor.SmsApiUser, error) {

	log.Info("incoming request username :", username, "password :", password)

	var u actor.SmsApiUser
	var errMessage error

	if r.connection == nil {
		log.Error("nil connection")
	} else {

		row, err := r.connection.Query("select USER_ID,VERSION,USER_NAME,PASSWORD,ACTIVE,COUNTER,LAST_ACCESS,CREATED_DATE,"+
			"UPDATED_DATE,CREATED_BY,UPDATED_BY,COBRANDER_ID,DELIVERY_STATUS_URL,URL_INVALID_COUNT,"+
			"URL_ACTIVE,URL_LAST_RETRY,USE_BLACKLIST,IS_POSTPAID,TRY_COUNT,INACTIVE_REASON,DATETIME_TRY "+
			"from USER where USER_NAME = ? AND PASSWORD=?", username, util.GetMd5Hash(password))

		if err != nil {

			log.Error("an error occured with message ", err.Error())
			errMessage = err

		} else {

			var active sql.NullString
			var urlActive sql.NullString
			var isUseBlackList sql.NullString
			var lastAccess sql.NullString
			var createdDate sql.NullString
			var updatedDate sql.NullString
			var deliveryStatusUrl sql.NullString
			var dateTimeTry sql.NullString
			var urlLastRetry sql.NullString
			var inactiveReason sql.NullString
			var isPostPaid sql.NullString

			if row.Next() {

				err := row.Scan(&u.UserId, &u.Version, &u.Username, &u.Password,
					&active, &u.Counter, &lastAccess, &createdDate,
					&updatedDate, &u.CreatedBy, &u.UpdatedBy, &u.Cobrander,
					&deliveryStatusUrl, &u.UrlInvalidCount, &urlActive,
					&urlLastRetry, &isUseBlackList, &isPostPaid,
					&u.TryCount, &inactiveReason, &dateTimeTry)

				log.Info("active :", active)

				if err != nil {
					log.Error("an error occured with message ", err.Error())
					errMessage = err
				}

			} else {

				log.Error("no record found")

			}

		}

	}

	return &u, errMessage
}

func (u userRepoImpl) FindUsers(limit int) []*actor.SmsApiUser {
	return nil
}
