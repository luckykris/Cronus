#coding=utf-8
from handler.__init__ import *
#from lib import prometheus
class server(BaseHandler):
	#@tornado.web.authenticated
	#@BaseHandler.Traceback
	@BaseHandler.apiProtocol
	def get(self):
		return self.prometheus.getServer()
	@BaseHandler.apiProtocol
	def post(self):
		DeviceName=self.get_argument("DeviceName",None)
		DeviceModelId=self.get_argument("DeviceModelId",None)
		if DeviceName== None or DeviceModelId == None:
			raise Exception("You Need Fill Device Name")
		server={"DeviceName":DeviceName,"DeviceModelId":DeviceModelId}
		return self.prometheus.addServer(server)
	@BaseHandler.apiProtocol
	def delete(self):
		DeviceId=self.get_argument("DeviceId",None)
		if DeviceId== None:
			raise Exception("You Need Fill Device Name")
		return self.prometheus.deleteDevice({'DeviceId':DeviceId})
