# user-api

# how to run

## db migration

for db migration, i use:

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

```
migrate.exe -source file://database/migrations/mysql -database "mysql://kenny:kenny@tcp(localhost:3306)/user_api" up
migrate.exe -source file://database/migrations/mysql -database "mysql://kenny:kenny@tcp(localhost:3306)/user_api" down
```

## set envar

```shell
export USER_API_MYSQL_HOST="localhost"
export USER_API_MYSQL_PASSWORD="kenny"
export USER_API_MYSQL_USERNAME="kenny"
export USER_API_MYSQL_SINGULAR_TABLE="false"
export USER_API_MYSQL_DB_NAME="user_api"
export ACCESS_TOKEN_SECRET="kenny"
export REFRESH_TOKEN_SECRET="kenny_juga"
export USER_API_REDIS_HOST=localhost
export USER_API_REDIS_PORT=6379
```

exec

```shell
set -o allexport; source .env; set +o allexport
```
# API Contracts
