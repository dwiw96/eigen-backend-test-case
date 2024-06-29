CREATE TABLE IF NOT EXISTS members(
    id INT GENERATED ALWAYS AS IDENTITY
        CONSTRAINT pk_members_id PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
        CONSTRAINT uq_members_code UNIQUE(code),
    name VARCHAR(255) NOT NULL
        CONSTRAINT ck_members_name_len CHECK(LENGTH(name) > 0)
);

CREATE INDEX ix_members_id ON members(id);
CREATE INDEX ix_members_code ON members(code);
CREATE INDEX ix_members_name ON members(name);