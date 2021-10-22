gorun:
	@go run app/*.go

docker-build:
	@docker build -f files/docker/Dockerfile -t live-chat .

docker-compose:
	@docker-compose -f files/docker/docker-compose.yml  up

docker-compose-build:
	@docker-compose -f files/docker/docker-compose.yml  up --build
