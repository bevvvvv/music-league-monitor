FROM golang:1.19-bullseye

WORKDIR /app

# module dependency information
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# does not traverse sub directories
COPY *.go ./
RUN go build -o /music-league-monitor

EXPOSE 80

CMD /music-league-monitor