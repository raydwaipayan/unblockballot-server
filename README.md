unblockballot-server
--------------------

Server component for block chain based voting system.

## File structure

```
.
├── docker-compose.yml
├── go.mod
├── go.sum
├── handler
│   ├── admin_handler.go
│   ├── handler.go
│   └── user_handler.go
├── main.go
├── migrations
│   ├── 1_init.go
│   └── migrations.go
├── models
│   └── models.go
├── README.md
├── router
│   └── routes.go
└── types
    ├── db.go
    └── types.go

```

## Running Migrations

```bash
go run migrations/*.go init #wil initialise the schemda
go run migrations/*.go #will create the tables
```

## API-Endpoints

| APIs | METHOD | DESCP |
|---|---|---|
| `/user/register` | POST | Inserting new user to DB |
| `/user/login` | POST | Checking user in DB, generating JWT token |
