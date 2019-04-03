package main

import "./etl"
func main()  {
	etl.ETL("../../data/building.csv", "127.0.0.1", "27017", "Topos", "BuildingFootprints3")
}