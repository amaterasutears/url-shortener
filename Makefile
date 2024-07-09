.PHONY: dcu dcd tidy

dcu:
	docker-compose up -d

dcd:
	docker-compose down

tidy:
	go mod tidy
