PACKAGES=$(shell go list ./...)
PACKAGES_UNITTEST=$(shell go list ./... | grep -v integration_test)
export GO111MODULE = on


proto-gen:
	@./third_party/protocgen.sh