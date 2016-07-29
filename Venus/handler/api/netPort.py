#coding=utf-8
from handler.__init__ import *
#from lib import prometheus
class netPort(BaseHandler):
	#@tornado.web.authenticated
	#@BaseHandler.Traceback
	@BaseHandler.apiProtocol
	def get(self):
		DeviceId=self.get_argument("DeviceId",None)
		if DeviceId== None:
			raise Exception("You Need Fill Device Name")
		return self.prometheus.getDeviceNetPorts(int(DeviceId))
	@BaseHandler.apiProtocol
	def post(self):
		DeviceId=self.get_argument("DeviceId",None)
		Ipv4Int=self.get_argument("Ipv4Int",None)
		if DeviceId== None:
			raise Exception("You Need Fill Device Name")
		netPort={"Ipv4Int":Ipv4Int,"Type":"Unknow"}
		return self.prometheus.addDeviceNetPorts(int(DeviceId),netPort)
	@BaseHandler.apiProtocol
	def delete(self):
		DeviceId=self.get_argument("DeviceId",None)
		NetPortId=self.get_argument("NetPortId",None)
		if DeviceId== None:
			raise Exception("You Need Fill Device Name")
		return self.prometheus.deleteDeviceNetPorts(int(DeviceId),int(NetPortId))