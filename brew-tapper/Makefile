BUILD := $(CURDIR)/_build
DIST := $(CURDIR)/_dist
REGISTRY := softleader

.PHONY: build
build:
	mkdir -p $(BUILD)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BUILD)/tapper -a -tags netgo ./cmd/tapper
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BUILD)/tapper-server -a -tags netgo ./cmd/tapper-server

.PHONY: dist
dist:
	mkdir -p $(DIST)
	cp Dockerfile $(DIST)
	cp -r script $(DIST)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(DIST)/tapper-server -a -tags netgo ./cmd/tapper-server
	docker build -t $(REGISTRY)/brew-tapper $(DIST)
	docker push $(REGISTRY)/brew-tapper

.PHONY: clean
clean:
	rm -rf _*