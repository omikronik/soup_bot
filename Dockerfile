# Build
FROM golang:1.22-alpine

RUN apk update && apk add --no-cache gcc libc-dev make

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make build

CMD ["make", "run"]
