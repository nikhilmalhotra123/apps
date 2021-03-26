startdb:
	brew services start mongodb-community@4.4

stopdb:
	brew services stop mongodb-community@4.4

run:
	go run main.go

dev-up:
	docker-compose up -d

dev-down:
	docker-compose down
