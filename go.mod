module auth

go 1.16

require (
	codeberg.org/coldwire/cwauth v0.0.0-20220308210551-1fc53c8ab924
	codeberg.org/coldwire/cwhydra v1.0.0
	github.com/BurntSushi/toml v1.0.0
	github.com/alexedwards/argon2id v0.0.0-20211130144151-3585854a6387
	github.com/gofiber/fiber/v2 v2.29.0
	github.com/joho/godotenv v1.4.0
	github.com/rs/zerolog v1.26.1
	golang.org/x/oauth2 v0.0.0-20220309155454-6242fa91716a // indirect
	gorm.io/driver/postgres v1.3.1
	gorm.io/driver/sqlite v1.3.1
	gorm.io/gorm v1.23.2
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/jackc/pgx/v4 v4.15.0 // indirect
	github.com/kataras/jwt v0.1.8
	github.com/mattn/go-sqlite3 v1.14.12 // indirect
	golang.org/x/crypto v0.0.0-20220307211146-efcb8507fb70 // indirect
	golang.org/x/sys v0.0.0-20220310020820-b874c991c1a5 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/ory/hydra-client-go v1.11.7 => github.com/monok-o/hydra-client-go v1.11.8-0.20220305235347-de77bfac3384
