
FRONT_END_BINARY=frontApp
VERSION=1.0.1

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build:  build_front
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_front: builds the front end binary
build_front:
	@echo "Building front end binary..."
	cd ../front-end && go build -o ${FRONT_END_BINARY} cmd/web/*.go && ./${FRONT_END_BINARY}  \
	-dbuser='postgres' \
	-pusherHost='localhost' \
	-pusherKey='abc123' \
	-pusherSecret='123abc' \
	-pusherApp="1"
	-pusherPort="4001"
	-pusherSecure=false 
	@echo "Built front end binary..."

## start: starts the front end
start: build_front
	@echo "Starting front end"
	./${FRONT_END_BINARY} &


## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"

push_front:
	@echo "Pushing front end ..."
	cd ../front-end && docker build -f front-end.dockerfile -t unixjoe/vigilate-front-end:${VERSION} .
	docker push unixjoe/vigilate-front-end:${VERSION}
	@echo "Done!"
