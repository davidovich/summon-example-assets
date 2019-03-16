SHELL=/bin/bash

HAS_PACKR2	:= $(shell command -v packr2)

SUMMON_NAME := summon
SUMMON_DEPS = $(shell find summon)

ASSETS := $(shell find assets)

all: $(SUMMON_NAME)/packrd/packed-packr.go bin/$(SUMMON_NAME)

bin/$(SUMMON_NAME): $(SUMMON_DEPS)
	@mkdir -p bin
	go build -o $@ $(@F)/summon.go

$(SUMMON_NAME)/packrd/packed-packr.go: $(ASSETS)
ifndef HAS_PACKR2
	$(error packr2 was not found on path and is needed to make assets)
endif
	cd $(SUMMON_NAME) && GO111MODULE=on packr2

.PHONY: clean
clean:
	cd $(SUMMON_NAME) && packr2 clean