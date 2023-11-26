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

INSERT INTO users(name, email, phone_number)
VALUES
    ('Senju Man', 'senju.man@shopee.com', '+62123456789'),
    ('Juju Man', 'juju.man@shopee.com', '+62987654321'),
    ('Manju Man', 'manju.man@shopee.com', '+62012301232');

SELECT * FROM authors;
SELECT * FROM users;
SELECT * FROM books;
SELECT * FROM borrows;

SELECT * FROM "borrows" WHERE id = 30 

DROP TABLE IF EXISTS borrows;
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS authors;
DROP TABLE IF EXISTS users;