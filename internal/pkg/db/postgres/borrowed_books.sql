CREATE TABLE IF NOT EXISTS borrowed_books(
    id INT GENERATED ALWAYS AS IDENTITY
        CONSTRAINT pk_borrowed_books_id PRIMARY KEY,
    book_id INT NOT NULL,
        CONSTRAINT fk_borrowed_books_book_id FOREIGN KEY(book_id) REFERENCES books(id),
    member_id INT NOT NULL,
        CONSTRAINT fk_borrowed_books_member_id FOREIGN KEY(member_id) REFERENCES members(id),
    borrowed_at TIMESTAMP NOT NULL DEFAULT NOW(),
    returned_at TIMESTAMP,
    is_returned BOOLEAN
);

CREATE INDEX ix_borrowed_books_book_id ON borrowed_books(book_id);
CREATE INDEX ix_borrowed_books_member_id ON borrowed_books(member_id);
CREATE INDEX ix_borrowed_books_borrowed_at ON borrowed_books(borrowed_at);
CREATE INDEX ix_borrowed_books_returned_at ON borrowed_books(returned_at);