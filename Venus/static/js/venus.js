APIVERSION="v1"
APIERRINFO="Some Error Duaring Interactive With Prometheus."
function ajax(method,async,url,data,callback,errcallback){
	$.ajax({
              type:method,
              async:async,
              url:url,
              dataType:"json",
              data:data,
              success:callback,
              error:errcallback
      });	
}

function  ApiGet(item,data){
	var a =null;
	ajax("get",false,"/"+APIVERSION+"/"+item,data,function(json){a=json},function(){alert(APIERRINFO)});
	return a;
}
function  ApiPost(item,data){
  var a =null;
  ajax("post",false,"/"+APIVERSION+"/"+item,data,function(json){a=json},function(){alert(APIERRINFO)});
  return a;
}
function  ApiDelete(item,data){
  var a =null;
  ajax("delete",false,"/"+APIVERSION+"/"+item,data,function(json){a=json},function(){alert(APIERRINFO)});
  return a;
}

function GetColumn(dic){
  column=[]
  $.each(dic,function(k,v){column.push(k)})
  return column
}


function Timestamp2string(timestamp){
  var newDate = new Date();
  newDate.setTime(timestamp * 1000);
  return newDate.toLocaleString();
}

function Ip4vInt2string(ipv4Int){
  d_c_b_a=[]
  IPV4MASK = 1 << 8
  for(s= 0;s < 4; s++ ){
    d_c_b_a[s] = (ipv4Int % IPV4MASK).toString()
    ipv4Int = ipv4Int >>>8
  }
  return d_c_b_a.reverse().join(".")
}

function String2ip4vInt(ipv4_str){
  a_b_c_d=ipv4_str.split(".")
  ipv4Int=0
  for(s= 0;s < 4; s++ ){
    ipv4Int=ipv4Int*256
    ipv4Int+=parseInt(a_b_c_d[s])
  }
  return ipv4Int
}


//
function DeleteDevice(id){
  a=ApiDelete("device",{"DeviceId":id})
  if(a.success){
    alert("deleted!")
  }else{
    alert(a.message)
  }
}
function DeleteNetPort(device_id,netPort_id){
  a=ApiDelete("netPort",{"DeviceId":device_id,"NetPortId":netPort_id})
  if(a.success){
    alert("deleted!")
  }else{
    alert(a.message)
  }
}

function GetCabinet(){
  a=ApiGet("cabinet",{})
  if(a.success){
    return a.data
  }else{
    alert(a.message)
  }
}

function GetDeviceMap(){
  var map={}
  a=ApiGet("device",{})
  if(!a.success){
    alert(a.message)
  }else{
    $.each(a.data,function(k,v){map[v.DeviceId]=v.DeviceName})
  }
  return map
}

function GetDeviceModel(){
  a=ApiGet("deviceModel",{})
  if(a.success){
    return a.data
  }else{
    alert(a.message)
  }
}

function GetDeviceModelMap(){
  var map={}
  a=ApiGet("deviceModel",{})
  if(!a.success){
    alert(a.message)
  }else{
    $.each(a.data,function(k,v){map[v.DeviceModelId]=v.DeviceModelName})
  }
  return map
}


function GetCabinetSpace(cabinet_id){
  a=ApiGet("space",{"cabinetId":cabinet_id})
  if(a.success){
    return a.data
  }else{
    alert(a.message)
  }
}