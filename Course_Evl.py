from flask import Flask,render_template,flash,request,url_for,redirect,session,send_file,send_from_directory,jsonify
from db_connect import connection
from MySQLdb import escape_string as thwart #avoid injection attack
from passlib.hash import sha256_crypt
import gc
import smtplib
from flask_mail import Mail, Message
import pygal

app = Flask (__name__,instance_path='/Users/xiangcui/Desktop/DataBase/Intro_To_Database/Project1_With_Flask/protected/')

app.secret_key= 'aearlfkdsjlejrflakdjfkafkdlajk'
"""
app.config.update(
	DEBUG=True,
	MAIL_SERVER='smtp.gmail.com',
	MAIL_PORT=465,
	MAIL_USE_SSL=True,
	MAIL_USERNAME='xc1008@nyu.edu',
	MAIL_PASSWORD='Heizhenzhu123'
	)
mail = Mail(app)"""

user = {
	'username':'xc1008',
	'password':'1234'
}

@app.route('/',methods=['GET','POST'])
def hello():
	if request.method == 'POST':
		
		
		if request.form['confirm_password'] == '': #if not confirm password login

			c,con = connection()
			if request.method == 'POST':

				data = c.execute("SELECT * FROM users WHERE username = '%s'"%thwart(request.form['username']))

				if  int(data) == 0:
					flash('invalid username or password, try again')
					c.close()
					con.close()
					return redirect(url_for('Invalid'))
				elif sha256_crypt.verify(request.form['password'],sha256_crypt.encrypt(request.form['password'])):
					data = c.fetchone()[2]
					session['logged_in']=True
					session['username'] = request.form['username']
					#flash('You ar now logged in')
					c.close()
					con.close()
					return redirect(url_for('Student'))
				else:
					c.close()
					con.close()
					return redirect(url_for('Invalid'))

		else: #else register
			username = request.form['username_reg']
			password_reg = request.form['password_reg']
			confirm_password = request.form['confirm_password']
			email = request.form['email']

			if password_reg!=confirm_password:
				flash('Confirm your password!')
				global error
				error = 'Confirm your password!'
				return redirect(url_for('Invalid'))
			c,conn = connection()
			c.execute("use Flask_App;")
			x = c.execute("SELECT * FROM users WHERE username = '%s'"%username)
			
			if int(x)>0:
				flash('The username is already taken!')
				c.close()
				conn.close()
				global error
				error = 'The username is already taken!'
				return redirect(url_for('Invalid'))
			else:
				encrpyt_password = sha256_crypt.encrypt(password_reg)
				A= "INSERT INTO users (username,password,email,tracking) VALUES (%s,%s,%s,%s);"%(thwart(username),thwart(encrpyt_password),thwart(email),'null')
				
				c.execute("INSERT INTO users (username, password, email, tracking) VALUES (%s, %s, %s, %s)",
				(thwart(username), thwart(encrpyt_password), thwart(email), thwart("/introduction-to-python-programming/")))
	               
				conn.commit()
				c.close()
				conn.close()
				flash('registered successfully!')
				return redirect(url_for('Professor'))

	return render_template("index.html")


@app.route("/Professor/")
def Professor():
	return render_template('Prefessor_Page.html')


@app.route("/student/")
def Student():
	return render_template('Student.html')

@app.route("/Invalid/")
def Invalid():
	return str(error)

@app.route("/gan/",methods=['GET','POST'])
def cao():
	if request.method == 'post':
		print('gan')
	return render_template('try.html')


@app.route('/submit/')
def submit():
	return render_template('submit.html')


@app.errorhandler(404)
def error(e):
	return render_template('404.html')

"""
@app.route('/send-mail/')
def send_mail():
	try:
		msg = Message('Send Mail Tutorial!',
			sender = 'xc1008@nyu.edu',
			recipients = ['xc1008@nyu.edu']
			)
		msg.body = "Yo What's update"
		mail.send(msg)
		return 'successfully send'

	except Exception as e:
		return str(e)"""


"""
@app.route('/index/') #login and sign up
def index():
	return render_template("index.html")




@app.route("/Student/")
def Professor():
	return 'hello Student'"""











	