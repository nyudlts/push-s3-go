## push-s3-go

This repo contains a super-basic code for pushing files to an Amazon S3 bucket

#### Script
`push-s3-go:`
    pushes a file to an Amazon AWS S3 bucket

    ```
    Usage: push-s3-go <path-to-file> <S3 key>
    e.g.,  push-s3-go ../foo.wav a/b/c/d/foo.wav
    ```

#### Configuration file:
    You must have YAML configuration file at this path `${HOME}/.push-s3-go/config.yaml`
    Please see refer to the (`config.yaml.template`)[./config.yaml.template] file for 
    details.


#### Installation
    Assumes that you have (Golang)[https://golang.org/] installed.
    `$ go get github.com/nyudlts/push-s3-go`
    `$ cd ${GOPATH}/src/github.com/nyudlts/push-s3-go`
    `$ go install push-s3-go.go`

