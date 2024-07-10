.PHONY: dcu dcd tidy test

dcu:
	docker-compose up -d

dcd:
	docker-compose down

tidy:
	go mod tidy

test:
	go test ./...