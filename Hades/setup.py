

from setuptools import setup, find_packages
import hades
setup(
      name="hades",
      version=hades.__version__,
      description="sub module of Cronus ",
      author="Kris Gu",
      url='https://github.com/luckykris',
      license="LGPL",
      #scripts=["scripts/test.py"],
      packages=[
      	"hades"
      ],
      install_requires=[
        'ansible<2.0',
        'requests>=2.11.1'
    ],
)
