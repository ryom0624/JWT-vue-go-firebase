debug:
	cp ./jwt/.realize.debug.yaml ./jwt/.realize.yaml
	cp ./session/.realize.debug.yaml ./session/.realize.yaml
	docker-compose up -d

dev:
	cp ./jwt/.realize.develop.yaml ./jwt/.realize.yaml
	cp ./session/.realize.develop.yaml ./session/.realize.yaml
	docker-compose up -d
