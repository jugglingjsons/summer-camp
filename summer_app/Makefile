build:
		GOOS=linux go build -o summer_app ./cmd/main.go
		docker build -f Dockerfile_build -t gcr.io/${TF_VAR_project}/summer_app:latest .
		docker build -f Dockerfile_build -t  ${DOCKER_HUB_USER}/summer-app:latest .
		docker push gcr.io/${TF_VAR_project}/summer_app:latest
		docker push ${DOCKER_HUB_USER}/summer-app:latest
		rm -f summer_app