
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 
	go build -C ./src -o ./../url-shortner ./cmd/ 

dockerize:
	docker build -t url-shortner .

run-docker:
	docker run -p 8080:8080 url-shortner
	