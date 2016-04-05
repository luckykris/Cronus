from flask import Flask,request,abort,render_template,make_response,redirect,url_for,session,escape
import setting
def init():
	for i in range(0,len(setting.data)):
		setting.data[i]["index"]=i
init()
app = Flask(__name__)

@app.route('/',methods=['GET', 'POST'])
def root():
	#abort(404)
	if request.method == 'GET':
		return redirect("/index")
@app.route('/index',methods=['GET', 'POST'])
def index():
	if request.method == 'GET':
		return render_template("index.html",doctors=setting.doctors,NAME=setting.NAME,items=setting.items,abouts=setting.abouts)
@app.route('/whoami',methods=['GET', 'POST'])
def whoami():
	if request.method == 'GET':
		return render_template("whoami.html",NAME=setting.NAME,ABOUT=setting.ABOUT)	
@app.route('/detail',methods=['GET', 'POST'])
def detail():
	if request.method == 'GET':
		i=int(request.args.get('i'))
		length=len(setting.data)
		if i <0:
			i=length-1
		elif i >=length:
			i=0
		return render_template("detail.html",NAME=setting.NAME,oneData=setting.data[i])				
@app.errorhandler(404)
def page_not_found(error):
	#return render_template('404.html'), 404
	resp = make_response(render_template('404.html'), 404)
	resp.headers['X-Something'] = 'XXXXX'
	return resp
    
if __name__ == '__main__':
	app.secret_key = 'A0Zr98j/3yX R~XHH!jmN]LWX/,?RT'
	app.run(debug=True,host='0.0.0.0',port=80)