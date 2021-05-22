build:
	go build .
	chmod +x wpscan
	sudo mv wpscan /usr/local/bin