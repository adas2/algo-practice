# Sample go_server
This is a sample README for myserver

## To build and run 
`go build` 
`./go_server <PORT>`

## To test
In a browser open URL:
 `http://<IP Addr>:<PORT>/v1` --> prints `hello`

 `http://<IP Addr>:<PORT>/cert` --> outputs list of available certificates for the host

 `http://<IP Addr>:<PORT>/cert/<hostname>` --> creates a new certificate for host
