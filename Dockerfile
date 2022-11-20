FROM golang:alpine AS build-env
WORKDIR /urban
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
COPY go.mod /urban/go.mod
COPY go.sum /urban/go.sum
RUN go mod download
COPY . /urban
RUN CGO_ENABLED=0 GOOS=linux go build -o build/urban ./urban


FROM scratch
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /urban/build/urban /
ENTRYPOINT ["/urban"]
CMD ["server"]
