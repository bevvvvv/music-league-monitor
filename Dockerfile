FROM golang:1.19-bullseye

WORKDIR /app

# module dependency information
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# does not traverse sub directories
COPY *.go ./
COPY discord_api/*.go ./discord_api/
RUN go build -o music-league-monitor

EXPOSE 80

ENTRYPOINT ["./music-league-monitor"]
