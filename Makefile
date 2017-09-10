build:
	GOOS=linux GOARCH=mipsle go build -ldflags="-s -w" main.go homekit.go lamp.go http.go
	
copy: 
	rsync -P -a main root@omega-5d69.local:/root/go