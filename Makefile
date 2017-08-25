APP := baja

build:
	cd cmd && go build -o ../$(APP)
