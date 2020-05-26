APP := baja
VERSION ?= $(shell git tag | head -n1)
GIT_COMMIT ?= $(shell git rev-list -1 HEAD)

ldflags = -ldflags "-X main.GitCommit=$(GIT_COMMIT) -X main.AppVersion=$(VERSION)"

build:
	cd cmd && go build -o ../out/$(APP) $(ldflags)

release:
	cd cmd && GOOS=linux go build -o ../out/linux $(ldflags)
	cd cmd && go build -o ../out/mac $(ldflags)

install:
	cp out/$(APP) ~/bin/
