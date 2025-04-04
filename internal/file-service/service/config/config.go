package config

type Config struct {
	Port       string
	BucketName string
	ChunkSize  int
	Endpoint   string
	AccessKey  string
	SecretKey  string
	UseSSL     bool
}

func NewConfig() *Config {

}
