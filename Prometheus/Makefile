.PHONY: clean install
program=prometheus
complier="go"
config="/etc/prometheus.toml"
SCRIPTDIR='script'

init:
	$(SCRIPTDIR)/clean_and_update_loacal_gosrc.sh
test:init
	$(complier) run $(program).go -config $(config) -debug
test_prometheus:init
	$(complier) test -v ./prometheus
clean:
	@echo "nothing to do"
