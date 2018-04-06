FROM golang:1.9.2

### START: Setting Environment ###

	ENV GOPATH /go
	ENV GOENV prod
	ENV PATH $GOPATH/bin:$PATH
    ENV GOAPP WhereIsMyDriver
    ENV PORT 3001
    ENV DB_NAME where_is_my_driver
    ENV DB_HOST localhost
    ENV DB_USER root
    ENV DB_PASSWORD root
    ENV DB_PORT=3306

### END: Setting Environment ###

### START: Set Date Time ###

	RUN ln -sf /usr/share/zoneinfo/Asia/Bangkok /etc/localtime
	RUN echo "Asia/Jakarta" > /etc/timezone && dpkg-reconfigure -f noninteractive tzdata
	ENV TZ=Asia/Jakarta
	RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

### END Set Date Time ###

### START: add source ###

	RUN mkdir -p /go/src/WhereIsMyDriver

	ADD . /go/src/WhereIsMyDriver

### END: add source ###


### START: Initialize dependency ###

	RUN go get -u github.com/kataras/iris/...

	RUN go get ./...

### END: Initialize dependency ###

CMD ["go","run","main.go"]