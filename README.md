# Auth

## How to use it
### Dev
```
git clone https://codeberg.org/coldwire/auth.git
cd auth
go mod tidy
go run main.go
```

## Endpoints

---

### Signing key

- GET `/api/pubkey`: Get server's public key for tokens verifications

  ```
  curl -X POST 'https://auth.coldwire.org/api/pubkey'
  ```

---

### User's public key

- GET `/api/pubkey/<username>`: Get a user public key for encryption

  ```
  curl -X POST 'https://auth.coldwire.org/api/pubkey/:username'
  ```


