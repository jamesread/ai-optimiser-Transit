default: compile test

compile:
	go build libtransit
	go build transitcli

run:
	./transitcli

clean:
	echo "clean"

test:
	go test libtransit
	go test transitcli

.PHONY: default compile test clean run
