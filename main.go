package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"io/ioutil"
	"log"
	"syscall"
	"wechat/router"
)

func main() {
	routers := router.SetRouter()
	port := ":9000"
	server := endless.NewServer(port, routers)
	server.BeforeBegin = func(add string) {
		pid := syscall.Getpid()
		log.Printf("Actual pid is %d", pid)
		ioutil.WriteFile("pid", []byte(fmt.Sprintf("%d", pid)), 0777)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err.Error())
	}
}
