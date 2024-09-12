# go-try-clean-arch

## Run app

### start database
```bash
make start-db
```

### start app
```bash
DATABASE_URL='postgresql://postgres:postgres@localhost:5432/expenses-db?sslmode=disable' go run main.go
```
