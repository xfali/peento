// Copyright (C) 2019, Xiongfa Li.
// All right reserved.
// @author xiongfa.li
// @version V1.0
// Description: 

package cluster

import (
    "github.com/hashicorp/memberlist"
    "io"
    "os"
    "peento/config"
    "strconv"
)

type Closable memberlist.Memberlist

func Startup(conf *config.Config) (io.Closer, error) {
    hostname, _ := os.Hostname()
    config := memberlist.DefaultLocalConfig()
    config.Name = hostname + "-" + strconv.Itoa(conf.Port)
    // config := memberlist.DefaultLocalConfig()
    config.BindPort = conf.Port
    config.AdvertisePort = conf.Port
    config.Delegate = conf.Delegate

    list, err := memberlist.Create(config)
    if err != nil {
        return nil, err
    }

    if len(conf.Members) > 0 {
        list.Join(conf.Members)
    }

    return (*Closable)(list), nil
}

func (c *Closable) Close() error {
    return (*memberlist.Memberlist)(c).Shutdown()
}

