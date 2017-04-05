package main

import (
  "log"
  "time"
  "fmt"
  "golang.org/x/net/context"
  "github.com/coreos/etcd/clientv3"
)

func main() {

  cfg := clientv3.Config{
    Endpoints: []string{"localhost:2379"},
    DialTimeout: 5 * time.Second,
  }

  cli, err := clientv3.New(cfg)

  if err != nil {
    // handle error
  }
  defer cli.Close()

  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  resp, err := cli.Get(ctx, "/services/", clientv3.WithPrefix())
  cancel()
  if err != nil {
    log.Fatal(err)
  }
  for _, ev := range resp.Kvs {
    fmt.Printf("%s : %s\n", ev.Key, ev.Value)
  }
}
