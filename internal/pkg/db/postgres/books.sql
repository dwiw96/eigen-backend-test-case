CREATE TABLE IF NOT EXISTS books(
    id INT GENERATED ALWAYS AS IDENTITY
        CONSTRAINT pk_books_id PRIMARY KEY,
    code VARCHAR(10) NOT NULL,
        CONSTRAINT uq_books_code UNIQUE(code),
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    stock INT
);