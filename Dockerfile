FROM golang:alpine AS buildenv

WORKDIR /go/src/build
COPY . .

RUN go mod download
RUN go vet -v

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o inviter

FROM gcr.io/distroless/static

LABEL org.opencontainers.image.source=https://github.com/FrostWalk/GitHub-Inviter

COPY --from=buildenv /go/src/build/inviter /app/inviter
COPY --from=buildenv /go/src/build/static/ /app/static/
COPY --from=buildenv /go/src/build/templates/ /app/templates/

ENTRYPOINT ["/app/inviter"]

ENV GITHUB_ORG_NAME=""
ENV GITHUB_TOKEN=""
ENV GITHUB_GROUP_NAME=""
ENV INVITE_CODE=""
ENV TLS_CERT=""
ENV TLS_KEY=""
ENV PORT="8080"
