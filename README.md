# Helpdesk Service Users

#### Helpdesk is test project of helpdesk page using microservices

## Helpdesk Service Users

- Build with go 1.20.2
- Uses the [Protocol Buffers](https://protobuf.dev/)
- Uses [gRPC](https://grpc.io/)
- Uses the [jackc pgx](https://github.com/jackc/pgx) PostgreSQL driver and toolkit for Go to communicate with Postgres db
- Uses PostgreSQL 15.2

-------------
### Helpdesk has below services:

- [Proto](https://github.com/dzwiedz90/helpdesk-proto)
- [service-agents]() - to manage agents
- [service-frontend](https://github.com/dzwiedz90/helpdesk-service-frontend) - to serve frontend UI of application and communicate with other services when necessary
- [service-notifications]() - to manage and send notification about events for tickets
- [service-tickets]() - to manage tickets
- [service-users](https://github.com/dzwiedz90/helpdesk-service-users) - to manage users

-------------
### Configuration before first run
- git pull origin master
- set up Postgres database
- create .env file and fill it with information as in the example below which will be loaded to the app's config:
```
GRPCAddress=0.0.0.0
GRPCPort=5002
Timeout=5
DB_HOST=localhost
DB_NAME=helpdesk-users
DB_USER=postgres
DB_PASSWORD=postgres
DB_PORT=5432
```
- run app with command ```go run main.go```
-------------
-------------