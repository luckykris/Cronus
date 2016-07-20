#!/usr/bin/python
import ansible.runner
import sys
import json
hosts=sys.argv[1]
hosts_ls=hosts.split(',')

runner = ansible.runner.Runner(
   module_name='setup',
   #module_name='ping',
   module_args='',
   pattern='all',
   forks=1,
   host_list=hosts_ls
)
datastructure = runner.run()
print(json.dumps(datastructure))
