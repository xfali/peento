// Copyright (C) 2019, Xiongfa Li.
// All right reserved.
// @author xiongfa.li
// @version V1.0
// Description: 

package main

import (
    "flag"
    "io"
    "log"
    "os"
    "os/signal"
    "peento/api"
    "peento/cluster"
    "peento/config"
    "peento/core"
    "strings"
    "syscall"
)

func main() {
    members := flag.String("join", "", "member list: HOST1:PORT1,HOST2:PORT2,HOST3:PORT3")
    cport := flag.Int("gossip-port", 17001, "gossip port")
    sport := flag.Int("api-port", 18001, "api port")
    flag.Parse()

    conf := &config.Config{}
    conf.Port = *cport
    conf.ApiPort = *sport
    memList := strings.Split(*members, ",")
    for _, v := range memList {
        mem := strings.TrimSpace(v)
        if mem != "" {
            conf.Members = append(conf.Members, mem)
        }
    }
    conf.Delegate = &core.GossipDelegate{}

    var closers []io.Closer
    c1, err1 := cluster.Startup(conf)
    if err1 != nil {
        log.Fatal(err1)
    }
    closers = append(closers, c1)

    c2, err2 := api.HttpStartup(conf)
    if err2 != nil {
        closeAll(closers)
        log.Fatal(err2)
    }
    closers = append(closers, c2)

    handleSignal(closers)
}

func handleSignal(c []io.Closer) {
    quitChan := make(chan os.Signal)
    signal.Notify(quitChan,
        syscall.SIGINT,
        syscall.SIGTERM,
        syscall.SIGHUP,
    )
    <-quitChan

    closeAll(c)
    log.Println("server gracefully shutdown")
    close(quitChan)
}

func closeAll(c []io.Closer) {
    for _, v := range c {
        if err := v.Close(); nil != err {
            log.Fatalf("server shutdown failed, err: %v\n", err)
        }
    }
}
