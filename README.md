# b64
A base64 Encode / URL Encode  / Decode / URL Decode written in Go

To install:
```
brew tap mahe54/homebrew-b64
brew install b64

```
base64 encode:
```
b64 -e SomethingToEncode
```

base64 url encode:
```
b64 -e -u SomethingToEncode
```

base64 decode:
```
b64 -d U29tZXRoaW5nVG9FbmNvZGU=
```

base64 url decode:
```
b64 -d -u U29tZXRoaW5nVG9FbmNvZGU
```