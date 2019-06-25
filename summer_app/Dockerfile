FROM golang

RUN mkdir -p /go/src/github.com/jadczakd/summer-camp/

ADD . /go/src/github.com/jadczakd/summer-camp/
WORKDIR /go/src/github.com/jadczakd/summer-camp/summer_app

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/jadczakd/summer-camp/summer_app/cmd  -watch github.com/jadczakd/summer-camp/summer_app
