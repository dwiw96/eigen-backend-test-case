--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.2 (Debian 15.2-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: books; Type: TABLE; Schema: public; Owner: dwiwahyudi
--

CREATE TABLE public.books (
    id integer NOT NULL,
    code character varying(10) NOT NULL,
    title character varying(255) NOT NULL,
    author character varying(255) NOT NULL,
    stock integer,
    total_amount integer,
    CONSTRAINT ck_books_author_len CHECK ((length((author)::text) > 0)),
    CONSTRAINT ck_books_stock CHECK ((stock >= 0)),
    CONSTRAINT ck_books_title_len CHECK ((length((title)::text) > 0))
);


ALTER TABLE public.books OWNER TO dwiwahyudi;

--
-- Name: books_id_seq; Type: SEQUENCE; Schema: public; Owner: dwiwahyudi
--

ALTER TABLE public.books ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.books_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: borrowed_books; Type: TABLE; Schema: public; Owner: dwiwahyudi
--

CREATE TABLE public.borrowed_books (
    id integer NOT NULL,
    book_id integer NOT NULL,
    member_id integer NOT NULL,
    borrowed_at timestamp without time zone DEFAULT now() NOT NULL,
    returned_at timestamp without time zone,
    is_returned boolean
);


ALTER TABLE public.borrowed_books OWNER TO dwiwahyudi;

--
-- Name: borrowed_books_id_seq; Type: SEQUENCE; Schema: public; Owner: dwiwahyudi
--

ALTER TABLE public.borrowed_books ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.borrowed_books_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: members; Type: TABLE; Schema: public; Owner: dwiwahyudi
--

CREATE TABLE public.members (
    id integer NOT NULL,
    code character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    CONSTRAINT ck_members_name_len CHECK ((length((name)::text) > 0))
);


ALTER TABLE public.members OWNER TO dwiwahyudi;

--
-- Name: members_id_seq; Type: SEQUENCE; Schema: public; Owner: dwiwahyudi
--

ALTER TABLE public.members ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.members_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: penalized_members; Type: TABLE; Schema: public; Owner: dwiwahyudi
--

CREATE TABLE public.penalized_members (
    id integer NOT NULL,
    member_id integer NOT NULL,
    penalty_start timestamp without time zone DEFAULT now() NOT NULL,
    penalty_end timestamp without time zone
);


ALTER TABLE public.penalized_members OWNER TO dwiwahyudi;

--
-- Name: penalized_members_id_seq; Type: SEQUENCE; Schema: public; Owner: dwiwahyudi
--

ALTER TABLE public.penalized_members ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.penalized_members_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: dwiwahyudi
--

COPY public.books (id, code, title, author, stock, total_amount) FROM stdin;
1	JK-45	Harry Potter	J.K Rowling	1	1
2	SHR-1	A Study in Scarlet	Arthur Conan Doyle	1	1
3	TW-11	Twilight	Stephenie Meyer	1	1
4	HOB-83	The Hobbit, or There and Back Again	J.R.R. Tolkien	1	1
5	NRN-7	The Lion, the Witch and the Wardrobe	C.S. Lewis	1	1
6	ACD-01	Sherlock Holmes Chapter 1	Sir Arthur Conan Doyle	3	3
8	ACD-03	Sherlock Holmes Chapter 3	Sir Arthur Conan Doyle	12	14
7	ACD-02	Sherlock Holmes Chapter 2	Sir Arthur Conan Doyle	7	7
\.


--
-- Data for Name: borrowed_books; Type: TABLE DATA; Schema: public; Owner: dwiwahyudi
--

COPY public.borrowed_books (id, book_id, member_id, borrowed_at, returned_at, is_returned) FROM stdin;
1	8	4	2024-06-29 15:29:13	\N	f
2	8	4	2024-06-29 15:29:14	\N	f
3	7	5	2024-06-29 15:29:54	2024-06-29 15:30:17	t
4	7	6	2024-06-29 15:29:58	2024-07-07 15:32:01	t
\.


--
-- Data for Name: members; Type: TABLE DATA; Schema: public; Owner: dwiwahyudi
--

COPY public.members (id, code, name) FROM stdin;
1	M001	Angga
2	M002	Ferry
3	M003	Putri
4	M004	Dwi
5	M005	Wahyu
6	M006	Yudi
\.


--
-- Data for Name: penalized_members; Type: TABLE DATA; Schema: public; Owner: dwiwahyudi
--

COPY public.penalized_members (id, member_id, penalty_start, penalty_end) FROM stdin;
1	6	2024-07-07 15:32:01	2024-07-10 15:32:01
\.


--
-- Name: books_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dwiwahyudi
--

SELECT pg_catalog.setval('public.books_id_seq', 8, true);


--
-- Name: borrowed_books_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dwiwahyudi
--

SELECT pg_catalog.setval('public.borrowed_books_id_seq', 4, true);


--
-- Name: members_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dwiwahyudi
--

SELECT pg_catalog.setval('public.members_id_seq', 6, true);


--
-- Name: penalized_members_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dwiwahyudi
--

SELECT pg_catalog.setval('public.penalized_members_id_seq', 1, true);


--
-- Name: books pk_books_id; Type: CONSTRAINT; Schema: public; Owner: dwiwahyudi
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT pk_books_id PRIMARY KEY (id);


--
-- Name: borrowed_books pk_borrowed_books_id; Type: CONSTRAINT; Schema: public; Owner: dwiwahyudi
--

ALTER TABLE ONLY public.borrowed_books
    ADD CONSTRAINT pk_borrowed_books_id PRIMARY KEY (id);


--
-- Name: members pk_members_id; Type: CONSTRAINT; Schema: public; Owner: dwiwahyudi
--

ALTER TABLE ONLY public.members
    ADD CONSTRAINT pk_members_id PRIMARY KEY (id);


--
-- Name: penalized_members pk_penalized_members_id; Type: CONSTRAINT; Schema: public; Owner: dwiwahyudi
--

ALTER TABLE ONLY public.penalized_members
    ADD CONSTRAINT pk_penalized_members_id PRIMARY KEY (id);


--
-- Name: books uq_books_code; Type: CONSTRAINT; Schema: public; Owner: dwiwahyudi
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT uq_books_code UNIQUE (code);


--
-- Name: members uq_members_code; Type: CONSTRAINT; Schema: public; Owner: dwiwahyudi
--

ALTER TABLE ONLY public.members
    ADD CONSTRAINT uq_members_code UNIQUE (code);


--
-- Name: ix_books_author; Type: INDEX; Schema: public; Owner: dwiwahyudi
--

CREATE INDEX ix_books_author ON public.books USING btree (author);


--
-- Name: ix_books_code; Type: INDEX; Schema: public; Owner: dwiwahyudi
--

CREATE INDEX ix_books_code ON public.books USING btree (code);


--
-- Name: ix_books_id; Type: INDEX; Schema: public; Owner: dwiwahyudi
--

CREATE INDEX ix_books_id ON public.books USING btree (id);


--
-- Name: ix_books_title; Type: INDEX; Schema: public; Owner: dwiwahyudi
--

CREATE INDEX ix_books_title ON public.books USING btree (title);


--
-- Name: ix_borrowed_books_book_id; Type: INDEX; Schema: public; Owner: dwiwahyudi
--

CREATE INDEX ix_borrowed_books_book_id ON public.borrowed_books USING btree (book_id);


--
-- Name: ix_borrowed_books_borrowed_at; Type: INDEX; Schema: public; Owner: dwiwahyudi
--

CREATE INDEX ix_borrowed_books_borrowed_at ON public.borrowed_books USING btree (borrowed_at);


--
-- Name: ix_borrowed_books_member_id; Type: INDEX; Schema: public; Owner: dwiwahyudi
--

CREATE INDEX ix_borrowed_books_member_id ON public.borrowed_books USING btree (member_id);


--
-- Name: ix_borrowed_books_returned_at; Type: INDEX; Schema: public; Owner: dwiwahyudi
--

CREATE INDEX ix_borrowed_books_returned_at ON public.borrowed_books USING btree (returned_at);


--
-- Name: ix_members_code; Type: INDEX; Schema: public; Owner: dwiwahyudi
--

CREATE INDEX ix_members_code ON public.members USING btree (code);


--
-- Name: ix_members_id; Type: INDEX; Schema: public; Owner: dwiwahyudi
--

CREATE INDEX ix_members_id ON public.members USING btree (id);


--
-- Name: ix_members_name; Type: INDEX; Schema: public; Owner: dwiwahyudi
--

CREATE INDEX ix_members_name ON public.members USING btree (name);


--
-- Name: ix_penalized_member_member_id; Type: INDEX; Schema: public; Owner: dwiwahyudi
--

CREATE INDEX ix_penalized_member_member_id ON public.penalized_members USING btree (member_id);


--
-- Name: ix_penalized_member_penalty_end; Type: INDEX; Schema: public; Owner: dwiwahyudi
--

CREATE INDEX ix_penalized_member_penalty_end ON public.penalized_members USING btree (penalty_end);


--
-- Name: ix_penalized_member_penalty_start; Type: INDEX; Schema: public; Owner: dwiwahyudi
--

CREATE INDEX ix_penalized_member_penalty_start ON public.penalized_members USING btree (penalty_start);


--
-- Name: borrowed_books fk_borrowed_books_book_id; Type: FK CONSTRAINT; Schema: public; Owner: dwiwahyudi
--

ALTER TABLE ONLY public.borrowed_books
    ADD CONSTRAINT fk_borrowed_books_book_id FOREIGN KEY (book_id) REFERENCES public.books(id);


--
-- Name: borrowed_books fk_borrowed_books_member_id; Type: FK CONSTRAINT; Schema: public; Owner: dwiwahyudi
--

ALTER TABLE ONLY public.borrowed_books
    ADD CONSTRAINT fk_borrowed_books_member_id FOREIGN KEY (member_id) REFERENCES public.members(id);


--
-- Name: penalized_members fk_penalized_members_member_id; Type: FK CONSTRAINT; Schema: public; Owner: dwiwahyudi
--

ALTER TABLE ONLY public.penalized_members
    ADD CONSTRAINT fk_penalized_members_member_id FOREIGN KEY (member_id) REFERENCES public.members(id);


--
-- PostgreSQL database dump complete
--

