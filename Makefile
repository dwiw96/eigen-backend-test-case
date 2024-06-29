dockerPgStart:
	sudo docker compose -f ./utils/driver/postgres/docker-compose.yml up -d
dockerPgExec:
	sudo docker exec -it postgres_eigen_book_project psql -h localhost -p 5432 -U dwiwahyudi eigen_book_project
dockerPgStop:
	sudo docker container stop postgres_eigen_book_project

pgCreateAllTable:
	sudo docker exec -i postgres_eigen_book_project psql -h localhost -p 5432 -U dwiwahyudi -W eigen_book_project < internal/pkg/db/postgres/create_all_table.sql
pgDropAllTable:
	sudo docker exec -i postgres_eigen_book_project psql -h localhost -p 5432 -U dwiwahyudi -W eigen_book_project < internal/pkg/db/postgres/drop_all_table.sql

pgBackup:
	sudo docker exec -t postgres_eigen_book_project pg_dump -U dwiwahyudi -d eigen_book_project > /home/dwiw22/test-job/eigen/eigen-backend-test-case/backup/backup_$(date).sql