default:
	@echo "=============building Local API============="
	docker build -f Dockerfile -t socket-dispatcher .

up: default
	@echo "=============starting api locally============="
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
	go build -o ./bin/socket-dispatcher ./cmd/daemon/

clean: down
	@echo "=============cleaning up============="
	rm -f api
	docker system prune -f
	docker volume prune -f