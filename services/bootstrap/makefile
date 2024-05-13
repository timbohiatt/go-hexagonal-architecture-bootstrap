APP_NAME := hexagonal-go
SVC_NAME := bootstrap
VERSION := v0.0.1-a.1

build: compile

run:
	#source cmd/.env.local && \
	go run cmd/main.go
	
vet: 
	go vet 

docker: docker-build docker-run

docker-build:
	docker build -t localhost/${APP_NAME}/${SVC_NAME}:${VERSION} .

docker-run:
	docker run -p 3000:8080 localhost/${APP_NAME}/${SVC_NAME}:${VERSION}

compile: 
	# Clean Up Previous Binaries
	rm -rf bin/ 

	echo "Building for "
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64  go build -o bin/darwin/${APP_NAME}-${SVC_NAME}-${VERSION}-amd64-darwin cmd/main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64  go build -o bin/darwin/${APP_NAME}-${SVC_NAME}-${VERSION}-arm64-darwin cmd/main.go
	
	# Linux
	CGO_ENABLED=0 GOOS=linux GOARCH=386  go build -o bin/linux/${APP_NAME}-${SVC_NAME}-${VERSION}-386-linux cmd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o bin/linux/${APP_NAME}-${SVC_NAME}-${VERSION}-amd64-linux cmd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm  go build -o bin/linux/${APP_NAME}-${SVC_NAME}-${VERSION}-arm-linux cmd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64  go build -o bin/linux/${APP_NAME}-${SVC_NAME}-${VERSION}-arm64-linux cmd/main.go
	
	# Windows
	CGO_ENABLED=0 GOOS=windows GOARCH=386  go build -o bin/windows/${APP_NAME}-${SVC_NAME}-${VERSION}-386-windows cmd/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -o bin/windows/${APP_NAME}-${SVC_NAME}-${VERSION}-amd64-windows cmd/main.go
 

	