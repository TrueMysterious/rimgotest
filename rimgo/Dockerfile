FROM --platform=$BUILDPLATFORM golang:alpine AS build

ARG TARGETARCH

WORKDIR /src
RUN apk --no-cache add ca-certificates git
COPY . .

RUN go mod download
RUN GOOS=linux GOARCH=$TARGETARCH CGO_ENABLED=0 go build -ldflags "-X codeberg.org/video-prize-ranch/rimgo/pages.VersionInfo=$(date '+%Y-%m-%d')-$(git rev-list --abbrev-commit -1 HEAD)"

FROM scratch as bin

WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/rimgo .

EXPOSE 3000

CMD ["/app/rimgo"]
