FROM golang

ADD . /go/src/github.com/harshadnawathe/osohello

ARG VERSION
ARG BUILDTIMESTAMP
ARG GITCOMMIT

RUN go install -ldflags "-X main.Version=$VERSION" github.com/harshadnawathe/osohello

ENTRYPOINT /go/bin/osohello

EXPOSE 8080
