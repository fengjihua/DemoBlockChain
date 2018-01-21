package lib

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/zheng-ji/goSnowFlake"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
ToInt32 :String To Int32
*/
func ToInt32(s string) (bool, int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return false, 0, err
	}
	return true, i, nil
}

/*
ToInt64 :String To Int64
*/
func ToInt64(s string) (bool, int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false, 0, err
	}
	return true, i, nil
}

/*
IntToString :Int To String
*/
func IntToString(i int) string {
	return strconv.Itoa(i)
}

/*
Int64ToString :Int64 To String
*/
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// func ToString(i int64) string {
// 	return strconv.FormatInt(i, 10)
// }

/*
ToObjectID :String To bson.ObjectId
*/
func ToObjectID(s string) (bson.ObjectId, error) {
	if len(s) != 24 {
		return bson.NewObjectId(), mgo.ErrNotFound
	}
	return bson.ObjectIdHex(s), nil
}

/*
MD5 :Crypt to MD5
*/
func MD5(s string) string {
	data := []byte(s)
	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash) //将[]byte转成16进制
}

/*
HandleError :Handle all errors, return status code
*/
func HandleError(err error) int {
	if err != nil {
		Log.Error(err)
	}
	status := handleErrorStatus(err)

	return status
}

func handleErrorStatus(err error) int {
	status := StatusUnknown
	if err != nil {
		if err == mgo.ErrNotFound {
			status = StatusNotFound
		}
	} else {
		status = StatusSuccess
	}

	return status
}

/*
GetStatusMessage :Handle all status code, return status message
*/
func GetStatusMessage(status int) string {

	message := ""
	switch status {
	case StatusSuccess:
		message = "success"
	case StatusBad:
		message = "bad request"
	case StatusNoAuth: //	用于API访问鉴权
		message = "请先登录授权"
	// case StatusExpired:
	// 	message = "登录身份已过期"
	case StatusNotFound:
		message = "数据不存在"
	// case StatusInvalidAuth: //	用于账号登录
	// 	message = "用户名或密码错误"
	case StatusUnknown:
		message = "服务器忙"
	}
	return message
}

var iw, _ = goSnowFlake.NewIdWorker(1)

/*
GetNewUID :Get New Unique ID
*/
func GetNewUID() (int64, error) {
	if iw == nil {
		iw, _ = goSnowFlake.NewIdWorker(1)
		// HandleError(err)
	}
	return iw.NextId()
}

/*
GetCurrentDirectory :Get current running directory
*/
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
