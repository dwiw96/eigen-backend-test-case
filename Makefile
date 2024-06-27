dockerPgStart:
	sudo docker compose -f ./utils/driver/postgres/docker-compose.yml up -d
dockerPgExec:
	sudo docker exec -it postgres_eigen_book_project psql -h localhost -p 5432 -U dwiwahyudi eigen_book_project
dockerPgStop:
	sudo docker container stop postgres_eigen_book_project