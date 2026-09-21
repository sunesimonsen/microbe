default: test

RELEASE_VERSION := $(word 2,$(MAKECMDGOALS))

# Allow `make release v0.3.0` to treat the version as an argument instead of
# as a separate make target.
ifneq ($(RELEASE_VERSION),)
.PHONY: $(RELEASE_VERSION)
$(RELEASE_VERSION):
	@:
endif

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
	VERSION=HEAD Test=true air

cover:
	go test ./... -coverprofile=cover.prof
	@covreport
	@echo
	@echo "Written file://$$PWD/cover.html"

clean:
	rm -rf tmp

.PHONY: release
release:
	@set -eu; \
	if [ "$$(git branch --show-current)" != "main" ]; then \
		echo "release must be run from the main branch" >&2; \
		exit 1; \
	fi; \
	version="$(RELEASE_VERSION)"; \
	if ! printf '%s\n' "$$version" | grep -Eq '^v[0-9]+\.[0-9]+\.[0-9]+$$'; then \
		echo "usage: make release v<major>.<minor>.<patch>" >&2; \
		exit 1; \
	fi; \
	if git rev-parse --verify --quiet "refs/tags/$$version" >/dev/null; then \
		echo "tag $$version already exists" >&2; \
		exit 1; \
	fi; \
	printf '%s\n' "$$version" > VERSION; \
	git add VERSION; \
	git commit -m "Release $$version"; \
	git tag "$$version"; \
	git push origin main "refs/tags/$$version"; \
	$(MAKE) deploy
