#coding=utf-8
from handler.__init__ import *
#from lib import prometheus
class deviceModel(BaseHandler):
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
		return self.prometheus.getDeviceModel()
