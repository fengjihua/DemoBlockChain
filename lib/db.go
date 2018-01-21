package lib

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

// GlobalMgoSession :
var GlobalMgoSession *mgo.Session

func init() {
	if !ConfigIsDB {
		return
	}

	// globalMgoSession, err := mgo.DialWithTimeout("localhost", 10*time.Second)
	conn := DbConnMongo
	globalMgoSession, err := mgo.DialWithTimeout(conn, 10*time.Second)
	if err != nil {
		panic(err)
	}
	GlobalMgoSession = globalMgoSession
	GlobalMgoSession.SetMode(mgo.Monotonic, true)
	//default is 4096
	GlobalMgoSession.SetPoolLimit(300)
}

/*
CloneSession :Get MongoDB Global Session
*/
func CloneSession() *mgo.Session {
	return GlobalMgoSession.Clone()
}

/*
Get :Get Session and Database
*/
func Get() (*mgo.Session, *mgo.Database) {
	session := CloneSession() //调用这个获得session
	db := session.DB("DemoBlockChain")

	return session, db
}
