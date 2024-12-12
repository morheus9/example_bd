FROM golang:1.23-alpine3.20 AS build

WORKDIR /app

RUN apk add git


ENV APP_NAME="music-library"
ARG COMMIT_HASH="latest"
ARG COMMIT_TIME="latest"
ARG VERSION="dev"

RUN mkdir /out
COPY . /app/

RUN go build  \
    -ldflags "-X github.com/Azaliya1995/music_library/version.Version=${VERSION} -X github.com/Azaliya1995/music_library/version.CommitHash=${COMMIT_HASH} -X github.com/Azaliya1995/music_library/version.CommitTime=${COMMIT_TIME}"  \
    -o /out/${APP_NAME}  \
    github.com/Azaliya1995/music_library/cmd


FROM alpine:3.20

WORKDIR /app

COPY --from=build /out/music-library /app/

EXPOSE 8080
CMD ["/app/music-library", "serve"]