#coding=utf-8
from handler.handlers  import *
route=[
#(r"/", index),
#(r"/space",space),
(r"/device",device.device),
(r"/space",space.space),
#################api##########
(r"/v1/space",api.space),
(r"/v1/cabinet",api.cabinet),
(r"/v1/server",api.server),
(r"/v1/device",api.device),
(r"/v1/deviceModel",api.deviceModel),
(r"/v1/netPort",api.netPort),
######################404##############
(r".*", BaseHandler)
]


