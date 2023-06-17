include .env
SERVER=api

init:
	cd ${SERVER} && go get .

build:
	cd ${SERVER} && docker build -t server .

update:
	-docker-compose down
	make build
	docker-compose up -d
