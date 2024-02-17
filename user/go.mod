module github.com/ushiradineth/cron-be/user

go 1.22.0

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.6.0
	github.com/jmoiron/sqlx v1.3.5
)

require golang.org/x/crypto v0.19.0 // indirect

require github.com/ushiradineth/cron-be/auth v1.0.0

replace github.com/ushiradineth/cron-be/auth => ../auth
