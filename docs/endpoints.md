# API Endpoints

## Users

### Create User

```http
POST /api/users
Content-Type: application/json

{
    "email": "user@example.com",
    "password": "password"
}
```

### Update User

```http
PUT /api/users
Authorization: Bearer <token>
Content-Type: application/json

{
    "email": "newemail@example.com",
    "password": "newpassword"
}
```

### Upgrade to Chirpy Red

```http
PUT /api/users/upgrade
Authorization: Bearer <token>
```

## Chirps

### Create Chirp

```http
POST /api/chirps
Authorization: Bearer <token>
Content-Type: application/json

{
    "message": "Hello, Chirpy!"
}
```

### Get All Chirps

```http
GET /api/chirps
```

### Get Chirp by ID

```http
GET /api/chirps/{id}
```

### Delete Chirp

```http
DELETE /api/chirps/{id}
Authorization: Bearer <token>
```

### Get User's Chirps

```http
GET /api/chirps/users/{user_id}
```

## Authentication

See [Authentication Documentation](./authentication.md) for detailed information about authentication endpoints.
