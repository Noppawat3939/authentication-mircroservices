# Authenication microservice

### Key Features

- generate access token and refresh token (payload jwt configuration)
- verify access token and refresh token

### API Documentation

This document describes the API endpoints available for authentication and authorization.

base URL `localhost:8000`

### Authentication Endpoints

#### 1. Generate JWT Token

- **Endpoint:** `/api/v1/jwt/generate`
- **Method:** `POST`
- **Description:** This endpoint generates access and refresh tokens. It requires a request body conforming to a specific interface (detailed below). You can optionally specify an expiration time for the generated JWT tokens.

- **Request Body:**

```json
{
  "user": "john", // Replace with actual interface fields
  "age": 30, // Replace with actual interface fields
  "...": "...", // Add other fields
  "expired_hour": 24 // Optional: Expiration time in hours. Defaults to a standard expiration if omitted.
}
```

- **Response (Success - 200 OK):**

```json
{
  "success": true,
  "code": 200,
  "data": {
    "access_token": "eyJ...", // JWT access token
    "refresh_token": "eyJ..." // JWT refresh token
  }
}
```

#### 2. Verify JWT Token

- **Endpoint:** `/api/v1/jwt/verify`
- **Method:** `POST`
- **Description:** This endpoint verifies and decodes a JWT (either access or refresh token). It requires a request body containing either the access_token or the refresh_token.

- **Request Body:**

```json
{
  "access_token": "eyJ...", // JWT access token (optional)
  "refresh_token": "eyJ..." // JWT refresh token (optional)
}
```

- **Response (Success - 200 OK):**

```json
{
  "success": true,
  "code": 200,
  "data": {
    "user": "john",
    "age": 30,
    "exp": 1678886400 // Example expiration timestamp
  }
}
```

#### 3. User Registration

- **Endpoint:** `/api/v1/user/register`
- **Method:** `POST`
- **Description:** This endpoint insert user data in to database (MongoDB) but required 3 fields.

- **Request Body:**

```json
{
  "email": "example@gmail.com",
  "username": "example_username",
  "password": "example_paSsw0rD1"
}
```

- **Response (Success - 200 OK):**

```json
{
  "success": true,
  "code": 200,
  "data": {
    "_id": "xxxx...", // id generated from MongoDB
    "email": "example@gmail.com",
    "username": "example_username",
    "password": "xxxx..." // password hashed from api
  }
}
```
