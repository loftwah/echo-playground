# Start with the latest Golang base image for building the app
FROM golang:latest as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a minimal alpine image for the final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .

# Copy the .env file into the container
COPY .env .
COPY data/data.csv data/data.csv

EXPOSE 1323
CMD ["./main"]
