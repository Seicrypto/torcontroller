all:checkPackage
	@echo "--------torcontroller--------"
checkPackage:
# curl
	@dpkg -s curl >/dev/null 2>&1 || (\
	echo "The 'netcat' package is not installed. Please install it using 'sudo apt install netcat' and try again." \
	&& exit 1\
	)
# netcat
	@dpkg -s netcat >/dev/null 2>&1 || (\
	echo "The 'netcat' package is not installed. Please install it using 'sudo apt install netcat' and try again." \
	&& exit 1\
	)
# tor
	@dpkg -s tor >/dev/null 2>&1 || (\
	echo "The 'tor' package is not installed. Please install it using 'sudo apt install tor' and try again." \
	&& exit 1\
	)
# privoxy
	@dpkg -s privoxy >/dev/null 2>&1 || (\
	echo "The 'privoxy' package is not installed. Please install it using 'sudo apt install privoxy' and try again." \
	&& exit 1\
	)
install: all
	@echo "Preparing to install torcontroller package..."
# Make sure tor and privoxy service already stopped.
	@service privoxy stop
	-@kill $(pidof tor) || echo "tor already stopped"
	@service tor stop
# Place tor and privoxy setting files.
# Backup user tor and privoxy setting before install.
	@if [ -f /etc/tor/torrc ]; then cp /etc/tor/torrc /etc/tor/torrc.back; fi
	install -D -m 755 ./installFiles/torrc $(DESTDIR)/etc/tor/torrc
	@if [ -f /etc/privoxy/config ]; then cp /etc/privoxy/config /etc/privoxy/config.back; fi
	install -D -m 755 ./installFiles/config $(DESTDIR)/etc/privoxy/config
# Place torcontroller scripts.
	install -D -m 755 torcontroller $(DESTDIR)/usr/bin/torcontroller
	install -D -m 555 ./lib/getIP.sh  $(DESTDIR)/usr/lib/torcontroller/getIP.sh
	install -D -m 555 ./lib/resetTorPassword.sh  $(DESTDIR)/usr/lib/torcontroller/resetTorPassword.sh
	install -D -m 555 ./lib/startTorcontrol.sh  $(DESTDIR)/usr/lib/torcontroller/startTorcontrol.sh
	install -D -m 555 ./lib/stopTorcontrol.sh  $(DESTDIR)/usr/lib/torcontroller/stopTorcontrol.sh
	install -D -m 555 ./lib/switchTorRouter.sh  $(DESTDIR)/usr/lib/torcontroller/switchTorRouter.sh
# set 'torcontroller' as the default tor authentication password.
	@hashTorPWD=$(tor --hash-password "torcontroller" | tail -n 1)
	@sed -i "/HashedControlPassword/s/.*/HashedControlPassword $hashTorPWD/" /etc/tor/torrc
	@unset hashTorPWD
# Makefile install finished.
	@echo "torcontroller package installed successfully."
clean:
# Do not thing.
disclean: clean

uninstall:
	@echo "Uninstalling torcontroller..."
# Make sure tor and privoxy service already stopped.
	@service privoxy stop
	@kill $(pidof tor)
	@service tor stop
# Remove tor and privoxy setting files.
# If tor backup existed, would replace it.
	rm -f $(DESTDIR)/etc/tor/torrc
	@if [ -f $(DESTDIR)/etc/tor/torrc.back ]; then \
	cp $(DESTDIR)/etc/tor/torrc.back $(DESTDIR)/etc/tor/torrc; \
	rm -f $(DESTDIR)/etc/tor/torrc.back; \
	fi
# If privoxy backup existed, would replace it.
	rm -f $(DESTDIR)/etc/privoxy/config
	@if [ -f $(DESTDIR)/etc/privoxy/config.back ]; then \
	cp $(DESTDIR)/etc/privoxy/config.back $(DESTDIR)/etc/privoxy/config; \
	rm -f $(DESTDIR)/etc/privoxy/config.back; \
	fi
# Remove torcontroller scripts.
	rm -f $(DESTDIR)/usr/bin/torcontroller
	rm -f $(DESTDIR)/usr/lib/torcontroller/getIP.sh
	rm -f $(DESTDIR)/usr/lib/torcontroller/resetTorPassword.sh
	rm -f $(DESTDIR)/usr/lib/torcontroller/startTorcontrol.sh
	rm -f $(DESTDIR)/usr/lib/torcontroller/stopTorcontrol.sh
	rm -f $(DESTDIR)/usr/lib/torcontroller/switchTorRouter.sh