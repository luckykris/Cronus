 #coding=utf-8
import requests
import urllib
import json
class PrometheusError(Exception):
	def __init__(self, value):
		self.value = value
	def __str__(self):
		return repr(self.value)
class Prometheus:
	def __init__(self,url):
		self._deviceModel="deviceModel"
		self._server="server"
		self._location="location"
		self._idc="idc"
		self._cabinet="cabinet"
		self.url=url
	def __apiRequest(self,api,method,data={}):
		try:
			if method == "GET":
				r = requests.get(self.url+api,data)
			elif method == "POST":
				r = requests.post(self.url+api,data)
			elif method == "DELETE":
				suffix=""
				if data!={}:
					suffix="?"+urllib.parse.urlencode(data)
				r = requests.delete(self.url+api+suffix)
			elif method == "UPDATE":
				r = requests.patch(self.url+api,data)
		except Exception as e:
			raise PrometheusError("Can`t connect to prometheus.:"+str(e))
		if r.status_code >399:
			raise PrometheusError("HTTP CODE:%d,Text:%s" % (r.status_code,r.json()))
		else:
			try:
				return r.json()
			except:
				return r.text
	def AddDeviceModel(self,js):
		return self.__apiRequest(self._deviceModel,'POST',js)
	def GetDeviceModel(self,name):
		return self.__apiRequest(self._deviceModel+"/"+name,'GET')
	def DeleteDeviceModel(self,name):
		return self.__apiRequest(self._deviceModel+"/"+name,'DELETE')
	def AddServer(self,js):
		return self.__apiRequest(self._server,'POST',js)
	def GetServer(self,name):
		return self.__apiRequest(self._server+"/"+name,'GET')
	def DeleteServer(self,name):
		return self.__apiRequest(self._server+"/"+name,'DELETE')
	def UpdateServer(self,name,js):
		return self.__apiRequest(self._server+"/"+name,'UPDATE',js)
	def AddNetPort(self,dt,name,js):
		return self.__apiRequest(dt+"/"+name+"/netPorts",'POST',js)
	def DeleteNetPort(self,dt,name,js):
		return self.__apiRequest(dt+"/"+name+"/netPorts",'DELETE',js)
	def GetNetPort(self,dt,name):
		return self.__apiRequest(dt+"/"+name+"/netPorts",'GET')
	def AddLocation(self,js):
		return self.__apiRequest(self._location,'POST',js)
	def GetLocation(self,name):
		return self.__apiRequest(self._location+"/"+name,'GET')
	def DeleteLocation(self,name):
		return self.__apiRequest(self._location+"/"+name,'DELETE')
	def UpdateLocation(self,name,js):
		return self.__apiRequest(self._location+"/"+name,'UPDATE',js)
	def AddIdc(self,js):
		return self.__apiRequest(self._idc,'POST',js)
	def GetIdc(self,name):
		return self.__apiRequest(self._idc+"/"+name,'GET')
	def DeleteIdc(self,name):
		return self.__apiRequest(self._idc+"/"+name,'DELETE')
	def UpdateIdc(self,name,js):
		return self.__apiRequest(self._idc+"/"+name,'UPDATE',js)
	def AddCabinet(self,js):
		return self.__apiRequest(self._cabinet,'POST',js)
	def GetCabinet(self,name):
		return self.__apiRequest(self._cabinet+"/"+name,'GET')
	def DeleteCabinet(self,name):
		return self.__apiRequest(self._cabinet+"/"+name,'DELETE')
	def UpdateCabinet(self,name,js):
		return self.__apiRequest(self._cabinet+"/"+name,'UPDATE',js)


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
		
