BEGIN;
CREATE TABLE IF NOT EXISTS members(
    id INT GENERATED ALWAYS AS IDENTITY
        CONSTRAINT pk_members_id PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
        CONSTRAINT uq_members_code UNIQUE(code),
    name VARCHAR(255) NOT NULL
);

CREATE INDEX ix_members_id ON members(id);
CREATE INDEX ix_members_code ON members(code);
CREATE INDEX ix_members_name ON members(name);

CREATE TABLE IF NOT EXISTS books(
    id INT GENERATED ALWAYS AS IDENTITY
        CONSTRAINT pk_books_id PRIMARY KEY,
    code VARCHAR(10) NOT NULL,
        CONSTRAINT uq_books_code UNIQUE(code),
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    stock INT
);

CREATE INDEX ix_books_id ON books(id);
CREATE INDEX ix_books_code ON books(code);
CREATE INDEX ix_books_title ON books(title);
CREATE INDEX ix_books_author ON books(author);

CREATE TABLE IF NOT EXISTS penalized_members(
    id INT GENERATED ALWAYS AS IDENTITY
        CONSTRAINT pk_penalized_members_id PRIMARY KEY,
    member_id INT NOT NULL,
        CONSTRAINT fk_penalized_members_member_id FOREIGN KEY(member_id) REFERENCES members(id),
    penalty_start TIMESTAMP NOT NULL,
    penalty_end TIMESTAMP
);

CREATE INDEX ix_penalized_member_member_id ON penalized_members(member_id);
CREATE INDEX ix_penalized_member_penalty_start ON penalized_members(penalty_start);
CREATE INDEX ix_penalized_member_penalty_end ON penalized_members(penalty_end);

CREATE TABLE IF NOT EXISTS borrowed_books(
    id INT GENERATED ALWAYS AS IDENTITY
        CONSTRAINT pk_borrowed_books_id PRIMARY KEY,
    book_id INT NOT NULL,
        CONSTRAINT fk_borrowed_books_book_id FOREIGN KEY(book_id) REFERENCES books(id),
    member_id INT NOT NULL,
        CONSTRAINT fk_borrowed_books_member_id FOREIGN KEY(member_id) REFERENCES members(id),
    borrowed_at TIMESTAMP NOT NULL,
    returned_at TIMESTAMP,
    is_returned BOOLEAN
);

CREATE INDEX ix_borrowed_books_book_id ON borrowed_books(book_id);
CREATE INDEX ix_borrowed_books_member_id ON borrowed_books(member_id);
CREATE INDEX ix_borrowed_books_borrowed_at ON borrowed_books(borrowed_at);
CREATE INDEX ix_borrowed_books_returned_at ON borrowed_books(returned_at);

COMMIT;