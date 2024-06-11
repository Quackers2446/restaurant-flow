# Backend

-   Install Go (1.22)
-   Install dependencies: `go mod download`
-   Run with `go run main.go`

## Helpful Links

-   [Go SQLX Docs](https://pkg.go.dev/github.com/jmoiron/sqlx#section-readme)
-   [Go MySQL Driver](https://github.com/go-sql-driver/mysql/wiki/Examples)
-   [Echo API Framework](https://github.com/labstack/echo?tab=readme-ov-file#example)

## Database Migrations

If you want to make a change to the database, create a new migration that executes some SQL commands (e.g `CREATE TABLE`, `ALTER TABLE`, etc)

-   This allows us to keep existing data while making changes to the schema. Note that the schema will have to be backwards-compatible (and ideally, forwards-compatible too, but this is not absolutely neccesary)
-   All migrations must have an `.up` (i.e applying the migration) and a `.down` (i.e rolling back the migration in case something goes wrong)
    -   For example, for every `CREATE TABLE`, the `.down` file must have a `DROP TABLE`, and for every `ALTER TABLE`, we have to revert the change depending on what was modified.

## Swagger

-   Install Swaggo: `go install github.com/swaggo/swag/cmd/swag@latest`
-   Add `~/go/bin` to your PATH (idk where this is on Windows)
-   Run `swag fmt && swag init` to generate API documentation

API doc comments documentation: [github.com/swaggo/swag](https://github.com/swaggo/swag). Documentation is at `localhost:3333/swagger/`.