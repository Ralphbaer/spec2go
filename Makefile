build:
	go build -o spec2go ./cmd

help:
	@echo "Management commands"
	@echo ""
	@echo "Usage:"
	@echo "  ## Root Commands"
	@echo "    make build           Build application."
	@echo "    make test            Run tests on all projects."
	@echo "    make lint            Run static code analysis (lint)."
	@echo "    make format          Run code formatter."
	@echo "    make tidy            Run go mod tidy."
	@echo "    make sec             Run gosec."
	@echo ""
	@echo "  ## Utility Commands"
	@echo "    make setup-git-hooks       Setup git hooks."
	@echo ""

test:
	go test -v ./... ./...

#lint:
#	./make.sh "lint"

#format:
#	./make.sh "format"

#setup-git-hooks:
#	./make.sh "setupGitHooks"

tidy:
	go mod tidy

sec:
	gosec ./...