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
		DeviceName=self.get_argument("DeviceName",None)
		if DeviceName== None:
			raise Exception("You Need Fill Device Name")
		server={"DeviceName":int(DeviceName)}
		return self.prometheus.addServer(server)
