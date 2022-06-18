.PHONY: database
database:
	docker-compose up database

.PHONY: api
api:
	docker-compose up api

.PHONY: express
express:
	docker-compose up mongo-express

.PHONY: down
down:
	docker-compose down

