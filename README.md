# Go web api
- Gin + Postgres
- Graceful shutdown
- Config: read from a file and environment variables  
- Middleware: rate limit, authentication: jwt
- Unit test: test real db and mock test
- Logger: customer, write logs to file
- Database: implement transaction
- Run app with minimal docker image 
- CI: Github action

### 1. Install 
- Install go-migrate: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
 
- Install postgres and create db
```
 make posgres
 make createdb
 make migrateup
```

### 2. Run
```
    make docker_build  // build docker with binary file
    make docker_run    // run docker-compose
```
### 3. Test
```
   curl -d '{"owner": "cathy", "balance": 150, "currency": "EUR"}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/v1/account/add
```
- Update config in file app/config/config.yaml

### 4. Other
- How use mock generate code:
  ```
    go get github.com/golang/mock/mockgen@v1.5.0
  ```
  - Add go/bin folder to $PATH environment variable: edit PATH in ~/.zshrc or ~/.bashrc
  ```
  vi ~/.zshrc  
  export PATH=$PATH:~/go/bin # Add line
  source ~/.zshrc
  which mockgen
  --> /home/$USER/go/bin/mockgen
  ```
  
### References
1. [simplebank](https://github.com/techschool/simplebank)
2. [gin-admin](https://github.com/LyricTian/gin-admin)