# go-sqlite-bootstrap

Some boilerplate to help bootstrap your next Go project backed by sqlite.

Includes:

* sqlite DB powered by [github.com/mattn/go-sqlite3](github.com/mattn/go-sqlite3)
* Database migrations powered by [github.com/golang-migrate/migrate](github.com/golang-migrate/migrate)
* Database migration files embedded into the binary powered by the stdlib's [embed](https://pkg.go.dev/embed) package
* A function to make it easy during test execution to spin up an ephemeral copy of your database with all migrations applied
