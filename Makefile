all : 
	go build  matrix.go

test:
	go test

clean: 
	rm -f matrix_tests matrix
