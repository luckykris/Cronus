#coding=utf-8
import tornado.web
import sys
import logging
import json
import traceback
import urllib
import time
import pymysql
from lib import prometheus
import setting
######logging conf path
formatter = logging.Formatter('[%(asctime)s][%(name)s][%(levelname)s]%(message)s') 


#  define a option  which can be load from command line,usage in program: options.XXXX
#try:
#        DB=db.ConnDb(setting.mysqldb)
#except Exception as e:
#        print("FATAL: Can not connect db :\n"+str(e))
#        sys.exit()




class BaseHandler(tornado.web.RequestHandler):
	def __init__(self,application, request, **kwargs):
		tornado.web.RequestHandler.__init__(self,application, request, **kwargs)
		self.prometheus=prometheus.Prometheus(setting.Prometheus['api'])
	def get_current_user(self):
		user=self.get_secure_cookie("user")
		if not user:
			return None
		try:
			self.PKey=json.loads(user.decode('utf-8'))
			return user
		except:
			self.write("您的cookie存在问题，请清除cookie后重新登录!")
			return None
	#@tornado.web.authenticated		
	def get(self):
		self.write_error(404)		
	def write_error(self, status_code, **kwargs):
		self.set_status(status_code)
		if status_code == 404:
			self.render('404.html')
	def getAllValue(self,ls):
		dc={}
		flg=0
		for i in ls:
			dc[i]=self.get_argument(i,None)
			if dc[i]=="":
				dc[i]=None
			if dc[i] is not None:
				flg+=1
		if flg==0:
		 	return None
		return dc
	# transport tuple list to dic list
	def tupleLs2dicLs(self,itemls,tls):
		return [dict(zip(itemls,t)) for t in tls]
	#trace back fatal error
	def Traceback(func):
		def wrap_f(self):
			try:
				func(self)
			except Exception as err :
				self.write({"success":False,"message":err})
		return wrap_f
	def apiProtocol(func):
		def wrap_f(self):
			success=False
			errDetail=None
			data=None
			try:
				data=func(self)
				success=True
			except prometheus.PrometheusError as err:
				errDetail=err
			except urllib.error.HTTPError as err:
				if err.code:
					errDetail="Not Found"
			except:
				errDetail=traceback.format_exc()
			finally:
				self.write({"success":success,"data":data,"message":errDetail})
		return wrap_f		