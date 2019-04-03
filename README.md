<p align="center">Apply for BACK-END ENGINEER internship</p>
<p align="center"><a href="https://topos.com" target="_blank" rel="noopener noreferrer"><img width="300" src="https://topos.com/static/logo-76148d2a15dff2c266e1f9bf32befd89.png" alt="Vue logo"></a></p>

## 1)The structure of this project

       .
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

```bash

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

### 1 Query API

This API supports both single query parameter and multiple parameters.

##### a.

| API usage                           	| Description                             	| Response type      	|
|-------------------------------------	|-----------------------------------------	|------------------	|
| /query/?LSTSTATYPE=Constructed      	| query by Feature last status type.      	| application/json 	|
| /query/?BIN=3245111                 	| query by Building Identification Number 	| application/json 	|
| /query/?FEAT_CODE=2100              	| query by FEAT_CODE                      	| application/json 	|
| ...                                 	| ...                                     	| ...              	|
| /query/?FEAT_CODE=2100&GROUNDELEV=6 	|  query by FEAT_CODE and GROUNDELEV      	| application/json 	|

##### b. response

```go
type QueryResponse struct {
	Message string           `json:"message"` # message for users
	Code int                 `json:"code"`    # response code (200 success, 201 failure)
	Result []orm.QueryResult `json:"result"`  # results (array)
	Rows int                 `json:"rows"`    # the number of results
}
```

##### c. response example

Request : /query/?BIN=3245111    
Response: 

```json
{
  "message": "Success",
  "code": 200,
  "result": [
    {
      "id": "5ca3f0fa76fd1accaae70974",
      "LSTMODDATE": "08/22/2017 12:00:00 AM +0000",
      "CNSTRCT_YR": 1928,
      "LSTSTATYPE": "Constructed",
      "FEAT_CODE": "2100",
      "GROUNDELEV": 6,
      "SHAPE_AREA": 1167.66412052299,
      "MPLUTO_BBL": "",
      "GEOMSOURCE": "Photogramm",
      "GEOLATLON": "-73.96113466505085 40.57743931616439, -73.96115106427175 40.577438626506336, -73.9611482066905 40.57739910513132, -73.96113180866084 40.57739979388887, -73.9611262071817 40.577322309983394, -73.96119013863581 40.577319622821236, -73.96120349754274 40.57750443832977, -73.9611403239488 40.57751761146039, -73.96113466505085 40.57743931616439",
      "HEIGHTROOF": 37.5,
      "BIN": "3245111",
      "DOITT_ID": "786626",
      "SHAPE_LEN": 183.80050188222,
      "BASE_BBL": "3086910048",
      "GEOTYPE": "MULTIPOLYGON"
    }
  ],
  "rows": 1
}
``` 
### 2 Transform API

This API can get the average height of the buildings by BASE_BBL

##### a.

| API usage                             	| Description                                         	| return type      	|
|---------------------------------------	|-----------------------------------------------------	|------------------	|
| /getAveByBaseBbl/?BASE_BBL=3086910048 	| get the average height of the buildings by BASE_BBL 	| application/json 	|


##### b. response

```go
type AveHeightResponse struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Result float64 `json:"result"`
}
```

##### c. response example

Request : /getAveByBaseBbl/?BASE_BBL=3086910048 	
Response: 

```json
{
  "message": "Success",
  "code": 200,
  "result": 37.5
}
``` 


## 5)API demo

### 1. Query API


<img src="http://chiao.me/swagger/gif/1.gif" alt="swagger"  />


### 2. Transform API

![Alt Text](http://chiao.me/swagger/gif/2.gif)



## 6)Visualizing and interacting with the API by Swagger tool

To facilitate the user to use the API, I wrote an <a href="https://chiao.me/swagger/" target="_blank" >API document </a>based on <a href="https://swagger.io/" target="_blank">Swagger</a>


You can click the link (<a href="https://chiao.me/swagger/" target="_blank" >https://chiao.me/swagger/</a>) to see the UI.

<img src="http://chiao.me/swagger/gif/swagger.png" alt="swagger"  />

## 7) Future Plans

There are some improvements I should do if I have enough time.

> Add more kinds of query APIs

> Support more powerful transformations 

> Consider the efficiency of APIs

> Consider the security of APIs

> Support the multiple-source data for ETL tools


