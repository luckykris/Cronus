#!/bin/bash
HTTP_SERVER_API="http://127.0.0.1:81/v1"
CheckApi(){
	resp_code=`curl -X $2 -o /dev/null -s -d "$3" -w %{http_code} $HTTP_SERVER_API/$1`
	echo -n "$HTTP_SERVER_API/$1 --->$2 --data:\"$3\"-->"
	if [ "$2" == "GET" ] && [ "$resp_code" == "200" ];then
		echo -e "\033[32m Success — \033[0m"
		return
	elif [ "$2" == "POST" ] && [ "$resp_code" == "201" ];then
		echo -e "\033[32m Success  \033[0m"
                return

	elif [ "$2" == "PATCH" ] && [ "$resp_code" == "201" ];then
                echo -e "\033[32m Success  \033[0m"
                return
	elif [ "$2" == "DELETE" ] && [ "$resp_code" == "204" ];then
		echo -e "\033[32m Success  \033[0m"
                return

	else
		echo -e "\033[31m Failed  \033[0m"
	fi
	
}

CheckApi "device" "GET" 
CheckApi "device" "POST"  "deviceName=function-testing&deviceModelId=1"
CheckApi "device/1" "GET" 
CheckApi "device/1/tags" "GET" 
CheckApi "device/1/netPorts" "GET" 
CheckApi "cabinet" "GET" 
CheckApi "deviceModel" "GET" 
CheckApi "location" "GET" 
CheckApi "space" "GET" 
