from hades import error
import sys
class new:
	def __init__(self,output_cfg,logger):
		self.output_cfg=output_cfg
		self.plugin=self.output_cfg.get('plugin')
		sys.path.append(output_cfg.get('module_path'))
		try:
			self.plugin_obj=__import__(self.plugin)
		except:
			self.log.error("can`t find the module plugin '%s'." % (self.plugin))
			raise error.AnsibleHttpError(error.EC_module_miss)
	def run_once(self,data):
		return self.plugin_obj.start(self.output_cfg,data)
	def start(self):
		pass
