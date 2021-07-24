ifeq ($(OS),Windows_NT)
  EXTERNAL=external/program.exe
else
  EXTERNAL=external/program
endif

.PHONY: run
run: $(EXTERNAL)
	go run main.go $(EXTERNAL) $(ARGS)

$(EXTERNAL): $(basename $(EXTERNAL)).c
	gcc -o $@ -Wall $^

.PHONY:
clean:
	rm -f $(EXTERNAL)
