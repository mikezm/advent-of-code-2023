# Makefile

ifeq ($(d),)
    $(error Please provide a value for d. Example: make gen d=1)
endif

all: gen

gen: ./scripts/gen.sh
	bash $< $(d)

clean:
	rm -rf "day$(d)"

.PHONY: run
