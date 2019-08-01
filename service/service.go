// Copyright (C) 2019, Xiongfa Li.
// All right reserved.
// @author xiongfa.li
// @version V1.0
// Description: 

package service

import "encoding/json"

type Discover interface {
    Update(data []byte) error
    Merge(data []byte) error
}

type Service struct {
    Addr string
}

type Cluster struct {
    Services []Service
}

func (c *Cluster)Update(data []byte) error {
    n := Cluster{}
    err := json.Unmarshal(data, &n)
    if err != nil {
        return err
    }



    return nil
}

func (c *Cluster)Merge(data []byte) error {
    n := Cluster{}
    err := json.Unmarshal(data, &n)
    if err != nil {
        return err
    }

    for _, v := range n.Services {
        found := false
        for _, old := range c.Services {
            if old.Addr == v.Addr {
                found = true
                break
            }
        }
        if !found {
            c.Services = append(c.Services, v)
        }
    }
    return nil
}
