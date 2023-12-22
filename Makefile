.DEFAULT_GOAL := run

.PHONY: build
build:
	@docker build -t otus .

.PHONY: run
run:
	@docker run -it --rm --name otus --hostname docker \
	-v $(PWD):/home/otus/otus \
	-v $(PWD)/../postgresql_cluster:/home/otus/postgresql_cluster \
	-v $(PWD)/.config/yandex-cloud:/home/otus/.config/yandex-cloud \
	otus

.PHONY: exec
exec:
	@docker exec -it otus bash
