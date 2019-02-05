APP := baja

build:
	cd cmd && go build -o ../out/$(APP)

install:
	cp out/$(APP) ~/bin/
