package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"msg":"Hello world!"}`))
}

func main() {
	http.HandleFunc("/", Test)

	_ = http.ListenAndServe(":9001", nil)
	return
	//var (
	//	err error
	//	user string
	//)
	//conn := Conn()
	//defer conn.Close()
	//_,err = conn.Do("set","user","cowboy","ex",60)
	//if err != nil {
	//	fmt.Println("redis set failed:",err)
	//	return
	//}
	//user,err = redis.String(conn.Do("GET","user"))
	//if err != nil {
	//	fmt.Println("redis get failed:",err)
	//	return
	//}
	//fmt.Println("Get user:%v",user)
}

func Conn() redis.Conn {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil
	}
	//defer conn.Close()
	return conn
}
