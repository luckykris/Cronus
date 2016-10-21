#coding=utf-8
import requests
from net import ipv4
def start(cfg,data):
	session = requests.Session()
	r = session.get("http://%s:%d/api/auth" % (cfg.get('host'),cfg.get("port")), data={"u":"cris.gu","p":"redhat","m":"on"} )
	if r.status_code > 399:
		raise Exception("身份验证失败:%s" % r.text)
	for x in data:
		data[x]['device_type']=cfg.get("extend")
		data[x]['device_id']=x
		r = session.get("http://%s:%d/api/llsd" % (cfg.get('host'),cfg.get("port")),data=data[x] )
		if r.status_code > 399:
			raise Exception("执行失败!错误内容:%s" % r.text)
		else:
			tmp_r=r.json()
			if tmp_r['success'] != True:
				raise Exception("执行失败!错误内容:%s" % tmp_r['message'])
			else:
				pass
