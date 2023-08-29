PROGRAM_NAME=go-opengl-template

build:
	GOARCH=amd64 GOOS=windows go build -o bin/${PROGRAM_NAME}-windows.exe main.go

run: 
	build
	bin/${PROGRAM_NAME}-windows.exe

clean:
	go clean
	rm bin/${PROGRAM_NAME}-windows.exe