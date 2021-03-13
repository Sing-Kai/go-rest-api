# Production Ready Go REST API 

Created following Tutorial Edge course

## Running with Docker 

docker-compose up --build

## Testing 

go test ./... -tags=e2e -v

## Running without Docker

### Docker Postgres DB

docker run --name some-postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres

### Running with out Docker Environment Variables

create .env file next to go.mod file

DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_TABLE=postgres
DB_PORT=5432
DB_HOST=localhost