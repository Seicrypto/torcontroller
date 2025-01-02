all:checkPackage
	@echo "--------torcontroller--------"

checkPackage:
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
	@dpkg -s procps >/dev/null 2>&1 || (\
	echo "The 'procps' package is not installed. Please install it using 'sudo apt install procps' and try again." \
	&& exit 1\
	)
# systemctl
	@dpkg -s systemctl >/dev/null 2>&1 || (\
	echo "The 'systemctl' package is not installed. Please install it using 'sudo apt install systemctl' and try again." \
	&& exit 1\
	)
# iptables
	@dpkg -s iptables >/dev/null 2>&1 || (\
	echo "The 'iptables' package is not installed. Please install it using 'sudo apt install iptables' and try again." \
	&& exit 1\
	)
# sudo
	@dpkg -s sudo >/dev/null 2>&1 || (\
	echo "The 'sudo' package is not installed. Please install it using 'sudo apt install sudo' and try again." \
	&& exit 1\
	)

install: all
	@echo "Preparing to install torcontroller package..."
# Install the torcontroller binary
	install -D -m 755 torcontroller $(DESTDIR)/usr/bin/torcontroller
# Install the setting files
	install -D -m 644 initializer/templates/torcontroller.yml $(DESTDIR)/usr/share/torcontroller/defaults/torcontroller.yml
	install -D -m 644 initializer/templates/tor.service $(DESTDIR)/usr/share/torcontroller/defaults/tor.service
	install -D -m 644 initializer/templates/privoxy.service $(DESTDIR)/usr/share/torcontroller/defaults/privoxy.service
	install -D -m 644 initializer/templates/tor/torrc $(DESTDIR)/usr/share/torcontroller/defaults/tor/torrc
	install -D -m 644 initializer/templates/privoxy/config $(DESTDIR)/usr/share/torcontroller/defaults/privoxy/config
	install -D -m 644 initializer/templates/sudoers.d/torcontroller $(DESTDIR)/usr/share/torcontroller/defaults/sudoers.d/torcontroller

	@echo "torcontroller package installation completed."

clean:
# Do not thing.
disclean: clean

uninstall:
	@echo "Uninstalling torcontroller..."
# Remove the torcontroller binary
	rm -f $(DESTDIR)/usr/bin/torcontroller
# Remove defualt setting files
	rm -f $(DESTDIR)/usr/share/torcontroller/defaults/torcontroller.yml
	rm -f $(DESTDIR)/usr/share/torcontroller/defaults/tor.service
	rm -f $(DESTDIR)/usr/share/torcontroller/defaults/privoxy.service
	rm -f $(DESTDIR)/usr/share/torcontroller/defaults/tor/torrc
	rm -f $(DESTDIR)/usr/share/torcontroller/defaults/privoxy/config
	rm -f $(DESTDIR)/usr/share/torcontroller/defaults/sudoers.d/torcontroller

	@echo "torcontroller package uninstallation completed."