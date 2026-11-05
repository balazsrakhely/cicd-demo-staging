# Multi-stage build, as in session 1 lesson 04.
# Stage 1 has the Go toolchain. Stage 2 has the binary and nothing else.
FROM golang:1.27 AS build
WORKDIR /src

COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /server .

FROM gcr.io/distroless/static-debian13:nonroot

# Links the GHCR package to its source repository.
LABEL org.opencontainers.image.source=https://github.com/balazsrakhely/cicd-demo

COPY --from=build /server /server

EXPOSE 8080

USER 65532:65532
ENTRYPOINT ["/server"]
