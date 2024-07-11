# test-task-go
## Deploy
```
cp .env.example .env
docker compose up -d
```

### Run migrations
```
docker compose exec -it app sh
sh ./scripts/migrate.sh
```

### Generate Api Key
```
go run ./scripts/generatekey.go
```

## Swagger
http://localhost:8080/docs/index.html

