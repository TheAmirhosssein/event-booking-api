# ~~~ Development Environment ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
up: dev-env
down: docker-stop

dev-env:
	@ docker compose up --build

docker-stop:
	@ docker compose down
