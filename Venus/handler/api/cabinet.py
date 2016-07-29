#coding=utf-8
from handler.__init__ import *
from lib import prometheus
class cabinet(BaseHandler):
	@BaseHandler.apiProtocol
	def get(self):
		return self.prometheus.getCabinet()