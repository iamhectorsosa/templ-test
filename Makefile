default: dev

format:
	@go fmt ./...
	@templ fmt .

build/templ:
	@templ generate

dev/templ:
	@templ generate --watch

build/tailwind:
	@pnpm dlx tailwindcss -i ./styles/style.css -o ./.dist/css/tailwind.css --minify

dev/tailwind:
	@pnpm dlx tailwindcss -i ./styles/style.css -o ./.dist/css/tailwind.css --watch

dev/server:
	go run github.com/air-verse/air@latest \
	--build.cmd "go build -o tmp/main" \
	--build.bin "./tmp/main" \
	--build.delay "0" \
	--build.include_ext "go,templ,css" \
	--build.exclude_dir ".dist" \
 	--build.stop_on_error "false" \
 	--proxy.app_port 3000 \
 	--proxy.enabled true \
 	--proxy.proxy_port 3001 \
	--misc.clean_on_exit true

dev:
	@make -j3 dev/templ dev/tailwind dev/server

preview:
	@make -j2 build/templ build/tailwind
	@go run .

build:
	@make -j2 build/templ build/tailwind
	@go run . --mode=build
