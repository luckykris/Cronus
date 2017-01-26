.PHONY: clean install
program=prometheus
complier="go"
config="/etc/prometheus.toml"
SCRIPTDIR='script'
build:init
	rm -rf /tmp/$(program) && mkdir /tmp/$(program) && $(complier) build -o /tmp/$(program)/$(program)  $(program).go
init:
	$(SCRIPTDIR)/clean_and_update_loacal_gosrc.sh
test_run:init
	$(complier) run $(program).go -config $(config) -debug
test_prometheus:init
	$(complier) test -v ./prometheus
install:
	cp /tmp/$(program)/$(program) /usr/bin/
clean:
	rm -rf /tmp/$(program)
	@echo "nothing to do"
