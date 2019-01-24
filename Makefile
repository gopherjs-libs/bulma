.PHONY: test

s:
	@open http://localhost:8080/github.com/kooksee/vvv/example/rr
	@gopherjs serve -m -q .
