#coding=utf8
__requires__ = ['ansible']
import hades.error
try:
	from ansible.runner import Runner
except:
	raise hades.error.AnsibleHttpError(EC_miss_module)
class new:
	def __init__(self,cfg,log):
		self.timeout=cfg.get('timeout')
		self.remote_user=cfg.get('remote_user')
		self.remote_pass=cfg.get('remote_pass')
		self.forks=cfg.get('forks')
		self.log=log
	def __del__(self):
		pass
	def get_property(self,hosts):
		tmp_hosts={}
		for x in hosts:
			tmp_hosts[hosts[x]]=x
		runner = Runner(
			module_name="setup",
			#module_path=options.module_path,
			#module_args=options.module_args,
			remote_user=self.remote_user,
			remote_pass=self.remote_pass,
			#inventory=["192.168.33.81"],
			host_list=tmp_hosts.keys(),
			timeout=self.timeout,
			#private_key_file=options.private_key_file,
			forks=self.forks
			#pattern=pattern,
			#callbacks=self.callbacks,
			#transport=options.connection,
			#subset=options.subset,
			#check=options.check,
			#diff=options.check,
			#vault_pass=vault_pass,
			#become=options.become,
			#become_method=options.become_method,
			#become_pass=becomepass,
			#become_user=options.become_user,
			#extra_vars=extra_vars,
		)
		perfect_data={}
		meta_data=runner.run()
		for x in meta_data['dark']:
			perfect_data[tmp_hosts.get(x)]=None
		for x in meta_data['contacted']:
			try:
				metric=meta_data['contacted'][x]['ansible_facts']
				perfect_data[tmp_hosts.get(x)]={"hostname" 	:metric['ansible_hostname'],
											"serial"   	:metric['ansible_product_serial'],
											"os"       	:metric['ansible_distribution'],
											"release"  	:metric['ansible_distribution_version'],
											"processor"	:metric['ansible_processor_count'],
											"memsize"   :metric['ansible_memtotal_mb']
											}
			except Exception as err:
					self.log.critical("host:%s analysis failed. meta:>>%s<< ,failed:>>%s<<" %(str(x),str(meta_data['contacted'][x]),str(err)) )
		return perfect_data
