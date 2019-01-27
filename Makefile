HAS_PACKR2	:= $(shell command -v packr2)

all: main-packr.go

main-packr.go:
ifndef HAS_PACKR2
	$(error packr2 was not found on path and is needed to make assets)
endif
	GO111MODULE=on packr2 --ignore-imports
