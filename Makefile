CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd

IMG_NAME=${APP}
REGISTRY=${REGISTRY}
TAG=latest
ENV_TAG=latest
NETWORK_NAME=default
PROJECT_NAME=learn-cloud

# Including
include .build_info

run:
	go run cmd/main.go

deploy:
	gcloud app deploy --project ${PROJECT_NAME} --version ${ENV_TAG} --no-promote

build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

clear:
	rm -rf ${CURRENT_DIR}/bin/*

network:
	docker network create --driver=bridge ${NETWORK_NAME}

migrate-up:
	docker run --mount type=bind,source="${CURRENT_DIR}/migrations,target=/migrations" --network ${NETWORK_NAME} migrate/migrate \
		-path=/migrations/ -database=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable up

migrate-down:
	docker run --mount type=bind,source="${CURRENT_DIR}/migrations,target=/migrations" --network ${NETWORK_NAME} migrate/migrate \
		-path=/migrations/ -database=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable down

mark-as-production-image:
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:production
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:production

build-image:
	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:${TAG} .
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:${ENV_TAG}

push-image:
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:${TAG}
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:${ENV_TAG}
swag-init:
	swag init -g cmd/main.go -o docs

.PHONY: proto
