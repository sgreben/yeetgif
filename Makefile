.PHONY: build push all README.md clean binary

APP=gif
NAME := yeetgif
REPOSITORY := quay.io/sergey_grebenshchikov/$(NAME)
VERSION := 1.10.0
VERSION_COMMIT := $(VERSION)-$(shell printf "%s" "$$(git rev-parse HEAD)" | cut -c 1-8)

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

README.md: binary examples
	<README.template.md subst USAGE_roll="$$($(APP) roll -h 2>&1)" USAGE_wobble="$$($(APP) wobble -h 2>&1)" USAGE_pulse="$$($(APP) pulse -h 2>&1)" USAGE_zoom="$$($(APP) zoom -h 2>&1)" USAGE_shake="$$($(APP) shake -h 2>&1)" USAGE_woke="$$($(APP) woke -h 2>&1)" USAGE_fried="$$($(APP) fried -h 2>&1)" USAGE_hue="$$($(APP) hue -h 2>&1)" USAGE_tint="$$($(APP) tint -h 2>&1)" USAGE_resize="$$($(APP) resize -h 2>&1)" USAGE_crop="$$($(APP) crop -h 2>&1)" USAGE_optimize="$$($(APP) optimize -h 2>&1)" USAGE_compose="$$($(APP) compose -h 2>&1)" USAGE_crowd="$$($(APP) crowd -h 2>&1)" USAGE_nop="$$($(APP) nop -h 2>&1)" USAGE_meta="$$($(APP) meta -h 2>&1)" USAGE="$$($(APP) -h 2>&1)" VERSION="$(VERSION)" APP="$(APP)" >README.md

examples: doc/terminal.gif doc/roll.gif doc/wobble.gif doc/pulse.gif doc/zoom.gif doc/shake.gif doc/woke.gif doc/fried.gif doc/hue.gif doc/tint.gif doc/compose.gif doc/crowd.gif

doc/terminal.gif: Makefile doc/wobble.gif
	<doc/terminal.png gif -n 5 fried -j 0 -a 0 -t 0 -u 1 -o 1 -n 0.7 | gif compose -s 1.0 -p right doc/wobble.gif | gif optimize -x 0 -y 0 --kb=999 > doc/terminal.gif

doc/roll.gif: Makefile
	<doc/eggplant.png gif roll > doc/roll.gif

doc/wobble.gif: Makefile
	<doc/eggplant.png gif -n 30 wobble -a 15 > doc/wobble.gif

doc/pulse.gif: Makefile
	<doc/eggplant.png gif pulse > doc/pulse.gif

doc/zoom.gif: Makefile
	<doc/eggplant.png gif zoom > doc/zoom.gif

doc/shake.gif: Makefile
	<doc/eggplant.png gif shake > doc/shake.gif

doc/woke.gif: Makefile
	<doc/yeet.png gif woke  "[[32,60],[92,60]]" > doc/woke.gif

doc/fried.gif: Makefile
	<doc/yeet.png gif fried > doc/fried.gif

doc/hue.gif: Makefile
	<doc/eggplant.png gif hue > doc/hue.gif

doc/tint.gif: Makefile
	<doc/eggplant.png gif tint > doc/tint.gif

doc/compose.gif: Makefile
	<doc/yeet.png gif compose -p left -s 0.8 -x 50 doc/wobble.gif > doc/compose.gif

doc/crowd.gif: Makefile
	<doc/wobble.gif gif crowd > doc/crowd.gif

binary:
	go install ./cmd/$(APP)

zip: release/$(APP)_$(VERSION)_osx_x86_64.tar.gz release/$(APP)_$(VERSION)_windows_x86_64.zip release/$(APP)_$(VERSION)_linux_x86_64.tar.gz release/$(APP)_$(VERSION)_osx_x86_32.tar.gz release/$(APP)_$(VERSION)_windows_x86_32.zip release/$(APP)_$(VERSION)_linux_x86_32.tar.gz release/$(APP)_$(VERSION)_linux_arm64.tar.gz

binaries: binaries/osx_x86_64/$(APP) binaries/windows_x86_64/$(APP).exe binaries/linux_x86_64/$(APP) binaries/osx_x86_32/$(APP) binaries/windows_x86_32/$(APP).exe binaries/linux_x86_32/$(APP)

release/$(APP)_$(VERSION)_osx_x86_64.tar.gz: binaries/osx_x86_64/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_osx_x86_64.tar.gz -C binaries/osx_x86_64 $(APP)

binaries/osx_x86_64/$(APP): $(GOFILES) Makefile
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION_COMMIT)" -o binaries/osx_x86_64/$(APP) ./cmd/$(APP)

release/$(APP)_$(VERSION)_windows_x86_64.zip: binaries/windows_x86_64/$(APP).exe
	mkdir -p release
	cd ./binaries/windows_x86_64 && zip -r -D ../../release/$(APP)_$(VERSION)_windows_x86_64.zip $(APP).exe

binaries/windows_x86_64/$(APP).exe: $(GOFILES) Makefile
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION_COMMIT)" -o binaries/windows_x86_64/$(APP).exe ./cmd/$(APP)

release/$(APP)_$(VERSION)_linux_x86_64.tar.gz: binaries/linux_x86_64/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_linux_x86_64.tar.gz -C binaries/linux_x86_64 $(APP)

binaries/linux_x86_64/$(APP): $(GOFILES) Makefile
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION_COMMIT)" -o binaries/linux_x86_64/$(APP) ./cmd/$(APP)

release/$(APP)_$(VERSION)_osx_x86_32.tar.gz: binaries/osx_x86_32/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_osx_x86_32.tar.gz -C binaries/osx_x86_32 $(APP)

binaries/osx_x86_32/$(APP): $(GOFILES) Makefile
	GOOS=darwin GOARCH=386 go build -ldflags "-X main.version=$(VERSION_COMMIT)" -o binaries/osx_x86_32/$(APP) ./cmd/$(APP)

release/$(APP)_$(VERSION)_windows_x86_32.zip: binaries/windows_x86_32/$(APP).exe
	mkdir -p release
	cd ./binaries/windows_x86_32 && zip -r -D ../../release/$(APP)_$(VERSION)_windows_x86_32.zip $(APP).exe

binaries/windows_x86_32/$(APP).exe: $(GOFILES) Makefile
	GOOS=windows GOARCH=386 go build -ldflags "-X main.version=$(VERSION_COMMIT)" -o binaries/windows_x86_32/$(APP).exe ./cmd/$(APP)

release/$(APP)_$(VERSION)_linux_x86_32.tar.gz: binaries/linux_x86_32/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_linux_x86_32.tar.gz -C binaries/linux_x86_32 $(APP)

binaries/linux_x86_32/$(APP): $(GOFILES) Makefile
	GOOS=linux GOARCH=386 go build -ldflags "-X main.version=$(VERSION_COMMIT)" -o binaries/linux_x86_32/$(APP) ./cmd/$(APP)

release/$(APP)_$(VERSION)_linux_arm64.tar.gz: binaries/linux_arm64/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_linux_arm64.tar.gz -C binaries/linux_arm64 $(APP)

binaries/linux_arm64/$(APP): $(GOFILES) Makefile
	GOOS=linux GOARCH=arm64 go build -ldflags "-X main.version=$(VERSION_COMMIT)" -o binaries/linux_arm64/$(APP) ./cmd/$(APP)
