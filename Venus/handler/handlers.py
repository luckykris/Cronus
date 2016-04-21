#coding=utf-8
from handler.__init__ import *
from handler import api
from handler import space
from handler import device

class index(BaseHandler):
	#@tornado.web.authenticated
	#@BaseHandler.Traceback
	def get(self):
		self.render("index.html")