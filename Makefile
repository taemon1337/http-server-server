IMAGE := containeryard.evoforge.org/tpstell/http-test-server
TAG := 0.0.1

build:
	go build
	sudo chown -R ${USER}. .

docker:
	docker build -t ${IMAGE}:${TAG} .
	sudo chown -R ${USER}. .

test:
	go test ./...
