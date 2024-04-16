# Makefile to handle pre-setup for Docker Compose by copying necessary folders

# Define the paths relative to the Makefile location
GREETING_HTTP_SERVER_PATH = ../greeting_http_server
GREETING_LAMBDA_PATH = ../greeting_lambda
DESTINATION = .

# Targets that do not produce files
.PHONY: all clean copy_folders dc_up

# Default target
all: clean copy_folders dc_up

# Clean up the existing directories
clean:
	@echo "Cleaning up existing directories..."
	rm -rf $(DESTINATION)/greeting_http_server
	rm -rf $(DESTINATION)/greeting_lambda

# Copy folders to the Docker Compose directory
copy_folders:
	@echo "Copying folders to the Docker Compose directory..."
	cp -R $(GREETING_HTTP_SERVER_PATH) $(DESTINATION)/greeting_http_server
	cp -R $(GREETING_LAMBDA_PATH) $(DESTINATION)/greeting_lambda
	@echo "Folders copied successfully."

# Start Docker Compose services
dc_up:
	@echo "Building Docker containers ..."
	docker compose build
	@echo "Starting Docker Compose services..."
	docker compose up -d
	@echo "Docker Compose services are up and running."

# Shut down Docker Compose
dc_down:
	@echo "Stopping Docker Compose services ..."
	docker compose down
	@echo "Docker Compose services have been shut down."
