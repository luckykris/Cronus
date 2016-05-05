#coding=utf-8
import urllib.request
import urllib.parse
import json
class Prometheus:
	def __init__(self,url):
		self.url=url
	def __apiRequest(self,api,method,data={}):
		req_data=urllib.parse.urlencode(data)
		req = urllib.request.Request(self.url+api,req_data.encode('ascii'),method=method)
		header={"Content-Type": "application/json"}
		for k in header:
			req.add_header(k, header[k])
		result = urllib.request.urlopen(req)
		return result.read()
		#resp=json.loads(result.read().decode('utf-8'))	
		#return resp
	def getDevice(self,deviceId=None):
		api="device"
		if deviceId !=None:
			api="device/%d" % deviceId
		return self.__apiRequest(api,'GET')
	def addDevice(self,device):
		return self.__apiRequest('device','POST',device)
	def deleteDevice(self,device_id):
		api='device/%d' % device_id
		return self.__apiRequest(api,'DELETE')
	def getCabinet(self,cabinetId=""):
		api="cabinet/%s" % str(cabinetId)
		print(api)
		return self.__apiRequest(api,'GET')	
	def getSpace(self,data={}):
		return self.__apiRequest('space','GET',data)
