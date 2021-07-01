DOCKER_RUN=docker run --rm -it --network host -v /etc/resolv.conf:/etc/resolv.conf -v ${PWD}:/app -w /app golang:1.16-buster
DOCKER_APP_EXC=docker exec account-transaction
DOCKER_MYSQL_EXC=docker exec -it mysql57

configure:
	- ${DOCKER_RUN} go mod download
	- ${DOCKER_RUN} go mod vendor
	- ${DOCKER_RUN} go clean -modcache

build:
	- ${DOCKER_RUN} go build -a -o main .

test:
	- ${DOCKER_RUN} go test -coverprofile cover.out -v ./app/... && go tool cover -html=cover.out -o cover.html

code-review: up
	- ${DOCKER_APP_EXC} go vet ./
	- ${DOCKER_APP_EXC} golint ./...
	- ${DOCKER_APP_EXC} errcheck -blank ./...

up:
	- docker-compose up -d

down:
	- docker-compose down --remove-orphans

migrate: up
	- ${DOCKER_MYSQL_EXC} sh -c "mysql --user=root -p xpto < migrations/xpto.sql"

log:
	- docker logs -f account-transaction