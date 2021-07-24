EXTERNAL=external/program

.PHONY: run
run: $(EXTERNAL)
	go run main.go

$(EXTERNAL): $(EXTERNAL).c
	cc -o $@ -Wall $^

.PHONY:
clean:
	rm -f $(EXTERNAL)
