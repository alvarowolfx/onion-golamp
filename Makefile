build:
	GOOS=linux GOARCH=mipsle go build -ldflags="-s -w" main.go homekit.go gpio.go
	
copy: 
	scp main root@omega-5d69.local:/root/go