from  __init__ import *
a=Prometheus("http://127.0.0.1:81/v1/")
#a=Prometheus("http://172.16.3.20/v1/")
device={
	"DeviceName":"functiontest",
	"DeviceModelId":1
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
def device_update():
	devices_map=deviceMap()
	a.updateDevice(devices_map['functiontest'],device)
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
def cabinet_get():
	a.getCabinet()

@deco
def space_get():
	a.getSpace()

#print(a.getDevice())
#print("get device ok" )
#print(a.deleteDevice(2))
#print("delete device ok" )
device_add()
device_netPort_add()
device_netPort_update()
device_netPort_get()
device_update()
device_delete()
cabinet_get()
space_get()
