.PHONY: all test clean glide fast

GO15VENDOREXPERIMENT=1

all: alice

alice: deps test
	go build -o ./build/alice ./main.go

fast:
	go build -i -o ./build/alice ./main.go

deps: glide
	./glide install

glide:
ifeq ($(shell uname),Darwin)
	curl -L https://github.com/Masterminds/glide/releases/download/0.5.0/glide-darwin-amd64.zip -o glide.zip
	unzip glide.zip
	mv ./darwin-amd64/glide ./glide
	rm -fr ./darwin-amd64
	rm ./glide.zip
else
	curl -L https://github.com/Masterminds/glide/releases/download/0.5.0/glide-linux-386.zip -o glide.zip
	unzip glide.zip
	mv ./linux-386/glide ./glide
	rm -fr ./linux-386
	rm ./glide.zip
endif

test:
	go test ./...

clean:
	rm ./glide
	rm -fr ./build