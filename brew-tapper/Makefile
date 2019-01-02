BUILD := $(CURDIR)/_build
REGISTRY := softleader

.PHONY: build
build:
	mkdir -p $(BUILD)
	cp Dockerfile $(BUILD)
	cp -r script $(BUILD)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BUILD)/tapper -a -tags netgo ./cmd/tapper
	tar -C $(BUILD) -zcvf $(BUILD)/tapper.tgz tapper
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BUILD)/tapper-server -a -tags netgo ./cmd/tapper-server
	docker build -t $(REGISTRY)/brew-tapper $(BUILD)

.PHONY: dist
dist: build
	docker push $(REGISTRY)/brew-tapper

.PHONY: clean
clean:
	rm -rf _*