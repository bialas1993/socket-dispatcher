BUILDER=go build -ldflags="-s -w"
BIN_NAME=socket-dispatcher
MAIN=cmd/main.go

build:
	go mod download
	go mod vendor
	env GOOS=linux $(BUILDER) -o bin/$(BIN_NAME)-linux $(MAIN)
	env GOOS=windows $(BUILDER) -o bin/$(BIN_NAME)-windows.exe $(MAIN)
	env GOOS=darwin $(BUILDER) -o bin/$(BIN_NAME)-mac $(MAIN)

dev:
	env GOOS=darwin $(BUILDER) -o bin/cmd cmd/main.go	

.PHONY: watch
watch: 
	kill -9 $(ps aux | grep fswatch | grep -v grep | awk '{print $2}') 2>>/dev/null || echo "Prepare..."
	fswatch -o ./cmd/**/*.go ./pkg/**/*.go ./internal/**/*.go | xargs -n1 -I{} make dev
.PHONY: clean
clean:
	rm -rf ./bin ./vendor
