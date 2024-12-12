-- +migrate Up
CREATE TABLE song_verses
(
    id           SERIAL PRIMARY KEY,
    song_id      INT  NOT NULL REFERENCES songs (id) ON DELETE CASCADE,
    verse_number INT  NOT NULL,
    text         TEXT NOT NULL,
    created_at   TIMESTAMP DEFAULT NOW(),
    updated_at   TIMESTAMP DEFAULT NOW(),
    CONSTRAINT unique_song_verse UNIQUE (song_id, verse_number)
);
