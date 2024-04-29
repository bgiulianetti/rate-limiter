# Define default target
all: test run

# DAO_TYPE options:
# memory
# redis
export DAO_TYPE := memory

# Run all tests
test:
	go test -v ./...

# Run the application
run:
	go run main.go

# Clean up temporary files (if needed)
clean:
	rm -f *.o

mock:
	moq -out ./controllers/mock_service_test.go -pkg controllers ./services Service
	moq -out ./services/mock_dao_test.go -pkg services ./dao Container

install-deps:
	go get -u github.com/gin-gonic/gin
	go install github.com/matryer/moq@latest
	go get github.com/stretchr/testify
	go get github.com/redis/go-redis/v9

initialize: install-deps mock test run

.PHONY: all run clean mock install-deps initialize
