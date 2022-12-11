package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want) //打印错误日志
	}
	t.Log(got) //打印日志 非必须
}
