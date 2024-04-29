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
	moq -out ./controllers/mock_notifications_service_test.go -pkg controllers ./controllers NotificationsService
	moq -out ./controllers/mock_rules_service_test.go -pkg controllers ./controllers RulesService
	moq -out ./services/mock_notifications_container_test.go -pkg services ./services NotificationsContainer
	moq -out ./services/mock_rules_container_test.go -pkg services ./services RulesContainer


install-deps:
	go get -u github.com/gin-gonic/gin
	go install github.com/matryer/moq@latest
	go get github.com/stretchr/testify
	go get github.com/redis/go-redis/v9

initialize: install-deps mock test run

.PHONY: all run clean mock install-deps initialize
