.PHONY: output.txt
output.txt:
	go test -bench=. | tee output.txt
