APP := baja

build:
	cd cmd && go build -o ../out/$(APP)

release:
	cd cmd && GOOS=linux go build -o ../out/linux
	cd cmd && go build -o ../out/mac

install:
	cp out/$(APP) ~/bin/
