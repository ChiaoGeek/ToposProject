# ToposProject

<p align="center"><a href="https://topos.com" target="_blank" rel="noopener noreferrer"><img width="300" src="https://topos.com/static/logo-76148d2a15dff2c266e1f9bf32befd89.png" alt="Vue logo"></a></p>

## 1)The structure of this project

       .
       ├── etl                   # Including etl.go. This directory holds the entire etl tools
       |
       etl
       │   ├── etl.go            # A simple ETL process written in Go that extract data from CSV file, clearn the  data, and stores data in a MongoDB.
       server
       │   ├── orm               # A simple tool used to query the MonogDB
       │   ├── server.go         # A small API written in Go that provides query and basic transformations
       │
       runetl.go                 # Used to run etl tool
       │
       runserver.go              # Used to run server API
       |
       .
## 2)Requirements

><a href="https://golang.org/doc/install#install" target ="_blank">Go 1.10 or higher. <a/> <br/>
><a href="https://docs.mongodb.com/manual/installation/" target ="_blank">MongoDB 2.6 and higher.<a/> <br/>
><a href="https://github.com/mongodb/mongo-go-driver" target ="_blank">MongoDB Go Driver v1.0.0 <a/> <br/>


            
## 3)How to run the etl tool and server API

```

git clone https://github.com/ChiaoGeek/ToposProject.git

# Run etl 
# Before run etl you should change the arguments such as database name, file name..
# runetl has a function named ETL("csvFileName", "MongoDBAdress", "MonogoDBPort", "MongoDBDatabaseName", "MongoDBDatabaseCollectionName")
# To make this progress easy, there is no password
go run runetl.go


# run server API
# Before you run server API, you should change the arguments such as database name, port number 
# runserver has a function named Runserver(serverPort, dbAddress, dbName, dbPort, collectionName)
go run runserver.go


```


## 4)The description of server API

1)Query API



## 2)Deployment diagram

<p align="center"><a href="https://raw.githubusercontent.com/USC-CSSL/MapYourMorals/master/readmefile/dd.png?token=AI0fHQMpyzaioVh2dWgepVvfSPMHbamSks5auvKEwA%3D%3D" target="_blank" rel="noopener noreferrer"><img width="1000" src="readmefile/dd.png" alt=""></a></p>

## 3)How to connect to system server

This system is deployed on Google Cloud. You can connect to the server using SSH.

### 1. connecting to the server in your LINUX or MACOS

```
ssh-keygen -t rsa -f ~/.ssh/[filename] -C [yourusername]

chmod 400 ~/.ssh/[filename]

cat ~/.ssh/[filename].pub

// The next step is to copy the contents of your public SSH key file and paste them in
to the instance configure.(You need to login in Google Cloud Platform and find the right
  VM server.)

// Tips: the public key has expired time. After a while, you need to regenerate it.

ssh -i ~/.ssh/[filename].pub [yourusername]@mapyourmorals.org

// Now, you can connect to the server.

sudo su  // get root permission

cd /home/changzhao619/Project/MapYourMoralServer/      //   back-end source code

cd /var/www/html/           // front-end source code

```

### 2.Connecting to the MySQL server.

```

mysql -h 35.230.66.253 -u root -p

// enter the password: MapYourMoralCsslUsc

```
## 3)How to run front end in localhost

```

cd /MapYourMorals/frontend

# install dependencies
npm install

# serve with hot reload at http://localhost:8088
npm run dev


```
## 4)How to upload the front end to the server

```

cd /MapYourMorals/frontend

# build for production with minification
npm run build

cd /MapYourMorals/frontend/dist

(and then uploading index.html and static to the www dir[ /var/www/html/] in the server).

```

## 5)How to run back end


```
cd /home/changzhao619/Project/MapYourMoralServer/

python  index.py  &

(To check if it works, you can run the command:  netstat -apn  | grep 8000)

```

## 6)The structure of front end

        .
        ├── dist                    # including index.html and static file. This directory holds the actual configurations for both the development server and the production webpack build. Normally you don't need to touch these files
        |
        config
        │   ├── index.js            # This is the main configuration file that exposes some of the most common configuration options for the build setup.
        src
        │   ├── assets              # module assets (processed by webpack)
        │   ├── components          # ui components (header, footer)
        |   ├── views               # the ui of the system
        |   ├── jsscript            # js scripts of map, search option and configure.
        ├── static                  # This directory is an escape hatch for static assets that you do not want to process with Webpack. They will be directly copied into the same directory where webpack-built assets are generated.



If you want to change the UI of the system, you should modify the following files:
- **src/views/index/index.vue**: `first page`.
- **src/views/index/left/index.vue**: `filter panel`.
- **src/views/index/map/index.vue**:  `map`.
- **src/views/components/header/index.vue**: `the header of page`.




## 7)Setting up ssl in Nginx

If you migrate the project to other server, you should set up nginx server again to enable https. You should modify the configure file of nginx.

```
server {
#	listen 443 default_server;
#	listen [::]:443 default_server;

	#SSL configuration
	#
	listen 443 ssl default_server;
	listen [::]:443 ssl default_server;
	ssl_certificate 	/var/https/489624b5bf5ab171.crt;
	ssl_certificate_key /var/https/www.mapyourmorals.org.key;
	ssl_protocols       TLSv1 TLSv1.1 TLSv1.2;
	ssl_ciphers         HIGH:!aNULL:!MD5;


	#
	# Note: You should disable gzip for SSL traffic.
	# See: https://bugs.debian.org/773332
	#
	# Read up on ssl_ciphers to ensure a secure configuration.
	# See: https://bugs.debian.org/765782
	#
	# Self signed certs generated by the ssl-cert package
	# Don't use them in a production server!
	#
	# include snippets/snakeoil.conf;

	root /var/www/html;

	# Add index.php to the list if you are using PHP
	index index.html index.htm index.nginx-debian.html;

	server_name www.mapyourmorals.org;


	location / {
		# First attempt to serve request as file, then
		# as directory, then fall back to displaying a 404.
		try_files $uri $uri/ =404;
	}

	location /app {
		#include uwsgi_params;
		proxy_pass http://127.0.0.1:8000;
		#include uwsgi_params;

	#	uwsgi_pass unix:///tmp/uwsgi.sock;
		uwsgi_read_timeout 60;
	}
	# pass PHP scripts to FastCGI server
	#
	#location ~ \.php$ {
	#	include snippets/fastcgi-php.conf;
	#
	#	# With php-fpm (or other unix sockets):
	#	fastcgi_pass unix:/var/run/php/php7.0-fpm.sock;
	#	# With php-cgi (or other tcp sockets):
	#	fastcgi_pass 127.0.0.1:9000;
	#}

	# deny access to .htaccess files, if Apache's document root
	# concurs with nginx's one
	#
	#location ~ /\.ht {
	#	deny all;
	#}
}


```

## 8)Some tools used to insert csv data to MySQL

There are some useful tools used to inserting csv data to MySQL.

    .
    ├── fullCounties.py                 # full_demographic_table_for_counties.csv
    ├── fullState.py                    # full_demographic_table_for_states.csv
    ├── staticCounties.py               # static_county_estimates_table.csv
    ├── staticStates.py                 # static_state_estimates_table.csv


---

If you have any questions, please feel free to contact me.

Name: Zhao Chang

Email: zhaochan@usc.edu || zhao-chang@outlook.com

<p align="center">:smiley:</p>
