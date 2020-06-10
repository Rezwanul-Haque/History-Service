package consts

const (
	APIVersion string = "v1"
)

type RoleDefine struct {
	ADMIN string
	USER  string
}

// Role
var Role = RoleDefine{"admin", "user"}

// QueueType ...
type QueueType string

// RabbitMQ ...
const (
	RabbitMQ QueueType = "rabbitmq"
)
