package history

import "github.com/rezwanul-haque/History-Service/src/utils/errors"

type History struct {
	Id                 int64  `json:"id"`
	Domain             string `json:"domain"`
	UserId             string `json:"user_id"`
	ClientTimeStampUTC string `json:"client_timestamp_utc"`
	ServerTimeStampUTC string `json:"server_timestamp_utc"`
	Longitude          string `json:"lon"`
	Latitude           string `json:"lat"`
}

type QueryParamRequest struct {
	UserId    string `form:"user_id" binding:"required"`
	StartDate *int64 `form:"start_date" binding:"required"`
	EndDate   *int64 `form:"end_date" binding:"required"`
}

func (hr *QueryParamRequest) Validate() *errors.RestErr {
	if hr.UserId == "" {
		return errors.NewBadRequestError("Required string parameter 'user_id' is not present")
	}
	if hr.StartDate == nil {
		return errors.NewBadRequestError("Required long 'start_date' parameter is not present")
	}
	if hr.EndDate == nil {
		return errors.NewBadRequestError("Required long 'end_date' parameter is not present")
	}
	return nil
}

type Path struct {
	ClientTimeStampUTC string `json:"client_timestamp_utc"`
	ServerTimeStampUTC string `json:"server_timestamp_utc"`
	Longitude          string `json:"lon"`
	Latitude           string `json:"lat"`
}

type Paths []Path

type HistoryResponse struct {
	Domain string `json:"domain"`
	UserId string `json:"user_id"`
	Paths  Paths  `json:"paths"`
}
