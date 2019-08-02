// Copyright (C) 2019, Xiongfa Li.
// All right reserved.
// @author xiongfa.li
// @version V1.0
// Description: 

package core

import (
    "github.com/hashicorp/memberlist"
    "log"
    "sync"
)

type handle func(*memberlist.Node)

type NodeDelegate struct {
    mu         sync.Mutex
    enabled    bool
    joinFunc   handle
    leaveFunc  handle
    updateFunc handle
}

func defaultJoin(n *memberlist.Node) {
    log.Printf("defaultJoin meta: %s\n", string(n.Meta))
}

func defaultLeave(n *memberlist.Node) {
    log.Printf("defaultLeave meta: %s\n", string(n.Meta))
}

func defaultUpdate(n *memberlist.Node) {
    log.Printf("defaultUpdate meta: %s\n", string(n.Meta))
}

func DefaultNodeDelegate() *NodeDelegate {
    return &NodeDelegate{
        enabled:    true,
        joinFunc:   defaultJoin,
        leaveFunc:  defaultLeave,
        updateFunc: defaultUpdate,
    }
}

func (d *NodeDelegate) NotifyJoin(n *memberlist.Node) {
    d.mu.Lock()
    defer d.mu.Unlock()
    if d.enabled {
        d.joinFunc(n)
    }
}

// NotifyLeave is invoked when a node is detected to have left.
// The Node argument must not be modified.
func (d *NodeDelegate) NotifyLeave(n *memberlist.Node) {
    d.mu.Lock()
    defer d.mu.Unlock()
    if d.enabled {
        d.leaveFunc(n)
    }
}

// NotifyUpdate is invoked when a node is detected to have
// updated, usually involving the meta data. The Node argument
// must not be modified.
func (d *NodeDelegate) NotifyUpdate(n *memberlist.Node) {
    d.mu.Lock()
    defer d.mu.Unlock()
    if d.enabled {
        d.updateFunc(n)
    }
}
