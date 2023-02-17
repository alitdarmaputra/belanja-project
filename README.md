# Belanja Project

### Description
Final Project Studyjam Golang Nest

### Tech Stack

- Go
- MySQL

### How to run?

1. Install Go
2. Get depedencies
    ```
    go get
    ```
3. Copy `.env.example` to `.env` and fill the value according to your own value.
    ```
    cp ./config/.env.example ./config/.env
    ```
    
4. Run Database Migration
   ```
   ./db/migrate -database "mysql://user:password@tcp(host:port)/dbname" -path "./db/migrations" up
   ```
5. Run the app
   
   App can be started with this command:
   ``` 
   go run ./cmd/api
   ```
### Documentation
   https://elements.getpostman.com/redirect?entityId=21970942-d2a24146-dfe0-44d2-b658-b2bc33023e53&entityType=collection
