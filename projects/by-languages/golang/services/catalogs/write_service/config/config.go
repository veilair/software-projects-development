package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/grpc"
	customEcho "github.com/mehdihadeli/store-golang-microservice-sample/pkg/http/custom_echo"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/logger"
	postgres "github.com/mehdihadeli/store-golang-microservice-sample/pkg/postgres_pgx"
	"os"

	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/constants"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/eventstroredb"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/gorm_postgres"
	kafkaClient "github.com/mehdihadeli/store-golang-microservice-sample/pkg/kafka"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/probes"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/tracing"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "catalogs write microservice config path")
}

type Config struct {
	DeliveryType     string                         `mapstructure:"deliveryType" env:"DeliveryType"`
	ServiceName      string                         `mapstructure:"serviceName" env:"ServiceName"`
	Logger           *logger.Config                 `mapstructure:"logger" envPrefix:"Logger_"`
	KafkaTopics      KafkaTopics                    `mapstructure:"kafkaTopics" envPrefix:"KafkaTopics_"`
	GRPC             *grpc.GrpcConfig               `mapstructure:"grpc" envPrefix:"GRPC_"`
	Http             *customEcho.EchoHttpConfig     `mapstructure:"http" envPrefix:"Http_"`
	Context          Context                        `mapstructure:"context" envPrefix:"Context_"`
	Postgresql       *postgres.Config               `mapstructure:"postgres" envPrefix:"Postgresql_"`
	GormPostgres     *gorm_postgres.Config          `mapstructure:"gormPostgres" envPrefix:"GormPostgres_"`
	Kafka            *kafkaClient.Config            `mapstructure:"kafka" envPrefix:"Kafka_"`
	Probes           probes.Config                  `mapstructure:"probes" envPrefix:"Probes_"`
	Jaeger           *tracing.Config                `mapstructure:"jaeger" envPrefix:"Jaeger_"`
	EventStoreConfig eventstroredb.EventStoreConfig `mapstructure:"eventStoreConfig" envPrefix:"EventStoreConfig_"`
}

type Context struct {
	Timeout int `mapstructure:"timeout" env:"Timeout"`
}

type KafkaTopics struct {
	ProductCreate  kafkaClient.TopicConfig `mapstructure:"productCreate" envPrefix:"ProductCreate_"`
	ProductCreated kafkaClient.TopicConfig `mapstructure:"productCreated" envPrefix:"ProductCreated_"`
	ProductUpdate  kafkaClient.TopicConfig `mapstructure:"productUpdate" envPrefix:"ProductUpdate_"`
	ProductUpdated kafkaClient.TopicConfig `mapstructure:"productUpdated" envPrefix:"ProductUpdated_"`
	ProductDelete  kafkaClient.TopicConfig `mapstructure:"productDelete" envPrefix:"ProductDelete_"`
	ProductDeleted kafkaClient.TopicConfig `mapstructure:"productDeleted" envPrefix:"ProductDeleted_"`
}

func InitConfig(environment string) (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			configPath = "./config"
		}
	}

	cfg := &Config{}

	viper.SetConfigName(fmt.Sprintf("config.%s", environment))
	viper.AddConfigPath(configPath)
	viper.SetConfigType(constants.Json)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	grpcPort := os.Getenv(constants.GrpcPort)
	if grpcPort != "" {
		cfg.GRPC.Port = grpcPort
	}

	postgresHost := os.Getenv(constants.PostgresqlHost)
	if postgresHost != "" {
		cfg.Postgresql.Host = postgresHost
	}
	postgresPort := os.Getenv(constants.PostgresqlPort)
	if postgresPort != "" {
		cfg.Postgresql.Port = postgresPort
	}
	jaegerAddr := os.Getenv(constants.JaegerHostPort)
	if jaegerAddr != "" {
		cfg.Jaeger.HostPort = jaegerAddr
	}
	kafkaBrokers := os.Getenv(constants.KafkaBrokers)
	if kafkaBrokers != "" {
		cfg.Kafka.Brokers = []string{kafkaBrokers}
	}

	return cfg, nil
}
