// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
    "flag"
    "log"
    "crypto/tls"
    "net/url"
    "os"
    //"os/signal"
    "bufio"
    "time"
    "fmt"
    "crypto/x509"

    "github.com/gorilla/websocket"
    "./cert"
)

var server = "liambeckman.com:8181"
//var server = "localhost:8181"

var addr = flag.String("addr", server, "http service address")

func main() {
    flag.Parse()
    log.SetFlags(0)

    cert, err := cert.Asset("cert.pem")
    if err != nil {
        // Asset was not found.
        log.Fatal(err)
    }


    //interrupt := make(chan os.Signal, 1)
    //signal.Notify(interrupt, os.Interrupt)

    u := url.URL{Scheme: "wss", Host: *addr, Path: "/"}
    log.Printf("connecting to %s", u.String())
    if err != nil {
        log.Fatal(err)
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(cert)


    var d = websocket.Dialer{
            TLSClientConfig: &tls.Config{
                RootCAs:      caCertPool,
            }}

    c, _, err := d.Dial(u.String(), nil)
    if err != nil {
        log.Fatal("dial:", err)
        }

    if err != nil {
        log.Fatal("dial:", err)
    }
    defer c.Close()

    done := make(chan struct{})

    go func() {
        defer close(done)
        for {
            _, message, err := c.ReadMessage()
            if err != nil {
                log.Println("read:", err)
                return
            }
            fmt.Print(string(message))
        }
    }()

    for true {
        scanner := bufio.NewScanner(os.Stdin)
        fmt.Print("> ")
        for scanner.Scan() {
            //fmt.Println(scanner.Text())
            err := c.WriteMessage(websocket.TextMessage, []byte(scanner.Text()))
            if err != nil {
                log.Println("write:", err)
                return
            }
        }

        if scanner.Err() != nil {
            // handle error.
        }

    }

    for {
        select {
        case <-done:
            return

            // Cleanly close the connection by sending a close message and then
            // waiting (with timeout) for the server to close the connection.
            err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
            if err != nil {
                log.Println("write close:", err)
                return
            }
            select {
            case <-done:
            case <-time.After(time.Second):
            }
            return
        }
    }
}

