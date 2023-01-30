build:
	go build -o ./bin/kv-cli

run: build
	./bin/kv-cli