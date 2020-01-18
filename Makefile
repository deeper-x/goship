install:
	@echo "Installing goship..."
	@go install src/goship.go
run:	
	@go run src/goship.go > /dev/null &
	@echo "Running goship in background...."
kill:
	@echo "Killing goship..."
	@pkill goship
