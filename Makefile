gorun:
	@go run app/*.go

docker-compose:
	@docker-compose -f files/docker/docker-compose.yml  up

docker-compose-build:
	@docker-compose -f files/docker/docker-compose.yml  up --build
