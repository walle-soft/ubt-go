package main

import "github.com/walle-soft/ubt-go/log"

func main()  {
	log.SendError(log.H{
		"message": "error test",
	})
}
