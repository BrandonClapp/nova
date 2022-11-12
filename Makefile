BINARY_NAME=novaApp

build:
	@go mod vendor
	@echo "Building..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Built!"

run: build
	@./tmp/${BINARY_NAME} &
	@echo "Started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

start_compose:
	docker-compose up -d

stop_compose:
	docker-compose down

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped!"

restart: stop start

## coverage: displays test coverage
coverage:
	@go test -cover ./...

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

