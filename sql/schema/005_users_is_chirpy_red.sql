-- +goose Up 
ALTER TABLE users add column is_chirpy_red boolean not null default false;


-- +goose Down
Alter table users drop column is_chirpy_red;