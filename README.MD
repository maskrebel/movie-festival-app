# Movie Festival App

## setup
### install dependency

```go
go get github.com/gin-gonic/gin@v1.10.0

go get gorm.io/gorm
go get gorm.io/driver/postgres
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
