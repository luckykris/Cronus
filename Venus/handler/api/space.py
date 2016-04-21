#coding=utf-8
from handler.__init__ import *
class space(BaseHandler):
	#@tornado.web.authenticated
	@BaseHandler.apiProtocol
	def get(self):
		cabinet_id=self.get_argument("cabinetId","")
		cabinet=self.prometheus.getCabinet(cabinet_id)
		total=cabinet['CapacityTotal']
		spaces=self.prometheus.getSpace({"cabinet_id":cabinet_id})
		dc=dict([(x['UPosition'],x['DeviceId']) for x in spaces])
		sd=dict([(i,None) for i in range(1,total+1)])
		for space in spaces:
			sd[space['UPosition']]=space['DeviceId']
		return [{"UPosition":i,"DeviceId":sd[i]} for i in range(1,total+1)[::-1]]