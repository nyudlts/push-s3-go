## push-s3-go

This repo contains code for pushing files to an Amazon S3 bucket

#### Scripts
```
push-s3:
    pushes a file to an Amazon AWS S3 bucket

    Usage: push-s3 <path-to-file>
    e.g.,  push-s3 ../foo.wav

Configuration file:
    You must have configuration YAML file at this path ${HOME}/.push-s3/config.yaml

---
s3:
  aws_access_key_id: "super secret value"
  aws_secret_access_key: "another super secret value"
  aws_region: "your region"
  bucket: "your target bucket name"
  base_url: "https://s3.amazonaws.com"
```
