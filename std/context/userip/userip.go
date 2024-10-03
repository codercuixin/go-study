package userip

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

func FromRequest(req *http.Request)(net.IP, error){
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip1: %q is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return nil, fmt.Errorf("userip2: %q is not IP:port", req.RemoteAddr)
	}
	return userIP, nil 
}

type key int 
const userIpKey key = 0

func NewContext(ctx context.Context, userIp net.IP)(context.Context){
	return context.WithValue(ctx, userIpKey, userIp)
}

func FromContext(ctx context.Context)(net.IP, bool){
	userIP, ok := ctx.Value(userIpKey).(net.IP)
	return userIP, ok 
}