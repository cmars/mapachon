FROM docker.io/golang:1.18-bullseye AS build
ENV GOPATH /build
WORKDIR /build/src/github.com/cmars/mapachon
COPY go.* /build/src/github.com/cmars/mapachon/
RUN go mod download
COPY . /build/src/github.com/cmars/mapachon/
RUN go install .

FROM gcr.io/distroless/base-debian11
COPY --from=build /build/bin/mapachon /mapachon
CMD ["/mapachon"]
