maninstall: man/goserve.1
	cp man/goserve.1 /usr/share/man/man1/goserve.1
	gzip /usr/share/man/man1/goserve.1

manuninstall:
	rm /usr/share/man/man1/goserve.1.gzip
