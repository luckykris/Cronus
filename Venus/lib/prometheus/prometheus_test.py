from  __init__ import *
a=Prometheus("http://127.0.0.1:81/v1/")
#a=Prometheus("http://172.16.3.20/v1/")
device={
	"DeviceName":"functiontest",
	"DeviceModelId":1
}

def deco(func):
	def _deco():
		try:
			func()
			print("%s---> \033[32m Success \033[0m  " % (func.__name__))
		except Exception as err:
			print("%s---> \033[31m Fail \033[0m  ERROR:%s" % (func.__name__,str(err)))
	return _deco
@deco	
def device_add():
	a.addDevice(device)
@deco
def device_get():
	devices=a.getDevice()
@deco	
def device_update():
	devices=a.getDevice()
	devices_map={x['DeviceName']:x['DeviceId'] for x  in devices}
	a.updateDevice(devices_map['functiontest'],device)
@deco
def device_delete():
	devices=a.getDevice()
	devices_map={x['DeviceName']:x['DeviceId'] for x  in devices}
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
device_update()
device_delete()
cabinet_get()
space_get()
