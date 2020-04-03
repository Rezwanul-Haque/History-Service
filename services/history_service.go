package services

import (
	"github.com/rezwanul-haque/History-Service/domain/history"
	"github.com/rezwanul-haque/History-Service/domain/queue"
	"github.com/rezwanul-haque/History-Service/utils/errors"
	"time"
)

var (
	HistoryService historyServiceInterface = &historyService{}
)

type historyService struct {
}

type historyServiceInterface interface {
	GetHistory(string, string, int64, int64) (*history.HistoryResponse, *errors.RestErr)
	CreateHistory(payload queue.RabbitLocationData) (bool, *errors.RestErr)
}

func (h *historyService) GetHistory(domain string, user_id string, start_date int64, end_date int64) (*history.HistoryResponse, *errors.RestErr) {
	results := &history.HistoryResponse{UserId: user_id, Domain: domain}

	startTime := convertIntToTimeStamp(start_date)
	endTime := convertIntToTimeStamp(end_date)

	paths, getErr := results.GetByDateRange(startTime, endTime)
	results.Paths = paths
	if getErr != nil {
		return nil, getErr
	}
	return results, nil
}

func (h *historyService) CreateHistory(payload queue.RabbitLocationData) (bool, *errors.RestErr) {
	inserted, err := history.Save(payload)
	if err != nil {
		return false, err
	}
	return inserted, nil
}

func convertIntToTimeStamp(date int64) time.Time {
	return time.Unix(0, date*1000*int64(time.Millisecond))
}
