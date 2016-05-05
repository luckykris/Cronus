from  __init__ import *
a=Prometheus("http://127.0.0.1:81/v1/")
#a=Prometheus("http://172.16.3.20/v1/")
device={
	"deviceName":"functiontest",
	"deviceModelId":1
}
#print(a.getDevice())
#print("get device ok" )
#print(a.deleteDevice(2))
#print("delete device ok" )
print(a.addDevice(device))
print("set device ok" )
