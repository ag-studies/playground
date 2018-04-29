FROM golang:1.10

COPY . /go/

RUN export GOPATH=$GOPATH:/go/

RUN cd /go/src/mydata && go build dataunit.go
RUN cd /go/src/app && go build modmain.go
RUN cd /go/src/app && go build -buildmode=plugin -o modcounter.so modcounter.go

WORKDIR /go/src/app

RUN ls -l

CMD ["/go/src/app/modmain"]