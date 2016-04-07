from flask import Flask,request,abort,render_template,make_response,redirect,url_for,session,escape

app = Flask(__name__)

@app.route('/',methods=['GET', 'POST'])
def root():
	#abort(404)
	if request.method == 'GET':
		return redirect("/home")
@app.route('/home',methods=['GET', 'POST'])
def home():
	if request.method == 'GET':
		return render_template("home.html")
@app.route('/device',methods=['GET', 'POST'])
def device():
	if request.method == 'GET':
		return render_template("device.html")
@app.route('/space',methods=['GET', 'POST'])
def space():
	if request.method == 'GET':
		return render_template("space.html")		
@app.errorhandler(404)
def page_not_found(error):
	#return render_template('404.html'), 404
	resp = make_response(render_template('404.html'), 404)
	resp.headers['X-Something'] = 'XXXXX'
	return resp
    
if __name__ == '__main__':
	app.secret_key = 'A0Zr98j/3yX R~XHH!jmN]LWX/,?RT'
	app.run(debug=True,host='0.0.0.0',port=80)