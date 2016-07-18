from  __init__ import *
a=Prometheus("http://127.0.0.1:81/v1/")
#a=Prometheus("http://172.16.3.20/v1/")
device={
	"DeviceName":"functiontest",
	"DeviceModelId":1
}
tag={
	"TagName":"functiontest_tag",
}
netPort={
	"Mac":"FF:FF:FF:FF:FF:FF",
	"Ipv4Int":1,
	"Type":"eth"
}

def deco(func):
	def _deco():
		try:
			func()
			print("%s---> \033[32m Success \033[0m  " % (func.__name__))
		except Exception as err:
			print("%s---> \033[31m Fail \033[0m  ERROR:%s" % (func.__name__,str(err)))
	return _deco

def deviceMap():
	devices=a.getDevice()
	devices_map={x['DeviceName']:x['DeviceId'] for x  in devices}
	return devices_map
def tagMap():
	tags=a.getTag()
	tags_map={x['TagName']:x['TagId'] for x  in tags}
	return tags_map
def device_netPortMap(device_id):
	netPorts=a.getDeviceNetPorts(device_id)
	netPorts_map={x['Ipv4Int']:x['NetPortId'] for x  in netPorts}
	return netPorts_map
@deco	
def device_add():
	a.addDevice(device)
@deco
def device_get():
	devices=a.getDevice()
@deco
def device_get_one():
	devices_map=deviceMap()
	a.getDevice(devices_map['functiontest'])
@deco
def tag_add():
	a.addTag(tag)
@deco
def tag_get():
	tags=a.getTag()
@deco
def tag_get_one():
	tags_map=tagMap()
	a.getTag(tags_map['functiontest_tag'])
@deco
def tag_update():
	tags_map=tagMap()
	a.updateTag(tags_map['functiontest_tag'],tag)
@deco	
def device_update():
	devices_map=deviceMap()
	a.updateDevice(devices_map['functiontest'],device)
@deco	
def device_tag_add():
	devices_map=deviceMap()
	tags_map=tagMap()
	a.addDeviceTags(devices_map['functiontest'],tags_map['functiontest_tag'])
@deco	
def device_tag_get():
	devices_map=deviceMap()
	tags_map=tagMap()
	a.getDeviceTags(devices_map['functiontest'],tags_map['functiontest_tag'])
@deco	
def device_tag_delete():
	devices_map=deviceMap()
	tags_map=tagMap()
	a.deleteDeviceTags(devices_map['functiontest'],tags_map['functiontest_tag'])
@deco	
def device_netPort_add():
	devices_map=deviceMap()
	a.addDeviceNetPorts(devices_map['functiontest'],netPort)
@deco	
def device_netPort_update():
	devices_map=deviceMap()
	netPorts_map=device_netPortMap(devices_map['functiontest'])
	a.updateDeviceNetPorts(devices_map['functiontest'],netPorts_map[1],netPort)
@deco	
def device_netPort_delete():
	devices_map=deviceMap()
	netPorts_map=device_netPortMap(devices_map['functiontest'])
	a.deleteDeviceNetPorts(devices_map['functiontest'],netPorts_map[1])
@deco	
def device_netPort_get():
	devices_map=deviceMap()
	a.getDeviceNetPorts(devices_map['functiontest'])
@deco
def device_delete():
	devices_map=deviceMap()
	a.deleteDevice(devices_map['functiontest'])
@deco
def tag_delete():
	tags_map=tagMap()
	a.deleteTag(tags_map['functiontest_tag'])
@deco
def cabinet_get():
	a.getCabinet()

@deco
def space_get():
	a.getSpace()

device_add()
device_get()
device_get_one()
tag_add()
tag_get()
tag_get_one()
tag_update()
device_netPort_add()
device_netPort_update()
device_netPort_get()
device_tag_add()
device_tag_get()
device_tag_delete()
device_update()
device_delete()
tag_delete()
cabinet_get()
space_get()
