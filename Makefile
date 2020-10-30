# Binary name
BINARY= getitle
# Builds the project
build:
		go build -ldflags "-s -w" -o ${BINARY} ./src/main/main.go
# Installs our project: copies binaries
install:
		go install
release-upx:
		# Clean
		#go clean
		rm -rf *.gz
		# Build for mac
		go build -ldflags "-s -w" -o ./bin/getitle-mac64-${VERSION} ./src/main/main.go
		upx -9 ./bin/getitle-mac64-${VERSION}
		# Build for linux
		#go clean
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/getitle-linux64-${VERSION} ./src/main/main.go
		upx -9 ./bin/getitle-linux64-${VERSION}
		#go clean
		CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o ./bin/getitle-linux32-${VERSION} ./src/main/main.go
		upx -9 ./bin/getitle-linux32-${VERSION}
		# Build for win
		#go clean
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/getitle-win64-${VERSION}.exe  ./src/main/main.go
		upx -9 ./bin/getitle-win64-${VERSION}.exe
		#go clean
		CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o ./bin/getitle-win32-${VERSION}.exe ./src/main/main.go
		upx -9 ./bin/getitle-win32-${VERSION}.exe
		#compress
		tar cvf getitle.tar.gz bin/*
# Cleans our projects: deletes binaries

release:
# Clean
		#go clean
		rm -rf *.gz
		# Build for mac
		go build -ldflags "-s -w" -o ./bin/getitle-mac64-${VERSION} ./src/main/main.go
		# Build for linux
		#go clean
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/getitle-linux64-${VERSION} ./src/main/main.go
		#go clean
		CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o ./bin/getitle-linux32-${VERSION} ./src/main/main.go
		# Build for win
		#go clean
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/getitle-win64-${VERSION}.exe  ./src/main/main.go
		#go clean
		CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o ./bin/getitle-win32-${VERSION}.exe ./src/main/main.go
		#compress
		tar cvf getitle.tar.gz bin/*

clean:
		go clean

.PHONY:  clean build