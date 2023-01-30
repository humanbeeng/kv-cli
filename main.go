package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/humanbeeng/distributed-cache/client"
)

func main() {

	leaderAddr := flag.String("leader", "localhost:3000", "Address of leader")
	flag.Parse()

	fmt.Println("KV Store CLI")

	client, err := client.New(*leaderAddr)

	if err != nil {
		log.Fatal(err)
	}

	for {
		var cmd string
		fmt.Print(":>> ")
		fmt.Scanf("%s", &cmd)

		switch cmd {
		case "exit":
			os.Exit(0)
		case "get":
			{
				var arg string
				fmt.Scan(&arg)
				val, err := client.Get(arg)
				if err != nil {
					fmt.Printf("GET:>> nil\n\n")
					break
				}
				fmt.Printf("GET:>> %v\n\n", val)
				break

			}
		case "set":
			{
				var key string
				var val string
				var ttl int32
				fmt.Scanf("%s %s %d", &key, &val, &ttl)
				err := client.Set(key, val, ttl)
				if err != nil {
					fmt.Println(err.Error())
					break
				}
				fmt.Printf("SET:>>%v\n\n", val)
				break

			}
		case "del":
			{
				var key string
				fmt.Scanf("%s", &key)
				err := client.Del(key)
				if err != nil {
					fmt.Println(err.Error())
					break
				}
				fmt.Printf("DEL:>>%v\n\n", key)
				break
			}

		default:
			fmt.Println("Invalid command")
		}

	}
}
