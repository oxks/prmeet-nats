include .env
#export

migrate: 
	migrate -source file://postgres/migrations  \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/permit?sslmode=disable up

rollback: 
	migrate -source file://postgres/migrations  \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/permit?sslmode=disable down
drop: 
	migrate -source file://postgres/migrations  \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/permit?sslmode=disable drop

migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir postgres/migrations $$name

sqlc: 
	sqlc generate



	