# Authentication

## Overview

Chirpy uses JWT (JSON Web Tokens) for authentication. The API supports both access tokens and refresh tokens.

## Authentication Flow

1. User logs in with email/password
2. Server returns an access token and refresh token
3. Access token is used for API requests
4. Refresh token can be used to obtain new access tokens

## Token Types

### Access Token

- Format: Bearer token
- Header: `Authorization: Bearer <token>`
- Expiration: Short-lived

### Refresh Token

- Used to obtain new access tokens
- Longer expiration time
- Stored securely in the database

## Endpoints

### Login

```http
POST /api/login
Content-Type: application/json

{
    "email": "user@example.com",
    "password": "password"
}
```

### Refresh Token

```http
POST /api/refresh
Authorization: Bearer <refresh_token>
```

## Security Considerations

- Access tokens are short-lived for security
- Refresh tokens can be revoked if compromised
- Passwords are hashed using bcrypt
- HTTPS is required in production
