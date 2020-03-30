package history

type RabbitLocationData struct {
	Domain             string  `json:"domain"`
	UserId             string  `json:"user_id"`
	ClientTimeStampUTC string  `json:"client_timestamp_utc"`
	ServerTimeStampUTC string  `json:"server_timestamp_utc"`
	Longitude          float64 `json:"lon"`
	Latitude           float64 `json:"lat"`
}
