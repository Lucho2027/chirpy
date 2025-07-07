-- +goose Up 
ALTER TABLE users add column hashed_password text not null default 'unset';


-- +goose Down
Alter table users drop column hashed_password;