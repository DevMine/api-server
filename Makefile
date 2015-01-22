PKG  = github.com/DevMine/api-server
EXEC = devmine

all: check test build

install:
	go install ${PKG}

build:
	go build -o ${EXEC} ${PKG}

test:
	go test -v ${PKG}/...

deps:
	go get -u code.google.com/p/biogo.matrix
	go get -u github.com/golang/glog
	go get -u github.com/gorilla/mux
	go get -u github.com/lib/pq

check:
	go vet ${PKG}/...
	golint ${GOPATH}/src/${PKG}/...

cover:
	go test -cover ${PKG}/...

clean:
	rm -f ./${EXEC}
