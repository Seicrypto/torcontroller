all:checkPackage
	@echo "--------torcontroller--------"
checkPackage:
# curl
# For version-1.0 just use for check ip address.
	@dpkg -s curl >/dev/null 2>&1 || (\
	echo "The 'curl' package is not installed. Please install it using 'sudo apt install curl' and try again." \
	&& exit 1\
	)
# netcat
	@dpkg -s netcat-traditional >/dev/null 2>&1 || (\
	echo "The 'netcat-traditional' package is not installed. Please install it using 'sudo apt install netcat-traditional' and try again." \
	&& exit 1\
	)
# torsocket
# torsocket seems provide torify bash command. Not sure now.
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
# procps
# Necessary for systemctl.
	@dpkg -s procps >/dev/null 2>&1 || (\
	echo "The 'procps' package is not installed. Please install it using 'sudo apt install procps' and try again." \
	&& exit 1\
	)
# systemctl
	@dpkg -s systemctl >/dev/null 2>&1 || (\
	echo "The 'systemctl' package is not installed. Please install it using 'sudo apt install systemctl' and try again." \
	&& exit 1\
	)
install: all
	@echo "Preparing to install torcontroller package..."
# Place torcontroller scripts.
	install -D -m 755 torcontroller $(DESTDIR)/usr/bin/torcontroller
	install -D -m 555 ./lib/getIP.sh  $(DESTDIR)/usr/lib/torcontroller/getIP.sh
	install -D -m 555 ./lib/resetTorPassword.sh  $(DESTDIR)/usr/lib/torcontroller/resetTorPassword.sh
	install -D -m 555 ./lib/startTorcontrol.sh  $(DESTDIR)/usr/lib/torcontroller/startTorcontrol.sh
	install -D -m 555 ./lib/stopTorcontrol.sh  $(DESTDIR)/usr/lib/torcontroller/stopTorcontrol.sh
	install -D -m 555 ./lib/switchTorRouter.sh  $(DESTDIR)/usr/lib/torcontroller/switchTorRouter.sh
# Place setting files for tor, privoxy's config, and supervisor.
	install -D -m 644 ./etc/tor.service $(DESTDIR)/usr/src/torcontroller/tmp/tor.service
	install -D -m 644 ./etc/privoxy.service $(DESTDIR)/usr/src/torcontroller/tmp/privoxy.service
	install -D -m 644 ./etc/torrc $(DESTDIR)/usr/src/torcontroller/tmp/torrc
	install -D -m 644 ./etc/config $(DESTDIR)/usr/src/torcontroller/tmp/config
# Place docs
	install -D -m 644 README.md $(DESTDIR)/usr/share/doc/torcontroller/README.md
# Makefile install finished.
	@echo "torcontroller package Makefile worked successfully."
clean:
# Do not thing.
disclean: clean

uninstall:
	@echo "Uninstalling torcontroller..."
# Remove torcontroller scripts.
# Just remove each of torcontroller scripts,
# 'cause users might put thier own files there.
	rm -f $(DESTDIR)/usr/bin/torcontroller
	rm -f $(DESTDIR)/usr/lib/torcontroller/getIP.sh
	rm -f $(DESTDIR)/usr/lib/torcontroller/resetTorPassword.sh
	rm -f $(DESTDIR)/usr/lib/torcontroller/startTorcontrol.sh
	rm -f $(DESTDIR)/usr/lib/torcontroller/stopTorcontrol.sh
	rm -f $(DESTDIR)/usr/lib/torcontroller/switchTorRouter.sh
	rm -f $(DESTDIR)/usr/share/doc/torcontroller/README.md
# If it was empty after unistalled,
# directory would be remove by postrm script.