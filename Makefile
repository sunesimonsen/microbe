default: test

tmp/main: **/*.go
	go build -o tmp/main

build: tmp/main

deploy:
	git push dokku main:master

browse:
	open "https://microbe.sune.one/"

run:
	go run .

test:
	go test ./...

test-update:
	UPDATE_SNAPS=true go test ./...

dev:
	Test=true air

cover:
	go test ./... -coverprofile=cover.prof
	@covreport
	@echo
	@echo "Written file://$$PWD/cover.html"

clean:
	rm -rf tmp
