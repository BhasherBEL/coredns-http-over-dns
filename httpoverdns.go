package httpoverdns

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"

	"github.com/miekg/dns"
)

var log = clog.NewWithPlugin("httpoverdns")

type HTTPOverDNS struct {
	Next plugin.Handler
}

func (e HTTPOverDNS) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}

	if state {

	}

	return plugin.NextOrFailure(e.Name(), e.Next, ctx, pw, r)
}

func (e HTTPOverDNS) Name() string { return "HTTPOverDNS" }
