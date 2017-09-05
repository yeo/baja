APP := baja

init:
	dep init

deps:
	dep ensure


build: deps compile install

compile:
	cd cmd && go build -o ../$(APP)

install:
	cp $(APP) ~/bin/
