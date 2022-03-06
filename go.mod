module auth

go 1.17

require (
	github.com/alexedwards/argon2id v0.0.0-20211130144151-3585854a6387
	github.com/coreos/go-oidc/v3 v3.1.0
	github.com/georgysavva/scany v0.3.0
	github.com/gofiber/fiber/v2 v2.28.0
	github.com/gofiber/template v1.6.24
	github.com/jackc/pgx/v4 v4.15.0
	github.com/joho/godotenv v1.4.0
	github.com/ory/hydra-client-go v1.11.7
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.11.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.10.0 // indirect
	github.com/jackc/puddle v1.2.1 // indirect
	github.com/klauspost/compress v1.14.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.33.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
	golang.org/x/net v0.0.0-20220111093109-d55c255bac03 // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)

replace github.com/ory/hydra-client-go v1.11.7 => github.com/monok-o/hydra-client-go v1.11.8-0.20220305235347-de77bfac3384
