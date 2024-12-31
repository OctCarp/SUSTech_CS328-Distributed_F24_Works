.PHONY: build-old build-new build-all

OLD_VERSION:=1.0.0
NEW_VERSION:=1.0.1

build-old:
	docker build --build-arg VERSION=$(OLD_VERSION) -t a4-flask:$(OLD_VERSION) .

build-new:
	docker build --build-arg VERSION=$(NEW_VERSION) -t a4-flask:$(NEW_VERSION) .

build-all: build-old build-new