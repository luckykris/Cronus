#coding: utf-8 
import os
import logging
class new:
	def __init__(self,cfg):
		self.level= eval("logging."+cfg.get('level').upper())
		self.formatter = logging.Formatter('[%(asctime)s][%(levelname)s]%(message)s') 
		self.logger = logging.getLogger(__name__)
		self.logger.setLevel(self.level)
		self.path=os.path.join("a/",)
		self.log_file_debug = logging.FileHandler(os.path.join(cfg.get('path'),"info.log"))
		self.log_file_error = logging.FileHandler(os.path.join(cfg.get('path'),"error.log"))
		self.log_file_debug.setLevel(logging.DEBUG)
		self.log_file_error.setLevel(logging.ERROR)
		self.log_file_debug.setFormatter(self.formatter)
		self.log_file_error.setFormatter(self.formatter)
		self.logger.addHandler(self.log_file_debug)
		self.logger.addHandler(self.log_file_error)
	def debug(self,m):
		self.logger.debug(m)
	def info(self,m):
		self.logger.info(m)
	def warn(self,m):
		self.logger.warn(m)
	def error(self,m):
		self.logger.error(m)
	def critical(self,m):
		self.logger.critical(m)