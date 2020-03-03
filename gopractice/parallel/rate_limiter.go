package main

import (
    "fmt"
    "time"
)

type Request struct {
    ID  int
    Msg string
}

type Response struct {
    Msg     string
    Success bool
}

// APi with rate limitation feature, returns bad response if limits crossed
func ProcessApi(req *Request, denied bool) *Response {
    // create dummy and send response
    rsp := new(Response)
    // get request
    if denied {
        fmt.Println("Denied request:", req.ID)
        rsp.Msg = "Server Overload"
        rsp.Success = false
        return rsp
    }

    fmt.Println("Processing request:", req.ID)
    // call API here
    time.Sleep(100 * time.Millisecond)
    rsp.Msg = "Some response"
    rsp.Success = true

    return rsp
}

// actual rate limiting logic
func RateLimiterApi(
    reqChannel chan *Request,
    respChannel chan *Response,
    limiter chan time.Time,
    done chan bool,
) {

    // processing := 0
    // denyRequest := false

    for {
    loop:
        select {
        case <-done:
            fmt.Println("Done: Flush All Pending Requests")
            return
        case t := <-limiter:
            fmt.Println("TICKER", t, "start processing")
            // denyRequest = true
            // for limit := 0; limit < 10; limit++ {
            if req, ok := <-reqChannel; ok {
                go func(rq *Request, flag bool) {
                    rsp := ProcessApi(rq, flag)
                    respChannel <- rsp
                }(req, false)
            } else {
                // if !ok reqChannel is closed
                // break loop
                reqChannel = nil
            }
            // }
        default:
            // Failed API calls till next tick
            // fmt.Println("Waiting till next tick", processing)
            if req, ok := <-reqChannel; ok {
                go func(req *Request, flag bool) {
                    rsp := ProcessApi(req, flag)
                    respChannel <- rsp
                }(req, true)
            } else {
                break loop
            }
        }

    }
    // return
}

func RateLimiter2(
    reqChannel chan *Request,
    respChannel chan *Response,
    burstyLimiter chan time.Time,
    done chan bool,
) {

    for {
        // loop:
        select {
        case <-burstyLimiter:
            fmt.Println("Case 1 limiter: ")
            for i := 0; i < 5; i++ {
                if req, ok := <-reqChannel; ok {
                    fmt.Println("processing request", req.ID, time.Now())
                } else {
                    fmt.Println("Case 1: no valid jobs")
                    reqChannel = nil
                    // break loop
                }
            }
        case <-done:
            fmt.Println("DONE")
            return

        default:
            // wait
            // if req, ok := <-reqChannel; ok {
            //     fmt.Println(" case default: Deny request", req.ID)
            // }
            // else {
            //     // channel closed
            //     // reqChannel = nil
            //     break loop
            // }
        }
    }
}

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    start := time.Now()
    // define channels for each API based on limit
    // test for one API first
    loginCh := make(chan *Request, 10000)
    respChannel := make(chan *Response, 10000)
    done := make(chan bool)
    timeChannel := make(chan time.Time, 5)

    // create some dummy requests
    go createRequests(loginCh)

    // add current time to kickstart the process
    // go func() {
    //     timeChannel <- time.Now()
    //     for t := range time.Tick(1 * time.Second) {
    //         timeChannel <- t
    //     }
    // }()
    rate := time.Second / 2 // x requests per second
    tick := time.NewTicker(rate)
    defer tick.Stop()
    go func() {
        for t := range tick.C {
            select {
            case timeChannel <- t:
                // bursts of 10 for every tick
                // for i := 0; i < 10; i++ {
                //     timeChannel <- time.Now()
                // }
            default:
            }
        } // does not exit after tick.Stop()
    }()

    // kill resp and req channels after fixed time
    go func() {
        // kill req
        time.Sleep(5 * time.Second)
        done <- true
        // kill resp
        // time.Sleep(2 * time.Second)
        // done <- true
    }()

    // process request through rate limiter helper
    // RateLimiterApi(loginCh, respChannel, timeChannel, done)
    // RateLimiter2(loginCh, respChannel, timeChannel, done)

    for req := range loginCh {
        <-timeChannel // rate limit our Service.Method RPCs
        // fmt.Println("Processing request for", req.ID)
        go func() {
            rsp := ProcessApi(req, false)
            respChannel <- rsp
        }()
    }

    // print response channel
    // displayResults(respChannel, done)

    elapsed := time.Since(start)
    fmt.Printf("Total time: %v\n", elapsed)

}

func displayResults(respChannel chan *Response, done chan bool) {
    // loop:
    for {
        select {
        case <-done:
            fmt.Println("Resp DONE")
            return
        // case resp, ok := <-respChannel:
        case resp, ok := <-respChannel:
            if ok {
                fmt.Println(resp)
            }
        default:
            // fmt.Println("Waiting")
        }
    }
    // return
}

func createRequests(reqChannel chan *Request) {
    // add reqs
    // time.Sleep(2 * time.Second)
    for i := 0; i < 25; i++ {
        req := &Request{
            ID:  i,
            Msg: "some req",
        }
        reqChannel <- req
    }

    // wait before adding new batch of requests
    fmt.Println("Adding more jobs at", time.Now())
    time.Sleep(3 * time.Second)
    for i := 25; i < 50; i++ {
        req := &Request{
            ID:  i,
            Msg: "some req",
        }
        reqChannel <- req
    }

    /*
       time.Sleep(1 * time.Second)
       for i := 200; i < 300; i++ {
           req := &Request{
               ID:  i,
               Msg: "some req",
           }
           reqChannel <- req
       }
    */

    // close channel to avoid deadlock
    close(reqChannel)
}
