// Copyright © 2013, 2014, The Go-LXC Authors. All rights reserved.
// Use of this source code is governed by a LGPLv2.1
// license that can be found in the LICENSE file.

//go:build linux && cgo
// +build linux,cgo

package main

import (
	"log"

	"github.com/lxc/go-lxc"
)

func main() {
	c := lxc.Containers()
	for i := range c {
		log.Printf("%s\n", c[i].Name())
		l, err := c[i].Snapshots()
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		}

		for _, s := range l {
			log.Printf("Name: %s\n", s.Name)
			log.Printf("Comment path: %s\n", s.CommentPath)
			log.Printf("Timestamp: %s\n", s.Timestamp)
			log.Printf("LXC path: %s\n", s.Path)
			log.Println()
		}
		c[i].Release()
	}
}
