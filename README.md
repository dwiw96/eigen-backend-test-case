# Overview
This project is technical test as backend developer for PT EIGEN TRI MATHEMA

## Supported Features
1. Members can borrow books with conditions
   - Members may not borrow more than 2 books
   - Borrowed books are not borrowed by other members
   - Member is currently not being penalized

2. Member returns the book with conditions
   - The returned book is a book that the member has borrowed
   - If the book is returned after more than 7 days, the member will be subject to a penalty. Member with penalty cannot able to borrow the book for 3 days

3. Check the book
   - Shows all existing books and quantities
   - Books that are being borrowed are not counted

4. Member check
   - Shows all existing members
   - The number of books being borrowed by each member

## Additional Features
1. Insert Books
   - User can insert more than 1 book at once.

2. Insert Members
   - User can insert more than 1 member at once.

## Information:
- API Documentation using Swagger in yaml file: /Swagger_APIDocumentation.yaml
- Postman collection: /eigen-backend-test-case.postman_collection.json
- Database: /backup/backup_20240629_Final.sql

## Algorith Test
github link for algorithm test: https://github.com/dwiw96/eigen-backend-test-algorithm

## How To Run This Project
* Run postgres with the same environment as in the utils/driver/postgres/docker-compose.yml
* Create table using file on /backup/backup_20240629_Final.sql
* Run by the executable file that compatible with your system
  note: I create executable for linux, macOS, and windows with an AMD64 architecture

## My System Specifications
This is the system I use for creating this project
* OS: Ubuntu 22.04.4 LTS (Jammy)
* Database: PostgreSQL 15.2 (Debian 15.2-1.pgdg110+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 10.2.1-6) 10.2.1 20210110, 64-bit
* Docker: Docker version 26.0.1, build d260a54
* golang: go version go1.20.3 linux/amd64
