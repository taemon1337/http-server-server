IMAGE := ${DOCKER_REPO}/tpstell/http-test-server
TAG := 0.0.1

build:
	go build
	sudo chown -R ${USER}. .

docker:
	docker build -t ${IMAGE}:${TAG} .
	sudo chown -R ${USER}. .

push:
	docker push ${IMAGE}:${TAG}

test:
	go test ./...
