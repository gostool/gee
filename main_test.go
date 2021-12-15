package gee

import (
	"testing"
	"github.com/gogf/gf/frame/g"
)

func Test_Server(t *testing.T) {
	Server()
}

func Test_G(t *testing.T) {
	g.Server().Run()
}