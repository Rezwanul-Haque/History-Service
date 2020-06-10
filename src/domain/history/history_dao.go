package history

import (
	"fmt"
	"github.com/rezwanul-haque/History-Service/src/datasources/mysql/hds_db"
	"github.com/rezwanul-haque/History-Service/src/domain/queue"
	"github.com/rezwanul-haque/History-Service/src/logger"
	"github.com/rezwanul-haque/History-Service/src/utils/errors"
	"time"
)

const (
	queryGetUserByDateRange = "SELECT client_timestamp_utc, server_timestamp_utc, longitude, latitude FROM location_history WHERE LOWER(domain)=LOWER(?) AND user_id=? AND client_timestamp_utc >= ? AND client_timestamp_utc <= ? ORDER BY client_timestamp_utc ASC;"
	queryInsertUserHistory  = "INSERT INTO hds_db.location_history(domain, user_id, client_timestamp_utc, server_timestamp_utc, longitude, latitude) VALUES ((SELECT domain FROM ids_db.company WHERE LOWER(domain)=LOWER(?)),?, ?, ?, ?, ?);"
)

func (hr *HistoryResponse) GetByDateRange(startDate time.Time, endDate time.Time) ([]Path, *errors.RestErr) {
	stmt, err := hds_db.Client.Prepare(queryGetUserByDateRange)
	if err != nil {
		logger.Error("error when trying to prepare get users by companyId and role statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(hr.Domain, hr.UserId, startDate, endDate)
	if err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]Path, 0)

	for rows.Next() {
		var path Path
		if err := rows.Scan(&path.ClientTimeStampUTC, &path.ServerTimeStampUTC, &path.Longitude, &path.Latitude); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, path)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no paths matching with domain %s, user id %v", hr.Domain, hr.UserId))
	}

	return results, nil
}

func Save(l queue.RabbitLocationData) (bool, *errors.RestErr) {
	stmt, err := hds_db.Client.Prepare(queryInsertUserHistory)
	if err != nil {
		logger.Error("error when trying to prepare get users by companyId and role statement", err)
		return false, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(l.Domain, l.UserId, l.ClientTimeStampUTC, l.ServerTimeStampUTC, l.Longitude, l.Latitude)

	if saveErr != nil {
		logger.Error("error when trying to save company", saveErr)
		return false, errors.NewInternalServerError("database error")
	}
	_, err = insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new company", err)
		return false, errors.NewInternalServerError("database error")
	}

	return true, nil
}
