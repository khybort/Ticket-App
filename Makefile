.PHONY: up down build test seed docker-seed up-dev up-prod down-dev down-prod build-dev build-prod restart-dev restart-prod

seed:
	go run cmd/seeder.go

docker-seed:
	docker-compose run api make seed

up-dev:
	docker-compose --profile development up -d

down-dev:
	docker-compose --profile development down

build-dev:
	docker-compose --profile development build

up-prod:
	docker-compose --profile production up -d

down-prod:
	docker-compose --profile production down

build-prod:
	docker-compose --profile production build

restart-dev:
	make build-dev && make down-dev && make up-dev

restart-prod:
	make build-prod && make down-prod && make up-prod

logs:
	docker-compose logs -f

restart:
	docker-compose restart

test:
	docker exec -it api-prod go test -v -cover ./...


cache-clean:
	docker builder prune

clean:
	docker system prune -f

status:
	docker-compose ps

logb-dev:
	docker-compose logs -f api
logf-dev:
	docker-compose logs -f ui

logb-prod:
	docker-compose logs -f api-prod
logf-prod:
	docker-compose logs -f ui-prod