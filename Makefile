all:checkPackage
	@echo "--------torcontroller--------"
checkPackage:
    @dpkg -s tor >/dev/null 2>&1 || (\
	echo "The 'tor' package is not installed. Please install it using 'sudo apt install tor' and try again." \
	&& exit 1\
	)
    @dpkg -s privoxy >/dev/null 2>&1 || (\
	echo "The 'privoxy' package is not installed. Please install it using 'sudo apt install privoxy' and try again." \
	&& exit 1\
	)
install: all
# Install 
	install -D -m 755 ./lib/installFiles/torrc $(DESTDIR)/etc/tor/torrc
	install -D -m 755 ./lib/installFiles/config $(DESTDIR)/etc/privoxy/config
clean:
# Do not thing.
disclean: clean

uninstall:
	rm -f $(DESTDIR)/bin/debtest/helloworld.sh