APP := baja

init:
	dep init

deps:
	dep ensure

build: deps
	cd cmd && go build -o ../$(APP)
