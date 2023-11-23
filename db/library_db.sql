CREATE DATABASE library_db; 

SELECT * FROM books;

INSERT INTO books (title, description, quantity, cover)
VALUES
    ('Cara Bikin Orang', 'Hello world Lalala', 10, 'Kertas'),
    ('Cara Bikin Orang', 'Orang adalah orang', 20, 'Sampul');

INSERT INTO authors(name)
VALUES
    ('Randy Steven'),
    ('Matthew Alfredo');

SELECT * FROM authors;
SELECT * FROM books;

DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS authors;