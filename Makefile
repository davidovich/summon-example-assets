SUMMONER_NAME := summon

ASSETS := $(shell find $(SUMMONER_NAME)/assets)

all: bin/$(SUMMONER_NAME)

bin/$(SUMMONER_NAME): $(SUMMONER_NAME)/summon.go $(ASSETS)
	go build -o $@ $<

.PHONY: clean
clean:
	rm -rf bin
