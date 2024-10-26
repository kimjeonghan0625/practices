package main

import (
	"fmt"
	"strings"
)

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire(any ...string)
}

type ReceivableEvent interface {
	Event
	OnRecv()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) {
	m.listener = listener
}

func (m *Mail) OnRecv() {
	m.listener.OnFire("메일", ":네이버")
}

type KakaoTalk struct {
	listener EventListener
}

func (k *KakaoTalk) Register(l EventListener) {
	k.listener = l
}

func (k *KakaoTalk) OnRecv() {
	k.listener.OnFire("카카오톡")
}

type Alarm struct {}
	
func (a *Alarm) OnFire(any ...string) {
	if len(any) != 0 {
		fmt.Printf("%s에서 알람왓숑\n", strings.Join(any, " "))
		return
	}
	fmt.Println("알람왓숑")
}

type DBInsert struct {}

func (d *DBInsert) OnFire(any ...string) {
	if len(any) != 0 {
		fmt.Printf("%s에서 받은 데이터를 db에 넣었숑\n", strings.Join(any, " "))
		return
	}
	fmt.Println("db에 삽입을 완료햇숑")
}

func main() {
	var revent ReceivableEvent = &Mail{}
	revent = &KakaoTalk{}
	var listener EventListener = &Alarm{}
	listener = &DBInsert{}

	revent.Register(listener)
	revent.OnRecv()
}

