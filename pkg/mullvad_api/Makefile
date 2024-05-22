API_BASE := https://api.mullvad.net
OPENAPI_GEN_PROPS := isGoSubmodule=true,packageName=mullvad_api,withGoMod=false,generateInterfaces=true
GET_SPEC_PATH = $(shell curl 'https://api.mullvad.net/app/documentation/' | grep spec-url | awk -F "'" '{print $$2}')
SET_SPEC_PATH = $(eval SPEC_PATH=$(GET_SPEC_PATH))
VERSION := "0.0.1"
UAGENT = "mully/$(VERSION)"

all :: clean dep fetch gen prune

dep ::
	if ! openapi-generator-cli version 2>&1 > /dev/null; then npm install @openapitools/openapi-generator-cli -g; fi

fetch ::
	mkdir -p api
	$(SET_SPEC_PATH)
	SPEC_URL="$(API_BASE)$(SPEC_PATH)"; \
		curl "$$SPEC_URL" -o api/mullvad_api.yaml;

gen ::
	openapi-generator-cli generate \
		-g go \
		-i api/mullvad_api.yaml \
		 --package-name mullvad_api \
		 --api-package mullvad_api \
		 --http-user-agent $(UAGENT) \
		 -p $(OPENAPI_GEN_PROPS);

prune ::
	rm .openapi-generator-ignore 2>/dev/null || true
	rm .gitignore 2>/dev/null || true
	rm .travis.yml 2>/dev/null || true
	rm git_push.sh 2>/dev/null || true
	rm openapitools.json 2>/dev/null || true
	rm -r .openapi-generator 2>/dev/null || true

clean ::
	rm -r api 2>/dev/null || true
	rm -r docs 2>/dev/null || true
	rm -r test 2>/dev/null || true
	rm *.go 2>/dev/null || true
	rm README.md 2>/dev/null || true
