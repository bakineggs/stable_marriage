stable_marriage: stable_marriage.c
	gcc -o $@ $^

stable_marriage.c: stable_marriage.bonsai .git/index
	git submodule update --init --recursive
	cd bonsai; ruby compile.rb ../stable_marriage.bonsai > ../stable_marriage.c.tmp
	mv stable_marriage.c.tmp stable_marriage.c

clean:
	rm -f interpreter.c.tmp interpreter.c interpreter
