#coding=utf-8
from handler.__init__ import *
from lib import prometheus
class cabinet(BaseHandler):
	#@tornado.web.authenticated
	#@BaseHandler.Traceback
	def get(self):
		self.write({"data":self.prometheus.getCabinet({})})