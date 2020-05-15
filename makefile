all: clean build

clean:
	rm -rf build/

build:
	go build -o build/sugoku main.go

run:
	go run main.go