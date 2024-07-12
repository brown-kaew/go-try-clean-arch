start-db:
	docker-compose -f docker-compose.yml up --build db

run-sandbox:
	docker-compose -f docker-compose.yml up --build app
