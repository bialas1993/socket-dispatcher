default:
	@echo "============= building ============="
	docker build -f Dockerfile -t socket-dispatcher .

up: default
	@echo "============= starting ============="
	docker-compose up -d

logs:
	docker-compose logs -f

down:
	docker-compose down

shell:
	docker exec -it socket-dispatcher /bin/bash

test:
	go test -v -cover ./...

build:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build  -ldflags "-s -w" -o ./bin/socket-dispatcher ./cmd/daemon/

clean: down
	@echo "============= cleaning up ============="
	docker system prune -f
	docker volume prune -f