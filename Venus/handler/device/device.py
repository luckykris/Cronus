#coding=utf-8
from handler.__init__ import *
class device(BaseHandler):
	#@tornado.web.authenticated
	#@BaseHandler.Traceback
	def get(self):
		self.render("device.html")
