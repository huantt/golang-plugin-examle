build-plugins:
	$(foreach file, $(wildcard ./plugins/*), go build -buildmode=plugin -o $(file)/speaker.so $(file)/speaker.go;)
