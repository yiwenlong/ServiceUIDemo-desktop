### How to build this App
#### Install therecipe/qt
```shell script
export GO111MODULE=off
xcode-select --install
go get -v github.com/therecipe/qt/cmd/...
$(go env GOPATH)/bin/qtsetup test && $(go env GOPATH)/bin/qtsetup -test=false
```
See Also: https://github.com/therecipe/qt
#### Build & Run
```shell script
qtdeploy test desktop
```