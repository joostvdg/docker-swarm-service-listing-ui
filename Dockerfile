FROM golang:latest AS build
WORKDIR /src
ENV LAST_UPDATE=20171020
RUN go get -v github.com/gorilla/mux/...
ADD . /src
RUN go build -v -tags netgo -o docker-swarm-service-listing-ui

FROM alpine:3.6
MAINTAINER 	Joost van der Griendt <joostvdg@gmail.com>
CMD ["docker-swarm-service-listing-ui"]
HEALTHCHECK --interval=5s --start-period=3s --timeout=5s CMD wget -qO- "http://localhost:8087/"
COPY --from=build /src/docker-swarm-service-listing-ui /usr/local/bin/docker-swarm-service-listing-ui
RUN chmod +x /usr/local/bin/docker-swarm-service-listing-ui