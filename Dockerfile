FROM golang:1.21 AS BUILD
WORKDIR /go/src/app
ARG GIT_ENERGY_USERNAME
ARG GIT_ENERGY_PASSWORD
RUN echo "machine git.energy\nlogin $GIT_ENERGY_USERNAME\npassword $GIT_ENERGY_PASSWORD\n" > ~/.netrc
RUN go env -w GOPRIVATE=.energy
COPY go.mod go.sum ./
RUN go mod download -x
COPY . .
RUN go build -o connector-accountant-service .

FROM debian:bookworm AS FINAL
RUN apt update && apt install -y ca-certificates
WORKDIR /app
RUN groupadd -g 1001 -r corepass && \
        useradd -u 1001 -r -s /bin/false -d /app -g corepass corepass && \
        chown -R corepass:corepass /app
USER corepass:corepass
COPY --from=BUILD --chown=corepass:corepass /go/src/app/connector-accountant-service /app
CMD /app/connector-accountant-service
