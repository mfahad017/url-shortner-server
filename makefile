.PHONY: dev prod

dev:
	@echo "Running in Dev mode. Watching for changes and reloading..."
	ENV=dev air

prod:
	@echo "Running in production mode"
	go run main.go -env=prod