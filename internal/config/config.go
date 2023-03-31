package config

import (
	"flag"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	GRPCPort int

	// Databases
	PostgresURL string

	// AMQP
	AMQPURL string

	// AWS S3
	S3ID       string
	S3Key      string
	S3Secret   string
	S3Bucket   string
	S3Region   string
	S3Endpoint string

	// Other
	Debug bool
}

func LoadEnv(e string) error {
	err := godotenv.Load(e)
	return err
}

func Parse() Config {

	var cfg = Config{}

	flag.IntVar(&cfg.GRPCPort, "grpc-port", intEnvOrDefault("GRPC_PORT", 9111), "gRPC port")
	flag.BoolVar(&cfg.Debug, "debug", boolEnvOrDefault("DEBUG", false), "debug mode")
	flag.StringVar(&cfg.PostgresURL, "postgres", stringEnvOrDefault("POSTGRES_URL", ""), "postgres database URL")
	flag.StringVar(&cfg.AMQPURL, "amqp-url", stringEnvOrDefault("AMQP_URL", ""), "AMQP server URL")
	flag.StringVar(&cfg.S3ID, "s3id", stringEnvOrDefault("S3ID", ""), "ID used to connect to S3")
	flag.StringVar(&cfg.S3Key, "s3key", stringEnvOrDefault("S3KEY", ""), "Key used to connect to S3")
	flag.StringVar(&cfg.S3Secret, "s3secret", stringEnvOrDefault("S3SECRET", ""), "Secret used to connect to S3")
	flag.StringVar(&cfg.S3Bucket, "s3bucket", stringEnvOrDefault("S3BUCKET", ""), "Bucket used to connect to S3")
	flag.StringVar(&cfg.S3Region, "s3region", stringEnvOrDefault("S3REGION", ""), "Region used to connect to S3")
	flag.StringVar(&cfg.S3Endpoint, "s3endpoint", stringEnvOrDefault("S3ENDPOINT", ""), "Endpoint used to connect to S3")

	flag.Parse()

	return cfg
}

func intEnvOrDefault(k string, d int) int {

	v, errB := os.LookupEnv(k)
	result, err := strconv.Atoi(v)

	if !errB || err != nil {
		return d
	}

	return result
}

func stringEnvOrDefault(k string, d string) string {

	v, errB := os.LookupEnv(k)
	if !errB {
		return d
	}

	return v
}

func boolEnvOrDefault(k string, d bool) bool {

	v, _ := os.LookupEnv(k)
	if v == "true" {
		return true
	} else {
		return false
	}
}
