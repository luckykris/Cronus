import ConfigParser
import os
import hades
BASE_DIR="/etc/%s" % hades.__program__
LOG_DIR="/var/log/%s" % hades.__program__
class new(object):
	def __init__(self, filename):
		self.config = ConfigParser.ConfigParser()
		self.config.read(filename)
		self.sections=self.config.sections()
		self.config_dict={}
		skv={'hades':
					 {"forks":{"default":10,"type":int},
					  "timeout":{"default":60,"type":int},
					  "remote_user":{"default":"root","type":basestring},
					  "remote_pass":{"default":None,"type":int},
					  "module_path":{"default":os.path.join(BASE_DIR,"module"),"type":basestring}
					 },
			 'logger':
			 		{"level":{"default":"info","type":basestring},
			 		 "path": {"default":LOG_DIR,"type":basestring}
			 		}
			}
		self.load_schema(skv)
		input_schema={
					 "plugin":{"default":None,"type":basestring},
					 "host":{"default":None,"type":basestring},
					 "port":{"default":80,"type":int},
					 "module_path":{"default":self.config_dict['hades']['module_path'],"type":basestring},
					 "extend":{"default":None,"type":basestring}
					}
		output_schema={"plugin":{"default":None,"type":basestring},
					 "host":{"default":None,"type":basestring},
					 "port":{"default":80,"type":int},
					 "module_path":{"default":self.config_dict['hades']['module_path'],"type":basestring},
					 "input":{"default":None,"type":basestring},
					 "extend":{"default":None,"type":basestring}
					}
		for x in ['hades','logger']:
			if x in self.sections:
				self.sections.remove(x)
		self.process_modules={}
		tmp_output=[]
		for s in self.sections:
			if not self.config.has_option(s,"type"):
				continue
			else:
				module_type=self.config.get(s,"type")
				if module_type == "input":
					self.load_schema({s:input_schema})
					self.process_modules[s]=[]
				elif module_type == "output":
					self.load_schema({s:output_schema})
					tmp_output.append(s)
		for s in tmp_output:
			tmp_input=self.config_dict[s]['input']
			if tmp_input in self.process_modules:
				self.process_modules[tmp_input].append(s)
	def load_schema(self,schema):
		for s in schema:
			self.config_dict[s]={}
			for k in schema[s]:
				if self.config.has_option(s,k):
					if schema[s][k]["type"] == basestring:
						tmp_item=self.config.get(s,k)
					elif schema[s][k]["type"] == int:
						tmp_item=self.config.getint(s,k)
				else:
					tmp_item=schema[s][k]["default"]
				self.config_dict[s][k]=tmp_item
	def get(self,s,k=None):
		if k==None:
			return self.config_dict[s]
		return self.config_dict[s][k]
	def get_process_modules(self):
		return self.process_modules