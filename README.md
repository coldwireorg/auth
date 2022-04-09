# Auth
## Account manager and Authorization server for coldwire

## Dev

in the "views" folder: 
```sh
npm i
npm run dev
```

in the root folder of ther project:
```sh
docker-compose -f ./hydra.yml up
DEV_FRONT_URL=http://127.0.0.1:8080/ go run main.go -config config.toml
```

then visit http://127.0.0.1:3003 :)