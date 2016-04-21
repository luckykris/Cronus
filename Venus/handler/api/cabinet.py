#coding=utf-8
from handler.__init__ import *
from lib import prometheus
class cabinet(BaseHandler):
	#@tornado.web.authenticated
	#@BaseHandler.Traceback
	def get(self):
		a=prometheus.Prometheus("http://192.168.33.81:81/v1/")
		self.write({"data":a.getCabinet({})})