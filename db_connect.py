import MySQLdb
def connection():
	conn = MySQLdb.connect(host='localhost', 
		user = 'root',
		passwd = '',
		db = 'Flask_App'
		)
	c= conn.cursor()
	return c,conn
