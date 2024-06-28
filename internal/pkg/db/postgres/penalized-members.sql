CREATE TABLE IF NOT EXISTS penalized_members(
    id INT GENERATED ALWAYS AS IDENTITY
        CONSTRAINT pk_penalized_members_id PRIMARY KEY,
    member_id INT NOT NULL,
        CONSTRAINT fk_penalized_members_member_id FOREIGN KEY(member_id) REFERENCES members(id),
    penalty_start TIMESTAMP NOT NULL,
    penalty_end TIMESTAMP
);

CREATE INDEX ix_penalized_member_member_id ON penalized_member(member_id);
CREATE INDEX ix_penalized_member_penalty_start ON penalized_member(penalty_start);
CREATE INDEX ix_penalized_member_penalty_end ON penalized_member(penalty_end);