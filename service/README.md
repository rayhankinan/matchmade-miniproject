# Backend Service

This is a backend service that provides a REST API for a simple movie watchlist application.

## Contents

The contents of the README are as follows:

- [Backend Service](#backend-service)
  - [Contents](#contents)
  - [Stack](#stack)
  - [File Structure](#file-structure)
  - [API Endpoints](#api-endpoints)
    - [Register](#register)
    - [Login](#login)
    - [Logout](#logout)
    - [Add Movie](#add-movie)
    - [Remove Movie](#remove-movie)
    - [Get Watchlist](#get-watchlist)
    - [Rate Movie](#rate-movie)
    - [Check is Movie in Watchlist](#check-is-movie-in-watchlist)
    - [Get Current Rating of Movie](#get-current-rating-of-movie)

## Stack

The service is built using the following stack:
- Go
- Echo
- GORM
- PostgreSQL
- Docker

## File Structure

The file structure of the service is as follows:

```
service
├── internal
│   ├── app
│   │   ├── start.go
│   ├── cmd
│   │   ├── cmd.go
│   ├── config
│   │   ├── environment.go
│   ├── handlers
│   │   ├── auth.go
│   │   ├── watchlist.go
│   ├── infrastructure
│   │   ├── db.go
│   │   ├── router.go
│   ├── middleware
│   │   ├── middleware.go
│   ├── migrations
│   │   ├── migrate.go
│   ├── models
│   │   ├── movies.go
│   │   ├── users.go
│   ├── repositories
│   │   ├── mocks
│   │   │   ├── movie
│   │   │   │   ├── movierepo.go
│   │   │   ├── user
│   │   │   │   ├── userrepo.go
│   │   ├── movierepo.go
│   │   ├── userrepo.go
│   ├── types
│   │   ├── auth.go
│   │   ├── resp.go
│   │   ├── watchlist.go
│   ├── usecase
│   │   ├── auth_test.go
│   │   ├── auth.go
│   │   ├── watchlist_test.go
│   │   ├── watchlist.go
│   ├── utils
│   │   ├── jwt.go
│   │   ├── password.go
│   │   ├── util.go
├── .env.example
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── README.md
```

## API Endpoints

The service provides the following API endpoints:

### Register

- **URL**: `localhost/api/users/register`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "username": "user",
    "email": "user@email.com",
    "password": "password"
  }
  ```

### Login

- **URL**: `localhost/api/users/login`
- **Method**: `POST`
- **Request Body**:
  
  For login, you can use either the username or email to login. The password is required.
  ```json
  {
    "identifier": "user",
    "password": "password"
  }
  ```

### Logout

- **URL**: `localhost/api/users/logout`
- **Method**: `POST`

### Add Movie

- **URL**: `localhost/api/movies/add`
- **Method**: `POST`
- **Request Body**:
  
  The movie ID (from the TMDB API) and title are required. The image and rating are optional, you can exclude them if you don't have the information.
  ```json
  {
    "movieID": 1,
    "title": "Movie Title",
    "image": "https://image.com",
    "rating": 5,
  }
  ```

### Remove Movie

- **URL**: `localhost/api/movies/remove/{movieID}`
  
  Note that the movie ID used here is the ID from the TMDB API.
  
- **Method**: `DELETE`

### Get Watchlist

- **URL**: `localhost/api/movies/watchlist`
- **Method**: `GET`

### Rate Movie

- **URL**: `localhost/api/movies/rate/{movieID}`
  
  Note that the movie ID used here is the ID from the TMDB API.

- **Method**: `PATCH`
- **Request Body**:
  ```json
  {
    "rating": 5
  }
  ```

### Check is Movie in Watchlist

- **URL**: `localhost/api/movies/watchlist/exist/{movieID}`
  
  Note that the movie ID used here is the ID from the TMDB API.

- **Method**: `GET`

### Get Current Rating of Movie

- **URL**: `localhost/api/movies/rate/{movieID}`
  
  Note that the movie ID used here is the ID from the TMDB API.

- **Method**: `GET`
