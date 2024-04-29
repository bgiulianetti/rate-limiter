# Define default target
all: run

# Set up environment variables for local development
set-env:
	export DAO_TYPE=memory

# Run the application
run:
	go run main.go

# Clean up temporary files (if needed)
clean:
	rm -f *.o
