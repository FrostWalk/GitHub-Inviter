FROM golang:alpine AS buildenv

WORKDIR /go/src/build
COPY . .

RUN go mod download
RUN go vet -v

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o inviter

FROM gcr.io/distroless/static

LABEL org.opencontainers.image.source=https://github.com/FrostWalk/GitHub-Inviter
LABEL org.opencontainers.image.description="linux/amd64"
LABEL org.opencontainers.image.licenses=MIT

COPY --from=buildenv /go/src/build/inviter /inviter
COPY --from=buildenv /go/src/build/static/ /static/
COPY --from=buildenv /go/src/build/templates/ /templates/

CMD ["/inviter"]

ENV GITHUB_ORG_NAME=""
ENV GITHUB_TOKEN=""
ENV GITHUB_GROUP_NAME=""
ENV INVITE_CODE_HASH=""
ENV TLS_CERT=""
ENV TLS_KEY=""
ENV HTTP_PORT="80"
ENV HTTPS_PORT="443"
