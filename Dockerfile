FROM golang AS builder
WORKDIR /app
# Using caches effectively
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -tags netgo -o ./server ./cmd/server

FROM scratch
COPY --from=builder /app/server ./
CMD [ "./server" ]
