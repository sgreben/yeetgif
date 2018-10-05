.PHONY: build push all README.md clean

APP=gif
NAME := yeetgif
REPOSITORY := quay.io/sergey_grebenshchikov/$(NAME)
VERSION := 1.0.10
VERSION_COMMIT := $(VERSION)-$(shell printf "%s" "$$(git rev-parse HEAD)")

PACKAGES := $(shell go list -f {{.Dir}} ./...)
GOFILES  := $(addsuffix /*.go,$(PACKAGES))
GOFILES  := $(wildcard $(GOFILES))

build: Dockerfile $(GOFILES)
	docker build -t $(NAME) -t $(REPOSITORY):latest -t $(REPOSITORY):$(VERSION) .

push: build
	docker push $(REPOSITORY):latest
	docker push $(REPOSITORY):$(VERSION)

all: build push

clean:
	rm -rf binaries zip

# go get -u github.com/github/hub
release: zip README.md
	git reset
	dep ensure
	git add vendor
	git add Gopkg.toml Gopkg.lock
	git commit -m "dep ensure" || true
	git reset
	git add README.template.md README.md Makefile
	git commit -m "Release $(VERSION)" || true
	git push
	hub release create $(VERSION) -m "$(VERSION)" -a release/$(APP)_$(VERSION)_osx_x86_64.tar.gz -a release/$(APP)_$(VERSION)_windows_x86_64.zip -a release/$(APP)_$(VERSION)_linux_x86_64.tar.gz -a release/$(APP)_$(VERSION)_osx_x86_32.tar.gz -a release/$(APP)_$(VERSION)_windows_x86_32.zip -a release/$(APP)_$(VERSION)_linux_x86_32.tar.gz -a release/$(APP)_$(VERSION)_linux_arm64.tar.gz

README.md:
	sed "s/\$${VERSION}/$(VERSION)/g;s/\$${APP}/$(APP)/g;" README.template.md > README.md

zip: release/$(APP)_$(VERSION)_osx_x86_64.tar.gz release/$(APP)_$(VERSION)_windows_x86_64.zip release/$(APP)_$(VERSION)_linux_x86_64.tar.gz release/$(APP)_$(VERSION)_osx_x86_32.tar.gz release/$(APP)_$(VERSION)_windows_x86_32.zip release/$(APP)_$(VERSION)_linux_x86_32.tar.gz release/$(APP)_$(VERSION)_linux_arm64.tar.gz

binaries: binaries/osx_x86_64/$(APP) binaries/windows_x86_64/$(APP).exe binaries/linux_x86_64/$(APP) binaries/osx_x86_32/$(APP) binaries/windows_x86_32/$(APP).exe binaries/linux_x86_32/$(APP)

release/$(APP)_$(VERSION)_osx_x86_64.tar.gz: binaries/osx_x86_64/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_osx_x86_64.tar.gz -C binaries/osx_x86_64 $(APP)

binaries/osx_x86_64/$(APP): $(GOFILES)
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION_COMMIT)" -o binaries/osx_x86_64/$(APP) ./cmd/$(APP)

release/$(APP)_$(VERSION)_windows_x86_64.zip: binaries/windows_x86_64/$(APP).exe
	mkdir -p release
	cd ./binaries/windows_x86_64 && zip -r -D ../../release/$(APP)_$(VERSION)_windows_x86_64.zip $(APP).exe

binaries/windows_x86_64/$(APP).exe: $(GOFILES)
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION_COMMIT)" -o binaries/windows_x86_64/$(APP).exe ./cmd/$(APP)

release/$(APP)_$(VERSION)_linux_x86_64.tar.gz: binaries/linux_x86_64/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_linux_x86_64.tar.gz -C binaries/linux_x86_64 $(APP)

binaries/linux_x86_64/$(APP): $(GOFILES)
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION_COMMIT)" -o binaries/linux_x86_64/$(APP) ./cmd/$(APP)

release/$(APP)_$(VERSION)_osx_x86_32.tar.gz: binaries/osx_x86_32/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_osx_x86_32.tar.gz -C binaries/osx_x86_32 $(APP)

binaries/osx_x86_32/$(APP): $(GOFILES)
	GOOS=darwin GOARCH=386 go build -ldflags "-X main.version=$(VERSION_COMMIT)" -o binaries/osx_x86_32/$(APP) ./cmd/$(APP)

release/$(APP)_$(VERSION)_windows_x86_32.zip: binaries/windows_x86_32/$(APP).exe
	mkdir -p release
	cd ./binaries/windows_x86_32 && zip -r -D ../../release/$(APP)_$(VERSION)_windows_x86_32.zip $(APP).exe

binaries/windows_x86_32/$(APP).exe: $(GOFILES)
	GOOS=windows GOARCH=386 go build -ldflags "-X main.version=$(VERSION_COMMIT)" -o binaries/windows_x86_32/$(APP).exe ./cmd/$(APP)

release/$(APP)_$(VERSION)_linux_x86_32.tar.gz: binaries/linux_x86_32/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_linux_x86_32.tar.gz -C binaries/linux_x86_32 $(APP)

binaries/linux_x86_32/$(APP): $(GOFILES)
	GOOS=linux GOARCH=386 go build -ldflags "-X main.version=$(VERSION_COMMIT)" -o binaries/linux_x86_32/$(APP) ./cmd/$(APP)

release/$(APP)_$(VERSION)_linux_arm64.tar.gz: binaries/linux_arm64/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_linux_arm64.tar.gz -C binaries/linux_arm64 $(APP)

binaries/linux_arm64/$(APP): $(GOFILES)
	GOOS=linux GOARCH=arm64 go build -ldflags "-X main.version=$(VERSION_COMMIT)" -o binaries/linux_arm64/$(APP) ./cmd/$(APP)
