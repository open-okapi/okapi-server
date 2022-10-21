package service

import (
	"github.com/open-okapi/data-driven-proxies/client"
	res "github.com/open-okapi/data-driven-proxies/request"
)

func ExecHttp(meta res.Meta) {
	var resp = client.Execute(meta)
	println(resp.Error)
}

func ExecWS() {
	println("TO IMPLEMENT")
}

func ExecTcp() {
	println("TO IMPLEMENT")
}

func ExecRpc() {
	println("TO IMPLEMENT")
}
