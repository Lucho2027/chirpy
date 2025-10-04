# Database Schema

## Users Table

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    hashed_password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    is_chirpy_red BOOLEAN DEFAULT FALSE
);
```

## Chirps Table

```sql
CREATE TABLE chirps (
    id UUID PRIMARY KEY,
    message TEXT NOT NULL,
    user_id UUID REFERENCES users(id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
```

## Refresh Tokens Table

```sql
CREATE TABLE refresh_tokens (
    token TEXT PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    created_at TIMESTAMP NOT NULL,
    revoked BOOLEAN DEFAULT FALSE
);
```

## Relationships

- Each chirp belongs to a user (user_id foreign key)
- Each refresh token belongs to a user (user_id foreign key)
- Users can have multiple chirps and refresh tokens

## Indexes

- Users email (UNIQUE)
- Chirps user_id (for faster lookup of user's chirps)
- Refresh tokens user_id (for faster token lookup)
