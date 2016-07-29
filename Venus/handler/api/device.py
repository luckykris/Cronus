#coding=utf-8
from handler.__init__ import *
#from lib import prometheus
class device(BaseHandler):
	#@tornado.web.authenticated
	#@BaseHandler.Traceback
	#@BaseHandler.apiProtocol
	#def get(self):
	#	return self.prometheus.getServer()
	#@BaseHandler.apiProtocol
	#def post(self):
	#	DeviceName=self.get_argument("DeviceName",None)
	#	if DeviceName== None:
	#		raise Exception("You Need Fill Device Name")
	#	server={"DeviceName":DeviceName}
	#	return self.prometheus.addServer(server)
	@BaseHandler.apiProtocol
	def get(self):
		return self.prometheus.getDevice()
	@BaseHandler.apiProtocol
	def delete(self):
		DeviceId=self.get_argument("DeviceId",None)
		if DeviceId== None:
			raise Exception("You Need Fill Device Name")
		return self.prometheus.deleteDevice(int(DeviceId))
