#coding=utf-8
import requests
from net import ipv4
def start(cfg):
	session = requests.Session()
	r = session.get("http://%s:%d/api/auth" % (cfg.get('host'),cfg.get("port")), data={"u":"cris.gu","p":"redhat","m":"on"} )
	if r.status_code > 399:
		raise Exception("身份验证失败:%s" % r.text)
	r = session.get("http://%s:%d/api/device" % (cfg.get('host'),cfg.get("port")),data={"device_types":cfg.get("extend"),"method":"get"} )
	if r.status_code > 399:
		raise Exception("执行失败!错误内容:%s" % r.text)
	else:
		tmp_r=r.json()
		if tmp_r['success'] != True:
			raise Exception("执行失败!错误内容:%s" % tmp_r['message'])
		else:
			server_dict={}
			for x in tmp_r['data']:
				if len(x['netPorts'])>0:
					server_dict[x['device_id']]=ipv4.Ipv4int2Ipv4string(x['netPorts'][0]['ipv4_int'])
			return server_dict
