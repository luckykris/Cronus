.PHONY: clean install
program=prometheus
complier="go"
config="/etc/prometheus.toml"
SCRIPTDIR='script'

init:
	$(SCRIPTDIR)/clean_and_update_loacal_gosrc.sh
test:
	$(complier) run $(program).go -config $(config)
clean:
	@echo "nothing to do"
