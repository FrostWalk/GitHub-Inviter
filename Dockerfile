FROM golang:alpine as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go vet -v

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o /go/bin/app

FROM gcr.io/distroless/static

COPY --from=build /go/bin/app /
ENTRYPOINT ["/app"]

ENV GITHUB_ORG_NAME=""
ENV GITHUB_TOKEN=""
ENV GITHUB_GROUP_NAME=""
ENV INVITE_CODE=""
ENV TLS_CERT=""
ENV TLS_KEY=""
ENV PORT="8080"
