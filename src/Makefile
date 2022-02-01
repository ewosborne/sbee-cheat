all: windows mintel m1 linux_x86 linux_rpi

m1:
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o ../bin/sbee-cheat-m1 ./sbee-cheat.go

mintel:
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ../bin/sbee-cheat-mintel ./sbee-cheat.go

windows:
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ../bin/sbee-cheat.exe ./sbee-cheat.go

linux_x86:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ../bin/sbee-cheat-linux-x86 ./sbee-cheat.go

linux_rpi:
	GOOS=linux GOARCH=arm go build -ldflags="-s -w" -o ../bin/sbee-cheat-linux-rpi ./sbee-cheat.go


clean:
	@rm -f ../bin/sbee*
