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

docker-run:
	docker run -it --rm -p 8080:8080 -p 8443:8443 ${IMAGE}:${TAG} --http --tls --clientauth mutual --min-tls-version 1.3
