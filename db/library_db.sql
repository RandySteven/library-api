CREATE DATABASE library_db; 

SELECT * FROM books;

INSERT INTO authors(name)
VALUES
    ('Randy Steven'),
    ('Matthew Alfredo');

INSERT INTO books (title, description, quantity, cover, author_id)
VALUES
    ('Cara Bikin Orang', 'Hello world Lalala', 10, 'Kertas', 1),
    ('Cara Bikin Orang 2', 'Orang adalah orang', 20, 'Sampul', 2);


SELECT * FROM authors;
SELECT * FROM books;

DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS authors;