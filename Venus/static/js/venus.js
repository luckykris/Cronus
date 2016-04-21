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
	ajax("get",false,APIVERSION+"/"+item,data,function(json){a=json},function(){alert(APIERRINFO)});
	return a;
}
