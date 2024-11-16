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
	@pnpm dlx tailwindcss -i ./styles/style.css -o ./.dist/css/tailwind.css

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

# Format the Go code
.PHONY: format
format:
	@echo "Formatting Go code..."
	@go fmt ./...
	@templ fmt .
