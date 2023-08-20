package config

type Config struct {
	Port    int           `yaml:"port"`
	Email   EmailConfig   `yaml:"email"`
	Tracing TracingConfig `yaml:"tracing"`
}

type EmailConfig struct {
	SMTPServer  string `yaml:"smtp_server"`
	SMTPPort    int    `yaml:"smtp_port"`
	SenderEmail string `yaml:"sender_email"`
	AppPassword string `yaml:"app_password"`
}

type TracingConfig struct {
	Enable bool         `yaml:"enable"`
	Jaeger JaegerConfig `yaml:"jaeger"`
}

type JaegerConfig struct {
	URL string `yaml:"url"`
}
