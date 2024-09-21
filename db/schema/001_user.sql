-- +goose Up
create table user (
    user_id UUID primary key ,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    phoneNumber TEXT UNIQUE NOT NULL  ,
    email varchar(50) UNIQUE NOT NULL,
    address TEXT NOT NULL,
    password TEXT NOT NULL,
    
 );
-- +goose Down 
DROP TABLE user;