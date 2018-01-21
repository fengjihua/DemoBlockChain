package lib

/*
Server Configuration
*/
var (
	ConfigIsLocal   = false
	ConfigIsDevelop = true
	ConfigIsProduct = false

	ConfigIsDebug   = true || ConfigIsLocal
	ConfigIsRelease = !ConfigIsDebug
)

/*
MongoDB Configuration
*/
var (
	ConfigIsDB  = false
	DbConnMongo = "mongodb://localhost:27017"
)
