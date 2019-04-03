package server

import (
	"./orm"
	"encoding/json"
	"github.com/fatih/structs"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"go.mongodb.org/mongo-driver/bson"
)

const SUCCESS = 200
const FAILTURE = 201


/*
	Response of Query interface
 */
type QueryResponse struct {
	Message string           `json:"message"`
	Code int                 `json:"code"`
	Result []orm.QueryResult `json:"result"`
	Rows int                 `json:"rows"`
}

/*
	Response of getting average height interface
 */
type AveHeightResponse struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Result float64 `json:"result"`
}

/*
parse the arguments from the url
 */
func parseUrl(values url.Values)  bson.M {
	m := structs.Map(orm.QueryResult{})
	newMap := make(map[string]interface{}, 0)
	for key, val := range values {
		if len(val) > 0 {
			if _, ok := m[key]; ok {
				if _, ok2 := newMap[key]; !ok2 {
					if key == "GROUNDELEV" || key == "SHAPE_AREA" || key == "HEIGHTROOF" || key == "SHAPE_LEN" || key == "CNSTRCT_YR"{
						newMap[key], _ = strconv.ParseFloat(val[0], 64)
					}else {
						newMap[key] = val[0]
					}
				}
			}
		}
	}
	return newMap
}

/*
handle the CROS
 */
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Runserver(serverPort, dbAddress, dbName, dbPort, collectionName string) {

	m := &orm.MonogoClient{
		dbAddress,
		dbName,
		dbPort,
		60,
		nil,
	}
	m.SetUp()

	queryHandler := func(w http.ResponseWriter, req *http.Request) {
		urlMap := parseUrl(req.URL.Query())
		message := ""
		code := SUCCESS
		res := m.Query(collectionName, urlMap)
		var response []byte
		if len(urlMap) != 0 && len(res) != 0{
			message = "Success"
			code = SUCCESS
			response, _ = json.Marshal(QueryResponse{message, code, res, len(res)})
		}else if len(urlMap) == 0{
			message = "Ivalid input"
			code = FAILTURE
			response, _ = json.Marshal(QueryResponse{message, code, res, len(res)})
		}else {
			message = "No results"
			code = SUCCESS
			response, _ = json.Marshal(QueryResponse{message, code, res, 0})
		}
		enableCors(&w)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

	aveByBaseBblHandler := func(w http.ResponseWriter, req *http.Request) {
		val := req.URL.Query()
		var res float64
		var code int
		var message string
		res = -1
		if baseBbl, ok := val["BASE_BBL"]; ok {
			code = SUCCESS
			res = m.GetAveHeightByBaseBbl(collectionName, baseBbl[0])
			if res == -1 {
				message = "No value"
			}else {
				message = "Success"
			}
		}else {
			message = "Invalid input"
			code = FAILTURE
		}
		response, _ := json.Marshal(AveHeightResponse{message, code, res})
		enableCors(&w)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
	http.HandleFunc("/query/", queryHandler)
	http.HandleFunc("/getAveByBaseBbl/", aveByBaseBblHandler)
	log.Fatal(http.ListenAndServe(":" + serverPort, nil))
}

