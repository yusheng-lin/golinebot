all: mongo run
mongo:
	docker-compose up -d
build:
	wire .;swag init
run:
	go run golinebot
down:
	docker-compose down