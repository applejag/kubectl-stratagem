
.PHONY: all
all: vhs

.PHONY: vhs
vhs: ./docs/vhs.gif

./docs/%.gif: kubectl-stratagem ./docs/%.tape
	vhs $(filter-out kubectl-stratagem, $^)

kubectl-stratagem: $(wildcard **.go) go.mod go.sum
	go build
