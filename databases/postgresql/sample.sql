CREATE TABLE mia (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255), --VARCHAR is a text column that has constraint on size, typically how long the text can be i.e 255 characters
    email VARCHAR(255),
    password TEXT,
    age INT
);