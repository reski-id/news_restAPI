

<h3 align="center">Portal Rest API <br>
<h5 align="center" >Golang Echo Clean Architecture <h5>
<br>
</h4>
<p align="left">
<h2>
  Content <br></h2>
  • Key Features <br>
  • Installing Using Github<br>
  • Installing Using Docker<br>
  • End Point<br>
  • Swagger Open API<br>
  • Technologi that i use<br>
  • Contact me<br>
</p>


## 📱 Features

* register
* login
* News
* Point


## ⚙️ Installing and Runing from Github

installing and running the app from github repository <br>
To clone and run this application, you'll need [Git](https://git-scm.com) and [Golang](https://go.dev/dl/) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/reski-id/news_restAPI.git

# Go into the repository
$ cd news_restAPI

# Install dependencies
$ go get

# load .env env 
# use bash cmd and type this..
$ source .env

# Run the app
$ go run main.go
```

> **Note**
> Make sure you allreaady create database for this app see local `.env` file.


## ⚙️ Installing and Runing with Docker
if you are using docker or aws/google cloud server you can run this application by creating a container. <br>

```bash
# Pull this latest app from dockerhub 
$ docker pull programmerreski/portal

# if you have mysql container you can skip this
$ docker pull mysql

$ docker run --name mysqlku -p 3306:3306 -d -e MYSQL_ROOT_PASSWORD="yourmysqlpassword" mysql 

# create app container
$ docker run --name wilapi -p 80:8000 -d --link mysqlku -e SECRET="secr3t" -e SERVERPORT=8000 -e Name="portal" -e Address=mysqlku -e Port=3306 -e Username="root" -e Password="yourmysqlpassword" programmerreski/portal

# Run the app
$ docker logs wilapi
```

## 📜 End Point  

User
| Methode       | End Point      | used for            | Using JWT
| ------------- | -------------  | -----------         | -----------
| `POST`        | /register      | Register            | No
| `POST`        | /login         | Login               | No 
| `GET`         | /users         | Get my Profile      | Yes
| `DELETE`      | /users         | Delete my Profile   | Yes
| `PUT`         | /users         | Update my Profile   | Yes

News/Post
| Methode       | End Point       | used for                | Using JWT
| ------------- | -------------   | -----------             | -----------
| `POST`        | /post       | Add New Post  | Yes
| `GET`         | /post       | Get Post            | Yes
| `DELETE`      | /post/:id   | Delete Post   | Yes
| `PUT`         | /post/:id   | Update Post   | Yes

Point
| Methode       | End Point       | used for                | Using JWT   | Only Accessible by
| ------------- | -------------   | -----------             | ----------- | -----------
| `POST`        | /point           | Add New Point       | Yes        | Admin Only
| `GET`         | /point           | Get Certain Point                | Yes        | Admin & User
| `DELETE`      | /point/:id       | Delete Point        | Yes        | Admin Only
| `PUT`         | /point/:id       | Update Point        | Yes        | Admin Only


## Swagger

you can find swagger link here:

[Swagger Open Api - Portal Rest API](https://app.swaggerhub.com/apis-docs/webdeveloper.reski/portal/1.0.0)

## 🛠️ Technologi

This software uses the following Tech:

- [Golang](https://go.dev/dl/)
- [Echo Framework](https://echo.labstack.com/)
- [Gorm](https://gorm.io/index.html)
- [mysql](https://www.mysql.com/)
- [Linux](https://www.linux.com/)
- [Docker](https://www.docker.com/)
- [Dockerhub](https://hub.docker.com/u/programmerreski)
- [Mockery Unit Testing](https://github.com/vektra/mockery)
- [Git Repository](https://github.com/reski-id)
- [Trunk Base Development](https://trunkbaseddevelopment.com/)
- [JSON Web Tokens (JWT)](https://jwt.io/)

## 📱 Contact me
feel free to contact me ... 
- Email programmer.reski@gmail.com 
- [Linkedin](https://www.linkedin.com/in/reski-id)
- [Github](https://github.com/reski-id)
- Whatsapp <a href="https://wa.me/+6281261478432?text=Hello">Send WhatsApp Message</a>
