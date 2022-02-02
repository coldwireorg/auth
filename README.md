# Website

## How to use it
### Dev
```
git clone https://codeberg.org/coldwire/website.git
cd website
go mod tidy
go run main.go
```

### Prod
#### Binary
```
git clone https://codeberg.org/coldwire/website.git
cd website
go mod tidy
go build
```

then move the generated binary with the others files (`static`, `views`, `.env`) to your server and just run `./main`

#### Docher
coming soon