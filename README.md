# codebase_go_api
- Install go-migrate: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
- Install postgres and create db
```
 make posgres
 make createdb
 make migrateup
```

- Run
```
    make docker_build  // build docker with binary file
    make docker_run    // run docker-compose
```
- Test
```
   curl -d '{"owner": "cathy", "balance": 150, "currency": "EUR"}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/v1/account/add
```
- Update config in file /config/config.yaml

### References
1. [go-database-sql](http://go-database-sql.org/retrieving.html)