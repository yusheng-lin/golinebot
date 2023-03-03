mongo:
	docker-compose up -d
build:
	wire .;swag init
run:
	go run golinebot