// Copyright (C) 2019, Xiongfa Li.
// All right reserved.
// @author xiongfa.li
// @version V1.0
// Description: 

package api

import (
    "fmt"
    "io"
    "net/http"
    "peento/config"
    "peento/core"
    "time"
)

type process struct {
    delegate *core.GossipDelegate
}

func HttpStartup(conf *config.Config) (io.Closer, error) {
    p := process{conf.Delegate}
    http.HandleFunc("/api/", p.process)
    //设置访问的ip和端口
    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", conf.ApiPort),
        Handler:        nil,
        ReadTimeout:    15 * time.Second,
        WriteTimeout:   15 * time.Second,
        IdleTimeout:    15 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    go s.ListenAndServe()

    return s, nil
}

func (p *process) process(resp http.ResponseWriter, req *http.Request) {
    data := getKey(req)
    p.delegate.Update([]byte(data))
    io.WriteString(resp, "ok")
}

func getKey(req *http.Request) string {
    return req.RequestURI[5:]
}
