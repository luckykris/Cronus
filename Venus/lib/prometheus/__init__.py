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
		try:
			if method == "GET":
				r = requests.get(self.url+api,data)
			elif method == "POST":
				r = requests.post(self.url+api,data)
			elif method == "DELETE":
				r = requests.delete(self.url+api)
			elif method == "UPDATE":
				r = requests.patch(self.url+api,data)
		except:
			raise PrometheusError("Can`t connect to prometheus.")
		if r.status_code >399:
			raise PrometheusError("HTTP CODE:%d,Text:%s" % (r.status_code,r.json()))
		else:
			try:
				return r.json()
			except:
				return r.text
	def getDevice(self,deviceId=None):
		api="devices"
		if deviceId !=None:
			api="devices/%d" % deviceId
		return self.__apiRequest(api,'GET')
	def addDevice(self,device):
		return self.__apiRequest('devices','POST',device)
	def updateDevice(self,device_id,device):
		return self.__apiRequest('devices/%d' % device_id,'UPDATE',device)
	def deleteDevice(self,device_id):
		api='devices/%d' % device_id
		return self.__apiRequest(api,'DELETE')
	def getServer(self,deviceId=None):
		api="server"
		if deviceId !=None:
			api="%s/%d" % (api,deviceId)
		return self.__apiRequest(api,'GET')
	def addServer(self,server):
		return  self.__apiRequest('server','POST',server)
	def getDeviceModel(self,deviceModelId=None):
		api="deviceModel"
		if deviceModelId !=None:
			api="%s/%d" % (api,deviceModelId)
		return self.__apiRequest(api,'GET')
	def getTag(self,tagId=None):
		api="tags"
		if tagId !=None:
		        api="tags/%d" % tagId
		return self.__apiRequest(api,'GET')
	def addTag(self,tag):
		return self.__apiRequest('tags','POST',tag)
	def updateTag(self,tag_id,tag):
		return self.__apiRequest('tags/%d' % tag_id,'UPDATE',tag)
	def deleteTag(self,tag_id):
		api='tags/%d' % tag_id
		return self.__apiRequest(api,'DELETE')
	def getCabinet(self,cabinetId=""):
		api="cabinet/%s" % str(cabinetId)
		return self.__apiRequest(api,'GET')
	def getCabinetSpace(self,cabinetId):
		api="cabinet/%d/space" % cabinetId
		return self.__apiRequest(api,'GET')	
	def getSpace(self,data={}):
		return self.__apiRequest('space','GET',data)
	def getDeviceNetPorts(self,device_id,netPort_id=None):
		api="devices/%d/netPorts/" % device_id
		if netPort_id is not None:
			api=api+str(netPort_id)
		return self.__apiRequest(api,'GET')	
	def addDeviceNetPorts(self,device_id,netPort):
		return self.__apiRequest('devices/%d/netPorts/' % device_id,'POST',netPort)
	def updateDeviceNetPorts(self,device_id,netPort_id,netPort):
		return self.__apiRequest('devices/%d/netPorts/%d' % (device_id,netPort_id),'UPDATE',netPort)
	def deleteDeviceNetPorts(self,device_id,netPort_id):
		api='devices/%d/netPorts/%d' % (device_id,netPort_id)
		return self.__apiRequest(api,'DELETE')
	def getDeviceTags(self,device_id,tag_id=None):
		api="devices/%d/tags/" % device_id
		if tag_id is not None:
			api=api+str(tag_id)
		return self.__apiRequest(api,'GET')	
	def addDeviceTags(self,device_id,tag_id):
		return self.__apiRequest('devices/%d/tags/%d' % (device_id,tag_id),'POST')
	def deleteDeviceTags(self,device_id,tag_id):
		return self.__apiRequest('devices/%d/tags/%d' % (device_id,tag_id),'DELETE')
		
