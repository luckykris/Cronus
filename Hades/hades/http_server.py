#coding=utf8
from flask import Flask
from flask_restful import  Resource, Api
import lib.ansible_api


class api_get_property(Resource):
	def get(self):
		self.ansible_api()
		return {'task': 'Hello world'},200
	def post(self):
		return None,200
class new:
	def __init__(self,ansible_api,port=8888):
		self.port=port
		self.app = Flask(__name__)
		self.api=Api(self.app)
		self.api.add_resource(api_get_property,"/property")
	def __exit__(self):
		pass
	def run(self):
		self.app.run(port=self.port,host="0.0.0.0")