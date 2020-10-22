# Copyright (C) 2019 Yunify, Inc.

VERBOSE = no

.PHONY: help
help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  format  to format the code"
	@echo "  vet     to run golang vet"
	@echo "  lint    to run the staticcheck"
	@echo "  check   to format, vet, lint"
	@echo "  test    to run test case"
	@exit 0

TEST = _test() {                              \
	args="$(filter-out $@,$(MAKECMDGOALS))";  \
	if [[ $(VERBOSE) = "yes" ]]; then         \
        bash -x scripts/gotest.sh $$args;     \
    else                                      \
        bash scripts/gotest.sh $$args;        \
    fi;                                       \
}

.PHONY: format
format:
	@[[ ${VERBOSE} = "yes" ]] && set -x; go fmt ./...;

.PHONY: vet
vet:
	@[[ ${VERBOSE} = "yes" ]] && set -x; go vet ./...;

.PHONY: lint
lint:
	@[[ ${VERBOSE} = "yes" ]] && set -x; staticcheck ./...;


.PHONY: check
check: format vet lint

.PHONY: test
test:
	@$(TEST); _test

.DEFAULT_GOAL = help

# Target name % means that it is a rule that matches anything, @: is a recipe;
# the : means do nothing
%:
	@:

