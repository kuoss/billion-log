IMG=ghcr.io/kuoss/log-routine:latest

docker:
	docker build -t $(IMG) . && docker push $(IMG)
	