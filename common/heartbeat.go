package common

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

const SVC = "svc#%s"

func Heartbeat(cli *clientv3.Client, name string) {
	ctx := context.TODO()
	svc := fmt.Sprintf(SVC, name)
	go func() {
		now := time.Now()
		_, err := cli.Put(ctx, svc, fmt.Sprint(now.Unix()))
		if err != nil {
			fmt.Println(svc, "heartbeat error:", err)
		} else {
			fmt.Println(svc, "heartbeat at", now.Unix())
		}
		for t := range time.Tick(time.Second * 10) {
			_, err := cli.Put(ctx, svc, fmt.Sprint(t.Unix()))
			if err != nil {
				fmt.Println(svc, "heartbeat error:", err)
			} else {
				fmt.Println(svc, "heartbeat at", t.Unix())
			}
		}
	}()
}
