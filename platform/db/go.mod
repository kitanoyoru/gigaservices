module github.com/kitanoyoru/gigaservices/platform/db

go 1.20

replace github.com/kitanoyoru/gigaservices/pkg/grpc => ../../pkg/grpc/

require (
	github.com/go-kivik/kivik/v4 v4.0.0-20230727163647-b793649be555
	github.com/google/uuid v1.3.0
	github.com/kitanoyoru/gigaservices/pkg/grpc v0.0.0-20230729182724-2d138552e06b
	github.com/samber/do v1.6.0
	github.com/samber/lo v1.38.1
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/viper v1.16.0
	golang.org/x/net v0.12.0
	golang.org/x/sys v0.10.0
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
