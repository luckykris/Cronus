
EC_module_miss				=0x001
EC_http_request				=0x002
EC_json_analysis			=0x002


class AnsibleHttpError(Exception):
	def __init__(self,value):
		self.value=value
	def __str__(self):
		return repr("Error Code:%d" % self.value)