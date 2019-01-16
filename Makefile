all: push

BUILDTAGS=

# Use the 0.0.0 tag for testing, it shouldn't clobber any release builds
APP?=k8s-introduction
USERSPACE?=andrewozhegov
CHARTS?=k8s-introcharts
RELEASE?=0.1.0
PROJECT?=github.com/${USERSPACE}/${APP}
GOOS?=linux
SERVICE_PORT?=8000

CONTAINER_IMAGE?=${USERSPACE}/${APP}

REPO_INFO=$(shell git config --get remote.origin.url)

ifndef COMMIT
	COMMIT := git-$(shell git rev-parse --short HEAD)
endif

vendor: clean
	go get -u github.com/Masterminds/glide \
	&& glide install

build: vendor
	CGO_ENABLED=0 GOOS=${GOOS} go build -a -installsuffix cgo \
		-ldflags "-s -w -X ${PROJECT}/version.RELEASE=${RELEASE} -X ${PROJECT}/version.COMMIT=${COMMIT} -X ${PROJECT}/version.REPO=${REPO_INFO}" \
		-o ${APP}

container: build
	docker build --pull -t $(CONTAINER_IMAGE):$(RELEASE) .

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)

run: container
	docker run --name ${APP} -p ${SERVICE_PORT}:${SERVICE_PORT} \
		-e "SERVICE_PORT=${SERVICE_PORT}" \
		-d $(CONTAINER_IMAGE):$(RELEASE)

deploy: push
	for t in $(shell find ./kubernetes/${APP}/ -type f -name "*.yaml"); do \
    cat $$t | \
        gsed -E "s/\{\{(\s*)\.Release(\s*)\}\}/$(RELEASE)/g" | \
        gsed -E "s/\{\{(\s*)\.ServiceName(\s*)\}\}/$(APP)/g"; \
    echo ---; \
    done > tmp.yaml
	kubectl apply -f tmp.yaml

fmt:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"gofmt -s -l {{.Dir}}"{{end}}' $(shell go list ${PROJECT}/... | grep -v vendor) | xargs -L 1 sh -c

lint:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"golint {{.Dir}}/..."{{end}}' $(shell go list ${PROJECT}/... | grep -v vendor) | xargs -L 1 sh -c

vet:
	@echo "+ $@"
	@go vet $(shell go list ${PROJECT}/... | grep -v vendor)

test: vendor fmt lint vet
	@echo "+ $@"
	@go test -v -race -tags "$(BUILDTAGS) cgo" $(shell go list ${PROJECT}/... | grep -v vendor)

cover:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"{{end}}' $(shell go list ${PROJECT}/... | grep -v vendor) | xargs -L 1 sh -c

clean:
	rm -f ${APP}
