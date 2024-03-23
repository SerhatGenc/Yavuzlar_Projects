
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    roles VARCHAR(100)
);



CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    bookname VARCHAR(255) NOT NULL,
    writer VARCHAR(255) NOT NULL,
    user_id INT REFERENCES users(id),
    deadline DATE
);

