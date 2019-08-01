// Copyright (C) 2019, Xiongfa Li.
// All right reserved.
// @author xiongfa.li
// @version V1.0
// Description: 

package core

import (
    "fmt"
    "sync"
)

type GossipDelegate struct {
    mu          sync.Mutex
    meta        []byte
    msgs        [][]byte
    broadcasts  [][]byte
    state       []byte
    remoteState []byte
}

func (m *GossipDelegate) Update(data []byte) {
    m.mu.Lock()
    defer m.mu.Unlock()

    m.state = data
}

func (m *GossipDelegate) NodeMeta(limit int) []byte {
    m.mu.Lock()
    defer m.mu.Unlock()

    fmt.Printf("NodeMeta :%v\n", limit)
    return m.meta
}

func (m *GossipDelegate) NotifyMsg(msg []byte) {
    m.mu.Lock()
    defer m.mu.Unlock()

    cp := make([]byte, len(msg))
    copy(cp, msg)
    fmt.Printf("NotifyMsg :%v\n", string(msg))
    m.msgs = append(m.msgs, cp)
}

func (m *GossipDelegate) GetBroadcasts(overhead, limit int) [][]byte {
    m.mu.Lock()
    defer m.mu.Unlock()

    b := m.broadcasts
    m.broadcasts = nil
    //fmt.Printf("GetBroadcasts :%v %v\n", overhead, limit)
    return b
}

func (m *GossipDelegate) LocalState(join bool) []byte {
    m.mu.Lock()
    defer m.mu.Unlock()

    fmt.Printf("LocalState :%v\n", join)

    return m.state
}

func (m *GossipDelegate) MergeRemoteState(s []byte, join bool) {
    m.mu.Lock()
    defer m.mu.Unlock()

    fmt.Printf("MergeRemoteState :%v %v\n", string(s), join)

    m.remoteState = s
}
