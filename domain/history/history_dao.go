package history

import (
	"fmt"
	"github.com/rezwanul-haque/History-Service/datasources/mysql/hds_db"
	"github.com/rezwanul-haque/History-Service/logger"
	"github.com/rezwanul-haque/History-Service/utils/errors"
	"time"
)

const (
	queryGetUserByDateRange = "SELECT client_timestamp_utc, server_timestamp_utc, longitude, latitude FROM location_history WHERE LOWER(domain)=LOWER(?) AND user_id=? AND client_timestamp_utc >= ? AND client_timestamp_utc <= ? ORDER BY client_timestamp_utc ASC;"
)

func (hr *HistoryResponse) GetByDateRange(start_date time.Time, end_date time.Time) ([]Path, *errors.RestErr) {
	stmt, err := hds_db.Client.Prepare(queryGetUserByDateRange)
	if err != nil {
		logger.Error("error when trying to prepare get users by companyId and role statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(hr.Domain, hr.UserId, start_date, end_date)
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
