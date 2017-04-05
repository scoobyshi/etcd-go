package main

import (
  "log"
  "time"
  "golang.org/x/net/context"
  "github.com/coreos/etcd/client"
)

func main() {
  cfg := client.Config{
    Endpoints: []string{"http://localhost:2379"},
    Transport: client.DefaultTransport,
  }
  c, err := client.New(cfg)
  if err != nil {
    log.Fatal(err)
  }
  kapi := client.NewKeysAPI(c)
  
  resp, err := kapi.Set(context.Background(), "bestLang", "javascript", &client.SetOptions{TTL: 20 * time.Second})
  if err != nil {
    log.Fatal(err)
  } else {
    log.Printf("Set is done. Metadata is %q\n", resp)
  }

  resp, err = kapi.Get(context.Background(), "bestLang", nil)
  if err != nil {
    log.Fatal(err)
  } else {
    log.Printf("Get is done. Metadata is %q\n", resp)
    log.Printf("%q is the value.", resp.Node.Value)
  }

  time.Sleep(10 * time.Second)
  resp, err = kapi.Get(context.Background(), "bestLang", nil)
  log.Printf("check again, %q", resp)
  log.Printf("value is still: %q", resp.Node.Value)

  time.Sleep(12 * time.Second)
  resp, err = kapi.Get(context.Background(), "bestLang", nil)
  log.Printf("is the value still: %q", resp.Node.Value)
}
