DOCKER_RUN=docker run --rm -it --network host -v /etc/resolv.conf:/etc/resolv.conf -v ${PWD}:/app -w /app golang:1.16-stretch
DOCKER_APP_EXC=docker exec go-account-transaction
DOCKER_MYSQL_EXC=docker exec -it mysql57

configure:
	- ${DOCKER_RUN} go mod download
	- ${DOCKER_RUN} go mod vendor
	- ${DOCKER_RUN} go clean -modcache

build:
	- ${DOCKER_RUN} go build -a -o main .

code-review:up
	-  ${DOCKER_EXC} go vet ./
	-  ${DOCKER_EXC} golint ./...
	-  ${DOCKER_EXC} errcheck -blank ./...

up:
	- docker-compose up -d

down:
	- docker-compose down --remove-orphans

migrate: up
	- ${DOCKER_MYSQL_EXC} sh -c "mysql --user=root -p xpto < migrations/xpto.sql"