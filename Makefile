APP := baja
VERSION ?= 0.4.1
GIT_COMMIT ?= $(shell git rev-list -1 HEAD)

ldflags = -ldflags "-X main.GitCommit=$(GIT_COMMIT) -X main.AppVersion=$(VERSION)"

clean:
	rm out/$(APP)

build: clean
	cd cmd && go build -o ../out/$(APP) $(ldflags)

release:
	cd cmd && GOOS=linux go build -o ../out/baja $(ldflags)
	cd out && zip baja-linux.zip baja
	cd cmd && go build -o ../out/baja $(ldflags)
	cd out && zip baja-mac.zip baja

install:
	rm ~/bin/baja
	cp out/$(APP) ~/bin/
