package main

import (
  "github.com/urturn/go-phantomjs"
)

func main() {
  p, err := phantomjs.Start();
  if err != nil {
    panic(err);
  }
  defer p.Exit();
  var result interface{}
  err = p.Run("function(){return 2 + 2 }",&result);
}
