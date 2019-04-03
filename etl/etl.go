package etl

import (
	"bufio"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)
func dealWithPerLine(s string) []string{
	res := make([]string, 0)
	token := strings.Split(s, ")))\",")
	if len(token) != 2 {
		return res
	}
	token0 := strings.Split(token[0], "(((")
	token1 := strings.Split(token[1], ",")
	if len(token0) != 2 {
		res = append(res, "", "")
		res = append(res, token1...)
		return res
	}
	geoTypeByte := []byte(token0[0])
	geoType := string(geoTypeByte[1:len(geoTypeByte) - 1])
	res = append(res, geoType, token0[1])
	res = append(res, token1...)
	return res
}
func ETL(filename, address, port, dbName, collName string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	appUrl := "mongodb://" + address + ":" + port
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(appUrl))

	collection := client.Database(dbName).Collection(collName)
	collection.Name()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {

		line := scanner.Text()
		token := dealWithPerLine(line)

		if len(token) == 16 {
			SHAPE_LEN, _ := strconv.ParseFloat(token[12], 64)
			SHAPE_AREA, _ := strconv.ParseFloat(token[11], 64)
			HEIGHTROOF, _ := strconv.ParseFloat(token[8], 64)
			GROUNDELEV, _ := strconv.ParseFloat(token[10], 64)
			CNSTRCT_YR, _ := strconv.ParseFloat(token[3], 32)
			res, err := collection.InsertOne(ctx, bson.M{
				"GEOTYPE": token[0],
				"GEOLATLON": token[1],
				"BIN": token[2],
				"CNSTRCT_YR": CNSTRCT_YR,
				"NAME": token[4],
				"LSTMODDATE": token[5],
				"LSTSTATYPE": token[6],
				"DOITT_ID": token[7],
				"HEIGHTROOF": HEIGHTROOF,
				"FEAT_CODE": token[9],
				"GROUNDELEV": GROUNDELEV,
				"SHAPE_AREA": SHAPE_AREA,
				"SHAPE_LEN": SHAPE_LEN,
				"BASE_BBL": token[13],
				"MPLUTO_BBL": token[14],
				"GEOMSOURCE": token[15],
			})
			if err != nil {
				fmt.Println("Insert error",err)
			}else {
				fmt.Println(res.InsertedID)
			}
		}
	}

	if err  := scanner.Err() ; err != nil {
		fmt.Println("Scanner error",err)
		//log.Fatal(err)
	}

}