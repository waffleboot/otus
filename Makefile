.PHONY: build
build:
	@docker build -t otus .

.PHONY: run
run:
	@docker run -it --rm --name otus -h docker -v $(PWD):/home/otus/otus otus
