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