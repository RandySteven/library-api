CREATE DATABASE library_db; 

SELECT * FROM books;

INSERT INTO books (title, description, quantity, cover)
VALUES
    ('Cara Bikin Orang', 'Hello world Lalala', 10, 'Kertas'),
    ('Cara Bikin Orang', 'Orang adalah orang', 20, 'Sampul');

DROP TABLE IF EXISTS books;