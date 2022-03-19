FROM golang:alpine AS build

WORKDIR /src
RUN apk --no-cache add git ca-certificates
RUN git clone https://codeberg.org/video-prize-ranch/rimgo .

RUN go mod download
RUN CGO_ENABLED=0 go build

FROM scratch as bin

WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/rimgo .

EXPOSE 3000

CMD ["/app/rimgo"]
