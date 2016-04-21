#coding=utf-8
from handler.__init__ import *
class space(BaseHandler):
	#@tornado.web.authenticated
	#@BaseHandler.Traceback
	def get(self):
		self.render("space.html")
