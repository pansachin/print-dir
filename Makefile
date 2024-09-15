build::
	go build .

rm::
	rm -f ./print-dir

run-recursive::
	@./print-dir --recursive

run::
	@./print-dir
