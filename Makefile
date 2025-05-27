.PHONY: run build clean add get update delete list tags

# Server commands
run:
	go run main.go

build:
	go build -o track-it

clean:
	rm -f track-it

# API commands
add:
	@read -p "Enter problem ID: " problem_id; \
	read -p "Enter title: " title; \
	read -p "Enter tags (comma-separated): " tags; \
	curl -X POST http://localhost:8080/problems \
		-H "Content-Type: application/json" \
		-d "{\"problem_id\": $$problem_id, \"title\": \"$$title\", \"tags\": [$$(echo $$tags | sed 's/,/","/g')]}"

get:
	@read -p "Enter problem ID: " id; \
	curl http://localhost:8080/problems/$$id

update:
	@read -p "Enter problem ID to update: " id; \
	read -p "Enter new problem ID: " problem_id; \
	read -p "Enter new title: " title; \
	read -p "Enter new tags (comma-separated): " tags; \
	curl -X PUT http://localhost:8080/problems/$$id \
		-H "Content-Type: application/json" \
		-d "{\"problem_id\": $$problem_id, \"title\": \"$$title\", \"tags\": [$$(echo $$tags | sed 's/,/","/g')]}"

delete:
	@read -p "Enter problem ID to delete: " id; \
	curl -X DELETE http://localhost:8080/problems/$$id

list:
	curl http://localhost:8080/problems

tags:
	@read -p "Enter tags to search (comma-separated): " tags; \
	curl "http://localhost:8080/problems/tags?tags=$$tags"

# Help command
help:
	@echo "Available commands:"
	@echo "  make run     - Run the server"
	@echo "  make build   - Build the application"
	@echo "  make clean   - Clean build artifacts"
	@echo "  make add     - Add a new problem"
	@echo "  make get     - Get a problem by ID"
	@echo "  make update  - Update a problem"
	@echo "  make delete  - Delete a problem"
	@echo "  make list    - List all problems"
	@echo "  make tags    - Search problems by tags" 