#coding=utf-8
import tornado.httpserver
import tornado.ioloop
import logging
#      load argument from command line
import tornado.options
from tornado.options import define ,options
import os
import sys
import time
import route
import version
#  define a option  which can be load from command line,usage in program: options.XXXX
define("port", default=80, help="run on the given port", type=int)
define("host", default="0.0.0.0", help="setting TCP listening address ", type=str)
define("processes", default=10, help="how manay processes to run on the multi processes module", type=int)

LOG = logging.getLogger(__name__)
def VenusWebServer():
    tornado.options.options.logging = "none"
#   start up options module to load argument from command line
    tornado.options.parse_command_line()
    app = tornado.web.Application(handlers=route.route,
			      template_path=os.path.join(os.path.dirname(__file__), "templates"),	
				  static_path=os.path.join(os.path.dirname(__file__), "static"),
#				  ui_modules={'login':LoginModule},
				  cookie_secret="124oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o/Vo=",
				  login_url="/login",
                  debug=True
    )
    http_server = tornado.httpserver.HTTPServer(app,no_keep_alive=True)
    #http_server.bind(options.port,'0.0.0.0')
    http_server.listen(options.port)
    #http_server.start(num_processes=50)
    tornado.ioloop.IOLoop.instance().start()
if __name__ == "__main__":
    VenusWebServer()