debug:
	cp ./api/.realize.debug.yaml ./api/.realize.yaml
	docker-compose up -d

dev:
	cp ./api/.realize.develop.yaml ./api/.realize.yaml
	docker-compose up -d
