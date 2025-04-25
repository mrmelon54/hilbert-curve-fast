mkdir -p dist/

CGO_ENABLED=0 go build -a -ldflags="-w -s" -o ./dist/hilbert-curve-fast.x86_64 .
CGO_ENABLED=0 GOOS=windows go build -a -ldflags="-w -s" -o ./dist/hilbert-curve-fast.exe .
