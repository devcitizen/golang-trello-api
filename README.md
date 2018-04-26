# GOLANG TRELLO API - SAMPLE

this is sample code for trello API with golang framework(gin-gonic)

## Getting Started

### Prerequisites

Look Golang docs [here](https://golang.org/doc/install)

```
install golang in your local machine
```


### Installing

Look gin-gonic docs [here](https://github.com/gin-gonic/gin)

after installing golang please follow :

```
clone the repo
```
```
cd repo
```
Create mysql database and setup on main.go line 19 :
```
const (
	MysqlURL = "[username]:[password]@/[name_your_database]?charset=utf8&parseTime=True&loc=Local"
)
```
```
$ go run main.go
```
```
the API run on http://localhost:8080
```

### Usage

Create Users Role API :
```
$ http -v --json POST localhost:8080/api/v1/roles name=admin
```

Register User API :
```
$ http -v --json POST localhost:8080/api/v1/users name=admin email=admin@admin.com role_id=1 password=password address=earth
```

Login User API :
```
$ http -v --json POST localhost:8080/login email=dmin@admin.com password=password address=earth
```


## Built With

* [Gin-gonic](https://github.com/gin-gonic/gin) - The web golang framework used
* [Gorm](http://gorm.io/) - ORM for golang
* [Appleboy/JWT](https://github.com/appleboy/gin-jwt) - JWT Auth

## Authors

* **Okky Muhamad Budiman** - *Initial work* - [Okiww](https://github.com/okiww)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone who's code was used
* Inspiration
* etc
