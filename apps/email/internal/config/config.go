package config

type Config struct {
	Port          int                 `yaml:"port"`
	Email         EmailConfig         `yaml:"email"`
	Database      DatabaseConfig      `yaml:"database"`
	Cache         CacheConfig         `yaml:"cache"`
	MessageBroker MessageBrokerConfig `yaml:"messageBroker"`
	Tracing       TracingConfig       `yaml:"tracing"`
}

type EmailConfig struct {
	SMTPServer  string `yaml:"smtp_server"`
	SMTPPort    int    `yaml:"smtp_port"`
	SenderEmail string `yaml:"senderEmail"`
	AppPassword string `yaml:"appPassword"`
}

type DatabaseConfig struct {
	URL      string `yaml:"url"`
	Password string `yaml:"password"`
}

type CacheConfig struct {
	URL      string `yaml:"url"`
	Password string `yaml:"password"`
}

type MessageBrokerConfig struct {
	BrokersUrl []string `yaml:"brokers"`
}

type TracingConfig struct {
	Enable bool         `yaml:"enable"`
	Jaeger JaegerConfig `yaml:"jaeger"`
}

type JaegerConfig struct {
	URL string `yaml:"url"`
}
