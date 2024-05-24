
.PHONY: all
all: vhs

.PHONY: vhs
vhs: ./docs/vhs.gif

./docs/%.gif: ./docs/%.tape
	vhs $^
