DROP DATABASE IF EXISTS album_store;
CREATE DATABASE album_store;
\c album_store;

CREATE TABLE albums (
    Album_Id SERIAL PRIMARY KEY,
    Title TEXT NOT NULL,
    Artist TEXT NOT NULL,
    Price NUMERIC NOT NULL
);

INSERT INTO albums (Title, Artist, Price)
VALUES
    ('Atlas Shrugged', 'Ayn Rand', 50.50),
    ('The Cellist of Sarajevo', 'Steven Galloway', 32.99),
    ('Crime and Punishment', 'Fyodor Dostoevsky', 35.00),
    ('The Pet Sematary', 'Stephen King', 29.99);

CREATE USER myuser WITH PASSWORD 'pass';
GRANT SELECT, INSERT, UPDATE, DELETE ON albums TO myuser;
