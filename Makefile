TARGET=./cli ./tests ./parse ./render ./types ./util

test:
	cp -r ./assets/template/ ./tests/assets/ &&\
	gotestsum --format testname -- -short ${TARGET} &&\
	rm -rf ./tests/assets/template ./tests/generated

check:
	go fmt ${TARGET} && \
		go vet ${TARGET} &&\
		cp -r ./assets/template/ ./tests/assets/ &&\
		gotestsum --format testname &&\
		rm -rf ./tests/assets/template ./tests/generated

test-run:
	make build &&\
	./bin/ghi gen wf001 gh-insight-cli

build:
	@rm -rf bin && \
	mkdir bin && \
	cd cli &&\
	go build -o ../bin/ghi . &&\
	cd ..

