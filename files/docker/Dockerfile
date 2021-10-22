FROM golang:1.17
WORKDIR /app
COPY . .
RUN go mod vendor
RUN go build -o live-chat app/*.go
ENTRYPOINT [ "/app/live-chat" ]
