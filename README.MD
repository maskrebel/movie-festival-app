# Movie Festival App

## setup
### install dependency

```go
go mod init
go get github.com/gin-gonic/gin@v1.10.0

go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/dgrijalva/jwt-go

go mod tidy
```
or
```go
go install // install from go.mod
```

## Endpoint API
### Get popular movie
```shell
curl -X GET "http://localhost:8080/admin/movies/most-viewed" \
     -H "Accept: application/json"
```

### Create list movie
```shell
curl -X POST "http://localhost:8080/admin/movies" \
     -H "Content-Type: application/json" \
     -d '{
           "title": "Movie Title",
           "description": "Movie Description",
           "duration": 120,
           "artists": ["Artist1", "Artist2"],
           "genres": ["Action", "Adventure"],
           "url": "https://example.com/movie.mp4"
         }'
```

### Create new user
```shell
curl -X POST "http://localhost:8080/auth/register" \
     -H "Content-Type: application/json" \
     -d '{
           "username": "user1", 
           "email": "user1@example.com",
           "password": "securepassword"
         }'
```

### User login
```shell
curl -X POST "http://localhost:8080/auth/login" \
     -H "Content-Type: application/json" \
     -d '{
           "email": "user1@example.com",
           "password": "securepassword"
         }'
```

### User vote Movie
```shell
curl --location --request POST 'http://localhost:8080/movies/:movieId/vote' \
--header 'Authorization: Bearer $token'
```

### User UnVote Movie
```shell
curl --location --request DELETE 'http://localhost:8080/movies/:movieId/unvote' \
--header 'Authorization: Bearer $token'
```

### User view record
```shell
curl --location 'http://localhost:8080/movies/1/view' \
--header 'Authorization: Bearer $token' \
--header 'Content-Type: application/json' \
--data '{
    "duration": 120
}'
```
