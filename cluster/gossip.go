// Copyright (C) 2019, Xiongfa Li.
// All right reserved.
// @author xiongfa.li
// @version V1.0
// Description: 

package cluster

import (
    "github.com/hashicorp/memberlist"
    "os"
    "peento/config"
    "strconv"
    "time"
)

type members memberlist.Memberlist

type Cluster interface {
    UpdateLocal(meta []byte) error
    UpdateAndWait(meta []byte, timeout time.Duration) error
    Close() error
}

func Startup(conf *config.Config) (Cluster, error) {
    hostname, _ := os.Hostname()
    config := memberlist.DefaultLocalConfig()
    config.Name = hostname + "-" + strconv.Itoa(conf.Port)
    // config := memberlist.DefaultLocalConfig()
    config.BindPort = conf.Port
    config.AdvertisePort = conf.Port
    if conf.Delegate != nil {
        config.Delegate = conf.Delegate
    }
    if conf.EventDelegate != nil {
        config.Events = conf.EventDelegate
    }

    list, err := memberlist.Create(config)
    if err != nil {
        return nil, err
    }

    if len(conf.Members) > 0 {
        list.Join(conf.Members)
    }

    return (*members)(list), nil
}

func (c *members) UpdateLocal(meta []byte) error {
    (*memberlist.Memberlist)(c).LocalNode().Meta = meta
    return nil
}

func (c *members) UpdateAndWait(meta []byte, timeout time.Duration) error {
    (*memberlist.Memberlist)(c).LocalNode().Meta = meta
    return (*memberlist.Memberlist)(c).UpdateNode(timeout)
}

func (c *members) Close() error {
    return (*memberlist.Memberlist)(c).Shutdown()
}
