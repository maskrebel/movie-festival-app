# Movie Festival API  

This project is a backend system for managing a **Movie Festival** application. It supports features like movie management, user viewership tracking, voting functionality, and secure authentication.

---

## Contents
1. [Features](#features-)
2. [End Point API](#endpoint-api)
3. [Database Design](#database-design)
4. [Setup Environment](#setup-environment)
5. [Running Tests](#running-tests)
6. [Contributing](#contributing)

## Features  

### Admin Features  
- Add movies  
- Update movie details  
- Retrieve most-viewed movies  

### User Features  
- Register and log in with JWT-based authentication  
- Vote for movies  
- Remove votes (unvote)  
- View movies with view tracking  
- Search movies by title, description, artists, or genres  
- Logout

---

## Endpoint API

## API Endpoints  

### **Admin API**  

#### 1. Add Movie  
- **Endpoint**: `POST /admin/movies` 
- **Headers**: `Authorization: Bearer <JWT_TOKEN>`  
- **Request Body**:  
```json
{
  "title": "Movie Title",
  "description": "Movie Description",
  "duration": 120,
  "artists": ["Artist1", "Artist2"],
  "genres": ["Action", "Adventure"],
  "watch_url": "https://example.com/movie.mp4"
}
```
- **Request Body**:  
```json
{
  "id": 1,
  "message": "Movie created successfully"
}
```

#### 2. Update Movie
- **Endpoint**: `PUT /admin/movies/{id}`
- **Headers**: `Authorization: Bearer <JWT_TOKEN>`
- **Request Body**: Same as Add Movie.
- **Response**:
```json
{
  "message": "Movie updated successfully"
}
```

#### 3. Retrieve Most-Viewed Movies
- **Endpoint**: `GET /admin/movies/most-viewed?page=1&per_page=10`
- **Headers**: `Authorization: Bearer <JWT_TOKEN>`
- **Response**:
```json
{
  "page": 1,
  "per_page": 10,
  "total": 100,
  "total_page": 10,
  "data": [
    {
        "artists": "Artist1, Artist2",
        "created_at": "2024-12-29T17:12:02.071+07:00",
        "description": "Movie Description",
        "duration": 120,
        "genres": "Action, Horor, Romance",
        "id": 1,
        "title": "Movie Title2",
        "updated_at": "2024-12-30T15:37:23.954253+07:00",
        "views": 7,
        "votes": 1,
        "watch_url": "https://example.com/movie-baru2.mp4",
        "year": 2010
    }
  ]
}
```

#### 4. Genre Most-Viewed Movies
- **Endpoint**: `GET /admin/movies/genre-most-viewed`
- **Headers**: `Authorization: Bearer <JWT_TOKEN>`
- **Response**:
```json
{
    "data": {
        "Action": 10,
        "Adventure": 5,
        "Horor": 5,
        "Romance": 5
    }
}
```
---

### **User API**
#### 1. Register  
- **Endpoint**: `POST /auth/register`  
- **Request Body**:  
```json
{
  "username": "user1", 
  "email": "user1@example.com",
  "password": "securepassword"
}
```
- **Response**:
```json
{
  "message": "User registered successfully"
}
```

#### 2. Login  
- **Endpoint**: `POST /auth/login`
- **Request Body**:  
```json
{
  "email": "user1@example.com",
  "password": "securepassword"
}
```
- **Response**:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### 3. Status  
- **Endpoint**: `GET /auth/status`
- **Headers**: `Authorization: Bearer <JWT_TOKEN>`  
- **Response**:
```json
{
    "message": "Welcome user1"
}
```

#### 4. Vote for a Movie
- **Endpoint**: `POST /movies/{id}/vote`
- **Headers**: `Authorization: Bearer <JWT_TOKEN>`
- **Response**:
```json
{
  "message": "Vote added successfully"
}
```

#### 5. Unvote a Movie
- **Endpoint**: `DELETE /movies/{id}/unvote`
- **Headers**: `Authorization: Bearer <JWT_TOKEN>`
- **Response**:
```json
{
  "message": "Vote removed successfully"
}
```

#### 6. View a Movie
- **Endpoint**: `POST /movies/{id}/view?duration=120`
- **Headers**: `Authorization: Bearer <JWT_TOKEN> (optional for guests)`
- **Response**:
```json
{
  "message": "View recorded successfully"
}
```

#### 7. Search movie by title/description/artists/genres
- **Endpoint**: `GET /movies/search?q=title`
- **Response**:
```json
{
  "data": [
    {
      "id": 1,
      "title": "Inception",
      "description": "A mind-bending thriller",
      "artists": ["Leonardo DiCaprio"],
      "genres": ["Sci-Fi", "Thriller"]
    }
  ]
}
```

#### 8. Logout  
- **Endpoint**: `GET /auth/logout`
- **Headers**: `Authorization: Bearer <JWT_TOKEN>`
- **Response**:
```json
{
    "message": "Successfully logged out"
}
```

---

## Database Design
### Users Table
| Column	  | Type	       | Description             |
|----------|-------------|-------------------------|
| id       | Primary Key | 	Unique user identifier |
| username | 	String     | 	User's username        |
| email    | 	String     | 	User's email address   |
| password | 	String     | 	Hashed password        |
| is_admin | 	Bool       | 	default: False         |

### Movies Table
| Column      |	Type| 	Description               |
|-------------|--------|----------------------------|
| id          |	Primary Key| 	Unique movie identifier   |
| title       |	String| 	Movie title               |
| description |	Text| 	Movie description         |
| year        |Integer| year's movie               |
| duration    |	Integer| 	Movie duration in minutes |
| artists     |	Array| 	List of artists           |
| genres      |	Array| 	List of genres            |
| watch_url   |	String| 	Streaming URL             |
| views       |	Integer| 	View count                |
| votes       |	Integer| 	Vote count                |

### Votes Table
|Column|	Type| 	Description |
|------|--------|--------------|
|id	|Primary Key|	Unique vote identifier|
|user_id|	Foreign Key|	References Users table|
|movie_id|	Foreign Key|	References Movies table|

### Views Table
|Column| 	Type                                | 	Description              |
|-|--------------------------------------|---------------------------|
|id| 	Primary Key                         | 	Unique view identifier   |
|user_id| 	Foreign Key                         | 	Nullable, guest views    |
|movie_id| 	Foreign Key	| References Movies table   |
|duration| 	Integer| 	View duration in seconds |

### TokenExpired
|Column|	Type| 	Description |
|------|--------|--------------|
|id|Primary Key| Unique for table|
|token| String| unique for token expired|

---

## Setup Environment
### install dependency

```shell
go mod init
go get github.com/gin-gonic/gin@v1.10.0

go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/dgrijalva/jwt-go

go mod tidy
```
or
```shell
git clone https://github.com/maskrebel/movie-festival-app.git
go install // install from go.mod
```

---
## Running Tests
### To run the tests:

```bash
go test ./tests -v
```
---

### Contributing
**Name**: `Wardi Fadillah`

**Email**: [wardi.fadillah@gmail.com](mailto:wardi.fadillah@gmail.com?subject=Internal%20Wallet%20Transactional%20System)

**Repo**: [GitHub](https://github.com/maskrebel/)