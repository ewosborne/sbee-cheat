all: windows mintel m1

m1:
	GOOS=darwin GOARCH=arm64 go build -o ../bin/sbee-cheat-m1 ./sbee-cheat.go

mintel:
	GOOS=darwin GOARCH=amd64 go build -o ../bin/sbee-cheat-mintel ./sbee-cheat.go

windows:
	GOOS=windows GOARCH=amd64 go build -o ../bin/sbee-cheat.exe ./sbee-cheat.go

clean:
	@rm -f ../bin/sbee*
