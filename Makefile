# Default target is to run the server
.PHONY: default
default: dev

# Common build steps
.PHONY: templ tailwind
templ:
	@echo "Running templ generate..."
	@templ generate

tailwind:
	@echo "Building TailwindCSS..."
	@pnpm tailwindcss -i ./styles/style.css -o ./.dist/css/tailwind.css

# Build the Go binary for serving the app
.PHONY: dev
dev:
	@echo "Running server..."
	@air

# Build the files into ./dist directory
.PHONY: build
build: templ tailwind
	@echo "Building static files..."
	@go run . --mode=build

# Build the files into ./dist directory and serve them
.PHONY: preview
preview: templ tailwind
	@echo "Building and serving static files..."
	@go run . --mode=preview

# Clean generated files
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf ./dist
	@rm -f myapp

# Run tests with coverage
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./... -cover

# Format the Go code
.PHONY: format
format:
	@echo "Formatting Go code..."
	@go fmt ./...

# Run the Go vet tool to check for potential issues
.PHONY: vet
vet:
	@echo "Running Go vet..."
	@go vet ./...

	@echo "Running staticcheck..."
	@staticcheck ./...
