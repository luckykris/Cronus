 #coding=utf-8
import urllib.request
import json
def ApiRequest(url,data):
	req_data=json.dumps(data)
	req = urllib.request.Request(url, req_data.encode('utf-8'))
	header={"Content-Type": "application/json"}
	for k in header:
		req.add_header(k, header[k])
	result = urllib.request.urlopen(req)
	resp=json.loads(result.read().decode('utf-8'))	
	return resp