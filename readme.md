# user-api

# how to run

## db migration

for db migration, i use:

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

```
migrate.exe -path .\database\migrations\mysql -database mysql mysql://kenny:kenny@localhost:3306/user_api up
```

## set envar

```shell
export USER_API_MYSQL_HOST="localhost"
export USER_API_MYSQL_PASSWORD="kenny"
export USER_API_MYSQL_USERNAME="kenny"
export USER_API_MYSQL_SINGULAR="false"
export USER_API_MYSQL_DB_NAME="user_api"
```

exec

```shell
set -o allexport; source .env; set +o allexport
```
# API Contracts

