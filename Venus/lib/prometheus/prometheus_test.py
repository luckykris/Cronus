from  __init__ import *
a=Prometheus("http://192.168.33.81:81/v1/")
print(a.getSpace({}))