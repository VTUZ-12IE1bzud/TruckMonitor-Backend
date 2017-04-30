FROM golang
MAINTAINER Pavel Annin  <anninpavel2014@gmail.com>

ADD . /go/src/TruckMonitor-Backend/

RUN go get gopkg.in/gin-gonic/gin.v1
RUN go get github.com/lib/pq
RUN go get github.com/dgrijalva/jwt-go

RUN go install gopkg.in/gin-gonic/gin.v1
RUN go install github.com/lib/pq
RUN go install github.com/dgrijalva/jwt-go

EXPOSE 8080
WORKDIR /go/src/TruckMonitor-Backend

ENTRYPOINT ["go"]
CMD ["run", "application.go"]