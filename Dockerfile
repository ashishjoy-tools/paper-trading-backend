FROM golang:1.19.3-alpine
WORKDIR /go/src/github.com/ashishkujoy/paper-trading-backend/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

FROM amd64/alpine:3.17
WORKDIR /root/
COPY --from=0 /go/src/github.com/ashishkujoy/paper-trading-backend/app ./
CMD ["./app"]