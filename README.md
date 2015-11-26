# Coco Splunk TCP Forwarder

## Building
```
CGO_ENABLED=0 go build -a -installsuffix cgo -o coco-splunk-tcp-forwarder .

docker build -t coco/coco-splunk-tcp-forwarder .
```

## Description
The Splunk forwarder is a golang application that streams TLS TCP stdin to `ingest.splunk.glb.ft.com:8443`.
Docker images builds a container that forwards the journalctl logs to the Splunk endpoint
 
## Usage ex
e.g. journalctl -f --output=json | ./coco-splunk-tcp-forwarder
