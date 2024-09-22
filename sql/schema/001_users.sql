-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL ,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    phoneNumber TEXT UNIQUE NOT NULL  ,
    email VARCHAR(50) UNIQUE NOT NULL,
    address TEXT NOT NULL,
    password TEXT NOT NULL
    
 );
-- +goose Down
DROP TABLE user;
