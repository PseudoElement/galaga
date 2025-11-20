build-mac:
	go build -o ./builds/galaga-mac

build-win64:
	GOOS=windows GOARCH=amd64 go build -o builds/galaga-win64.exe

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o builds/galaga-linux-amd64

build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o builds/galaga-linux-arm64

run-mac:
	make build-mac && ./builds/galaga-mac