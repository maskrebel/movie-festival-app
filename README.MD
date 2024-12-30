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
--header 'Authorization: Bearer $token'
```

### Get popular genre movie
```shell
curl --location 'http://localhost:8080/admin/movies/genre-most-viewed' \
--header 'Authorization: Bearer $token'
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
           "watch_url": "https://example.com/movie.mp4"
         }'
```

### Update Movie
```shell
curl --location --request PUT 'http://localhost:8080/admin/movies/1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer $token' \
--data '{
    "title": "Dora Explorer",
    "description": "Dora Explorer Description",
    "duration": 120,
    "artists": [
        "Artist1",
        "Artist2"
    ],
    "genres": [
        "Action",
        "Adventure"
    ],
    "watch_url": "https://example.com/dora.mp4"
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

### Search movie by title/description/artists/genres
```shell
curl --location 'http://localhost:8080/movies/search?q=dora
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
