from  __init__ import *
a=Prometheus("http://127.0.0.1:81/v1/")
#a=Prometheus("http://172.16.3.20/v1/")
name_suffix="test"

DeviceModel={
	"DeviceModelName":name_suffix+"-server-device-model",
	"DeviceType":"Server",
	"U":2,
	"HALF_FULL":"half"}
Server={
	"DeviceName":name_suffix+"-server",
	"DeviceModelName":DeviceModel["DeviceModelName"],
	"GroupId":10,
	"Env":3,
	"Serial":"asdasd",
	"Hostname":"myhost",
	"Memsize":9999,
	"Os":"RedHat",
	"Release":6.5,
	"Ipv4":"200.200.200.1"
}
tag={
	"TagName":"functiontest_tag",
}
Location={
	"LocationName":name_suffix+"-location"
}
Idc={
	"IdcName":name_suffix+"-idc",
	"LocationName":Location["LocationName"]
}
Cabinet={
	"CabinetName":name_suffix+"-cabinet",
	"CapacityTotal":200,
	"IdcName":Idc["IdcName"]
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
def DeviceModel_add():
	a.AddDeviceModel(DeviceModel)
@deco
def DeviceModel_get():
	a.GetDeviceModel(DeviceModel["DeviceModelName"])
@deco
def DeviceModel_delete():
	a.DeleteDeviceModel(DeviceModel["DeviceModelName"])


@deco
def Server_add():
	a.AddServer(Server)
@deco
def Server_get():
	a.GetServer(Server["DeviceName"])
@deco
def Server_delete():
	a.DeleteServer(Server["DeviceName"])
@deco
def Server_update():
	a.UpdateServer(Server["DeviceName"],Server)
@deco
def NetPort_add():
	a.AddNetPort("server",Server["DeviceName"],Server)
@deco
def NetPort_get():
	a.GetNetPort("server",Server["DeviceName"])
@deco
def NetPort_delete():
	a.DeleteNetPort("server",Server["DeviceName"],Server)


@deco
def Location_add():
	a.AddLocation(Location)
@deco
def Location_get():
	a.GetLocation(Location["LocationName"])
@deco
def Location_delete():
	a.DeleteLocation(Location["LocationName"])
@deco
def Location_update():
	a.UpdateLocation(Location["LocationName"],Location)



@deco
def Idc_add():
	a.AddIdc(Idc)
@deco
def Idc_get():
	a.GetIdc(Idc["IdcName"])
@deco
def Idc_delete():
	a.DeleteIdc(Idc["IdcName"])
@deco
def Idc_update():
	a.UpdateIdc(Idc["IdcName"],Idc)


@deco
def Cabinet_add():
	a.AddCabinet(Cabinet)
@deco
def Cabinet_get():
	a.GetCabinet(Cabinet["CabinetName"])
@deco
def Cabinet_delete():
	a.DeleteCabinet(Cabinet["CabinetName"])
@deco
def Cabinet_update():
	a.UpdateCabinet(Cabinet["CabinetName"],Cabinet)

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


Location_add()
Location_get()
Location_update()
Idc_add()
Idc_get()
Idc_update()
Cabinet_add()
Cabinet_get()
Cabinet_update()
DeviceModel_add()
DeviceModel_get()
Server_add()
Server_get()
Server_update()
NetPort_add()
NetPort_get()
NetPort_delete()
Server_delete()
DeviceModel_delete()
Cabinet_delete()
Idc_delete()
Location_delete()
