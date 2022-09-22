package structure

import "sync"

type ChanMgr struct {
	C    chan int
	once sync.Once
}

func (cm *ChanMgr) SafeClose() {
	cm.once.Do(func() { close(cm.C) })
}

func NewChanMgr() *ChanMgr {
	return &ChanMgr{C: make(chan int)}
}
