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
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o ./bin/socket-dispatcher ./cmd/daemon/

clean: down
	@echo "============= cleaning up ============="
	docker system prune -f
	docker volume prune -f