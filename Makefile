
all:
	go build -o build/inspect inspect.go

install:
	cp build/inspect ${HOME}/bin

