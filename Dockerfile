# build stage
FROM golang:alpine AS build-env
MAINTAINER atoato88@gmail.com
COPY src/* /src/
RUN cd /src && go build -o goapp


# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/goapp /app/
ENTRYPOINT ./goapp
