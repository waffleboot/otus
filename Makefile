.DEFAULT_GOAL := run

.PHONY: build
build:
	@docker build -t otus .

.PHONY: run
run:
	@docker run -it --rm --name otus -h docker -v $(PWD):/home/otus/otus -v $(PWD)/.config/yandex-cloud:/home/otus/.config/yandex-cloud otus
