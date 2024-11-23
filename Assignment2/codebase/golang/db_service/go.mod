module octcarp/sustech/cs328/a2/db

go 1.23

replace octcarp/sustech/cs328/a2/gogrpc v0.1.0 => ../gogrpc

require octcarp/sustech/cs328/a2/gogrpc v0.1.0

require (
	google.golang.org/grpc v1.68.0
	google.golang.org/protobuf v1.35.2
	gorm.io/driver/postgres v1.5.10
	gorm.io/gorm v1.25.12
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
)
