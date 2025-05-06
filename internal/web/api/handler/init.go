package handler

import "sync"

var once sync.Once

func Init() {
	once.Do(func() {})
}
