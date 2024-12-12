TASK=task

build:
	$(TASK) build

unit-test:
	$(TASK) test

lint:
	$(TASK) lint

install-tools:
	go install github.com/go-task/task/v3/cmd/task@latest

image:
	$(TASK) image