package tool

import "sync"

type GoMap struct {
	data map[interface{}]interface{}
	lock *sync.RWMutex
}

func NewGoMap(size int) *GoMap {
	return &GoMap{
		data: make(map[interface{}]interface{}, size),
		lock: new(sync.RWMutex),
	}
}

func (gm *GoMap) Add(key interface{}, value interface{}) {
	gm.lock.Lock()
	defer gm.lock.Unlock()
	gm.data[key] = value
}

func (gm *GoMap) Get(key interface{}) (interface{}, bool) {
	gm.lock.Lock()
	defer gm.lock.Unlock()
	v, ok := gm.data[key]
	return v, ok
}

func (gm *GoMap) Exist(key interface{}) bool {
	gm.lock.Lock()
	defer gm.lock.Unlock()
	_, ok := gm.data[key]
	return ok
}

func (gm *GoMap) Delete(key interface{}) {
	gm.lock.Lock()
	defer gm.lock.Unlock()
	delete(gm.data, key)
}

func (gm *GoMap) Remove() {
	gm.lock.Lock()
	defer gm.lock.Unlock()
	gm.data = make(map[interface{}]interface{})
}
