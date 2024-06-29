CREATE TABLE IF NOT EXISTS books(
    id INT GENERATED ALWAYS AS IDENTITY
        CONSTRAINT pk_books_id PRIMARY KEY,
    code VARCHAR(10) NOT NULL,
        CONSTRAINT uq_books_code UNIQUE(code),
    title VARCHAR(255) NOT NULL,
        CONSTRAINT ck_books_title_len CHECK(LENGTH(title) > 0),
    author VARCHAR(255) NOT NULL,
        CONSTRAINT ck_books_author_len CHECK(LENGTH(author) > 0),
    stock INT,
        CONSTRAINT ck_books_stock CHECK(stock >= 0),
    total_amount INT
);

CREATE INDEX ix_books_id ON books(id);
CREATE INDEX ix_books_code ON books(code);
CREATE INDEX ix_books_title ON books(title);
CREATE INDEX ix_books_author ON books(author);