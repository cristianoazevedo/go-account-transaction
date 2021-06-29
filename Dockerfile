FROM golang:1.16 AS PROD
WORKDIR /app
COPY . /app

FROM PROD AS DEV
WORKDIR /app
COPY . /app

RUN go get github.com/pilu/fresh \
    && go get -u golang.org/x/lint/golint \
    && go get -u github.com/kisielk/errcheck

EXPOSE 3001
ENTRYPOINT ["fresh", "-c", "runner.conf"]