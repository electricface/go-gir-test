all: bin

bindata.go: ui.glade
	go-bindata ui.glade

bin: bindata.go
	go build -o giteditor
