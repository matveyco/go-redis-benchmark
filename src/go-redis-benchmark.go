package main

import (
	"fmt"
	"time"
	"sync"
	"gopkg.in/redis.v3/redis"
	uuid "github.com/matveyco/go.uuid"
	"os"
//	"flag"
)

func WriteToRedisWithConnection(client *redis.Client, wg * sync.WaitGroup ) {
	key := "link:"+uuid.NewV4().String()
	client.HMSet(key,
		"clickurl",			"http://mylinksuggest.com/result/?affiliate=36046&subid=311764&subsid=0&terms=casino&p=1&clickid=NDgwMjUzYmJiZjk1NTNmN2RiYzQzYjU2ODYwMzNkNTY6c1RPNlI0ZDoxNTMzMi4",
		"bid_original",		"0.8",
		"bid",				"0.7",
		"user_agent",		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.152 Safari/537.36",
		"user_ip",			"123.222.123.221",
		"user_referrer",	"http://mylinksuggest.com/result/?affiliate=36046&subid=311764&subsid=",
		"user_keyword",		"casino+online+bonus",
		"clicked",			"false",
		"sent",				"false",
		"wm_id",			"5345",
		"wm_subid",			"6565346",
		"provider_name",	"agnnet_1",
		"timestamp",		"3345364564575.4353456457",
	)
	expireTime := time.Duration(30) * time.Second
	client.Expire(key,expireTime)
	wg.Done()
}

var hostVar string
var portVar string

//func init() {
//	flag.StringVar(&hostVar, "h", "127.0.0.1", "connect to redis host")
//	flag.StringVar(&portVar, "p", "6379", "connect to redis port")
//}

func main() {

//	flag.Parse()
	hostVar = os.Getenv("CONN_HOST")
	portVar = os.Getenv("CONN_PORT")

	if (hostVar == "") {
		hostVar = "127.0.0.1"
	}
	if (portVar == "") {
		portVar = "6379"
	}

	conn := hostVar + ":" + portVar
	//conn := "127.0.0.1:6379"
	//	conn := "192.168.59.103:6370"
	fmt.Println("CONNECTION:",conn)

	client := redis.NewClient(&redis.Options{
		//	client := redis.New (&redis.Options{
		//		Network:  "tcp",
	Addr:     conn,
	Password: "", // no password set
	DB:       0,  // use default DB
})
	var wg sync.WaitGroup

	connStart := time.Now()
	for i:=0; i<5000; i++{
		wg.Add(1)
		go WriteToRedisWithConnection(client,&wg)
	}
	wg.Wait()


	connElapsed := time.Since(connStart)
	fmt.Println("USING CONN : Written 5000 records to redis: ",connElapsed)

}
