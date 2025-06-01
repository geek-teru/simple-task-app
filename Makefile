.PHONY: run

run:
	docker compose up -d --build

.PHONY: dt

dt:
	docker exec -it postgres.local psql -U admin -d sampledb -c "\dt"

.PHONY: healthcheck

healthcheck:
	curl -X GET http://localhost:8080/healthcheck

.PHONY: migrate

migrate:
	docker compose run --rm migrate bash -c'go run cmd/main.go migrate'
