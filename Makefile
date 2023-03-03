mongo:
	docker-compose up -d
build:
	wire .
run:
	go run golinebot