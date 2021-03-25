startdb:
	brew services start mongodb-community@4.4

stopdb:
	brew services stop mongodb-community@4.4

run:
	go run main.go

docker build:
	docker build -t jobapps:latest .
