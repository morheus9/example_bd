-- +migrate Up
CREATE TABLE songs
(
    id           SERIAL PRIMARY KEY,
    group_name   VARCHAR(255) NOT NULL,
    title        VARCHAR(255) NOT NULL,
    release_date DATE,
    link         VARCHAR(2083),
    created_at   TIMESTAMP DEFAULT NOW(),
    updated_at   TIMESTAMP DEFAULT NOW(),
    CONSTRAINT unique_group_song UNIQUE (group_name, title)
);

