#coding=utf-8
from handler.handlers  import *
route=[
#(r"/", index),
#(r"/space",space),
(r"/device",device.device),
(r"/space",space.space),
#################api##########
(r"/v1/space",api.space),
######################404##############
(r".*", BaseHandler)
]


