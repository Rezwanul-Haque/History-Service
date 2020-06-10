package queue

import "fmt"

type RabbitLocationData struct {
	Domain             string  `json:"domain"`
	UserId             string  `json:"user_id"`
	ClientTimeStampUTC string  `json:"client_timestamp_utc"`
	ServerTimeStampUTC string  `json:"server_timestamp_utc"`
	Longitude          float64 `json:"lon"`
	Latitude           float64 `json:"lat"`
}

func (rld RabbitLocationData) ToString() string {
	result := fmt.Sprintf("RabbitLocationData{userId='%s', "+
		"clientTimestampUtc=%s, "+
		"serverTimestampUtc=%s, "+
		"latitude=%f, "+
		"longitude=%f, "+
		"domain='%s'"+
		"}",
		rld.UserId,
		rld.ClientTimeStampUTC,
		rld.ServerTimeStampUTC,
		rld.Latitude,
		rld.Longitude,
		rld.Domain)
	return result
}
