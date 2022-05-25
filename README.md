# go-pg-stat

App require enabled [pg_stat_statements extension](https://www.postgresql.org/docs/14/pgstatstatements.html).

You can [get sample PostgreSQL database here](https://www.postgresqltutorial.com/postgresql-getting-started/load-postgresql-sample-database/).

Build and run:

> go build && ./go-pg-stat

By default, app will be working here:

> http://127.0.0.1:13013/api/stat/get

For proper params list check swagger.

Check test coverage:

> go test ./...  \
-coverpkg=./app/... \
-coverprofile ./coverage.out && go tool cover -func ./coverage.out