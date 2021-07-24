.PHONY: run
run: sample
	go run main.go

sample: sample.c
	cc -o $@ -Wall sample.c
