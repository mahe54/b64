# b64
A base64 Raw / URL Encode / Decode, tool written in Go.
It does not add the trailing new line character (K).

##### To install:
```
brew tap mahe54/homebrew-b64
brew install b64

```
##### base64 std encode:
```
b64 -e SomethingToEncode
echo "SomethingToEncode" | b64 -e
```

##### base64 std decode:
```
b64 -d U29tZXRoaW5nVG9FbmNvZGU=
echo "U29tZXRoaW5nVG9FbmNvZGU=" | b64 -d
```

##### base64 url encode (raw):
```
b64 -e -u SomethingToEncode
echo "SomethingToEncode" |  b64 -e -u
```

##### base64 url decode (raw):
```
b64 -d -u U29tZXRoaW5nVG9FbmNvZGU
echo "U29tZXRoaW5nVG9FbmNvZGU" | b64 -d -u
```