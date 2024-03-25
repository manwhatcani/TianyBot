package utils

import (
	"github.com/siku2/arigo"
)

func SendAria2(url string) string {
	c, err := arigo.Dial("ws://localhost:16800/jsonrpc", "159357")
	if err != nil {
		panic(err)
	}
	gid, err := c.AddURI(arigo.URIs(url), &arigo.Options{Continue: true})
	Mylog(gid)
	if err != nil {
		panic(err)
	}
	c.Close()
	return gid.GID
}
