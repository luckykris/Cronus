.PHONY: clean install
program=hades



install: py_module_install  exe_file_intall  cfg_file_install module_file_install log_dir_mk

py_module_install:
	python setup.py install
exe_file_intall:
	cp bin/$(program) /usr/local/bin/ && chmod 755 /usr/local/bin/$(program)
cfg_file_install:
	if [ ! -d "/etc/$(program)" ];then mkdir /etc/$(program) &&cp $(program).ini /etc/$(program)/; fi 
module_file_install:
	if [ ! -d "/etc/$(program)/module" ];then mkdir /etc/$(program)/module ; fi;cp module/* /etc/$(program)/module/; 
log_dir_mk:
	if [ ! -d "/var/log/$(program)" ];then mkdir /var/log/$(program); fi 
clean:
	rm -rf build dist *.egg-info;find . -name *.pyc -exec rm -rvf {} \;
