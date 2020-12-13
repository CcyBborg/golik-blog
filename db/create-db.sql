DROP DATABASE IF EXISTS golikblog;

CREATE DATABASE golikblog WITH ENCODING = 'UTF8';

\connect golikblog

SET client_encoding = 'UTF8';

CREATE TABLE "user" (
    id INT GENERATED ALWAYS AS IDENTITY,
    username CHAR(30) NOT NULL,
    email VARCHAR(254) NOT NULL,
    password_hash VARCHAR(32) NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    info TEXT,
    PRIMARY KEY(id)
);

CREATE TABLE post (
    id INT GENERATED ALWAYS AS IDENTITY,
    author_id INT NOT NULL REFERENCES "user"(id),
    title VARCHAR(80) NOT NULL,
    summary VARCHAR(300),
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    published_at timestamp,
    content TEXT,
    PRIMARY KEY(id)
);

CREATE TABLE post_comment (
    id INT GENERATED ALWAYS AS IDENTITY,
    author_id INT NOT NULL REFERENCES "user"(id),
    post_id INT NOT NULL REFERENCES post(id),
    created_at timestamp NOT NULL DEFAULT NOW(),
    content TEXT NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE category (
    id INT GENERATED ALWAYS AS IDENTITY,
    title VARCHAR(80) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE post_category (
    post_id INT NOT NULL REFERENCES post(id),
    category_id INT NOT NULL REFERENCES category(id)
);
