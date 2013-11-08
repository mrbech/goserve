all: goinstall

OBJS = main.go log.go routes.go filehandler.go

goinstall: ${OBJS}
	go get
	go install

install: ${OBJS}
	go build
	cp goserve /usr/bin

build: ${OBJS}
	go build

maninstall: man/goserve.1
	cp man/goserve.1 /usr/share/man/man1/goserve.1
	gzip /usr/share/man/man1/goserve.1

clean:
	go clean

uninstall:
	rm /usr/share/man/man1/goserve.1.gzip
	rm /usr/bin/goserve
