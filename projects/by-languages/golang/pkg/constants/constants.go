package constants

import "time"

const (
	GrpcPort       = "GRPC_PORT"
	HttpPort       = "HTTP_PORT"
	ConfigPath     = "CONFIG_PATH"
	KafkaBrokers   = "KAFKA_BROKERS"
	JaegerHostPort = "JAEGER_HOST"
	RedisAddr      = "REDIS_ADDR"
	MongoDbURI     = "MONGO_URI"
	PostgresqlHost = "POSTGRES_HOST"
	PostgresqlPort = "POSTGRES_PORT"

	ReaderServicePort = "READER_SERVICE"

	Yaml          = "yaml"
	Json          = "json"
	Tcp           = "tcp"
	Redis         = "redis"
	Kafka         = "kafka"
	Postgres      = "postgres"
	MongoDB       = "mongo"
	ElasticSearch = "elasticSearch"

	GRPC     = "GRPC"
	SIZE     = "SIZE"
	URI      = "URI"
	STATUS   = "STATUS"
	HTTP     = "HTTP"
	ERROR    = "ERROR"
	METHOD   = "METHOD"
	NAME     = "NAME"
	METADATA = "METADATA"
	REQUEST  = "REQUEST"
	REPLY    = "REPLY"
	TIME     = "TIME"

	Topic     = "topic"
	Partition = "partition"
	Message   = "message"
	WorkerID  = "workerID"
	Offset    = "offset"
	Time      = "time"

	Validate        = "validate"
	FieldValidation = "field validation"
	RequiredHeaders = "required header"
	Base64          = "base64"
	Unmarshal       = "unmarshal"
	Uuid            = "uuid"
	Cookie          = "cookie"
	Token           = "token"
	Bcrypt          = "bcrypt"
	SQLState        = "sqlstate"

	Page   = "page"
	Size   = "size"
	Search = "search"
	ID     = "id"
)

const (
	ErrBadRequest          = "Bad Request"
	ErrConflict            = "Conflict Error"
	ErrEmailAlreadyExists  = "User with given email already exists"
	ErrNoSuchUser          = "User NotFound"
	ErrWrongCredentials    = "Wrong Credentials"
	ErrNotFound            = "Not Found"
	ErrUnauthorized        = "Unauthorized"
	ErrForbidden           = "Forbidden"
	ErrBadQueryParams      = "Invalid query params"
	ErrRequestTimeout      = "Request Timeout"
	ErrInvalidEmail        = "Invalid Email"
	ErrInvalidPassword     = "Invalid Password"
	ErrInvalidField        = "Invalid Field"
	ErrInternalServerError = "Internal Server Error"
	ErrDomain              = "Domain Model Error"
	ErrApplication         = "Application Service Error"
	ErrApi                 = "Api Error"
)

const (
	MaxHeaderBytes       = 1 << 20
	StackSize            = 1 << 10 // 1 KB
	BodyLimit            = "2M"
	ReadTimeout          = 15 * time.Second
	WriteTimeout         = 15 * time.Second
	GzipLevel            = 5
	WaitShotDownDuration = 3 * time.Second
	Dev                  = "development"
	Production           = "production"
)
