 #coding=utf-8
import requests
import json
class PrometheusError(Exception):
	def __init__(self, value):
		self.value = value
	def __str__(self):
		return repr(self.value)
class Prometheus:
	def __init__(self,url):
		self.url=url
	def __apiRequest(self,api,method,data={}):
		if method == "GET":
			r = requests.get(self.url+api,data)
		elif method == "POST":
			r = requests.post(self.url+api,data)
		elif method == "DELETE":
			r = requests.delete(self.url+api)
		elif method == "UPDATE":
			r = requests.patch(self.url+api,data)
		if r.status_code >399:
			raise PrometheusError("HTTP CODE:%d,Text:%s" % (r.status_code,r.json()))
		else:
			try:
				return r.json()
			except:
				return r.text
	def getDevice(self,deviceId=None):
		api="device"
		if deviceId !=None:
			api="device/%d" % deviceId
		return self.__apiRequest(api,'GET')
	def addDevice(self,device):
		return self.__apiRequest('device','POST',device)
	def updateDevice(self,device_id,device):
		return self.__apiRequest('device/%d' % device_id,'UPDATE',device)
	def deleteDevice(self,device_id):
		api='device/%d' % device_id
		return self.__apiRequest(api,'DELETE')
	def getCabinet(self,cabinetId=""):
		api="cabinet/%s" % str(cabinetId)
		return self.__apiRequest(api,'GET')	
	def getSpace(self,data={}):
		return self.__apiRequest('space','GET',data)
