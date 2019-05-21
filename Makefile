build:
	go mod download
	go mod vendor
	env GOOS=linux go build -ldflags="-s -w" -o bin/cmd cmd/main.go
	env GOOS=windows go build -ldflags="-s -w" -o bin/cmd.exe cmd/main.go

configure:
	brew install go --with-cc-common

dev:
	env GOOS=darwin go build -o bin/cmd cmd/main.go	

.PHONY: watch
watch: 
	kill -9 $(ps aux | grep fswatch | grep -v grep | awk '{print $2}') 2>>/dev/null || echo "Prepare..."
	fswatch -o ./cmd/**/*.go ./pkg/**/*.go ./internal/**/*.go | xargs -n1 -I{} make dev
.PHONY: clean
clean:
	rm -rf ./bin ./vendor
