.PHONY: css-watch
css-watch:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch
	
.PHONY: dev
dev:
	make css-watch && air && templ generate --watch --proxy=http://localhost:8080

.PHONY: build
build:
	make tailwind-build
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/main.go