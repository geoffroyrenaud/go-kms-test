# go-kms-test

My first GO program to begin
It will encrypt or decrypt message with AWS KMS

## Usage

```
./go-kms-test -h
Usage of ./go-kms-test:
  -decrypt string
        Decrypt message (base64 format)
  -encrypt string
        Encrypt message
  -keyid string
        AWS KeyID (required)
```

```
./go-kms-test -keyid 1234abcd-12ab-34cd-56ef-1234567890ab -encrypt "Hello world !"
```
