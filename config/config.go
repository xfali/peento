// Copyright (C) 2019, Xiongfa Li.
// All right reserved.
// @author xiongfa.li
// @version V1.0
// Description: 

package config

import "peento/core"

type Config struct {
    Port     int
    Members  []string
    ApiPort  int
    Delegate *core.GossipDelegate
}
