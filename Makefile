build:
	@mkdir -p bin
	@go build -o bin ./cmd/***
	@cp -r resources bin

clean:
	@rm -rf bin/*

install:
	@go install ./cmd/***