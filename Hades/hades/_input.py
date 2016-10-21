from hades import error
import sys
class new:
	def __init__(self,input_cfg,logger):
		self.input_cfg=input_cfg
		self.log=logger
		self.plugin=self.input_cfg.get('plugin')
		sys.path.append(input_cfg.get('module_path'))
		try:
			self.plugin_obj=__import__(self.plugin)
		except:
			self.log.error("can`t find the module plugin '%s'." % (self.plugin))
			raise error.AnsibleHttpError(error.EC_module_miss) 
	def run_once(self):
		return self.plugin_obj.start(self.input_cfg)
	def start(self):
		pass
