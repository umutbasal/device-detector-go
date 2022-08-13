# device-detector-go

[![GitHub issues](https://img.shields.io/github/issues/umutbasal/device-detector-go)](https://github.com/umutbasal/device-detector-go/issues)
[![GitHub stars](https://img.shields.io/github/stars/umutbasal/device-detector-go)](https://github.com/umutbasal/device-detector-go/stargazers)
[![GitHub license](https://img.shields.io/github/license/umutbasal/device-detector-go)](https://github.com/umutbasal/device-detector-go/blob/main/LICENSE)

Device-detector-go is a precise user agent parser and device detector written in Golang, backed by the largest and most up-to-date open-source user agent database and library written js.

Device-detector-go will parse any user agent and detect the browser, operating system, device used (desktop, tablet, mobile, tv, cars, console, etc.), brand and model. Works with go and in the browser.

    This library uses [device-detector-js](https://github.com/etienne-martin/device-detector-js) and originial library maintainer says;
    This library is heavily tested and relies on over 10,000 tests to detect thousands of user agent strings, even from rare and obscure browsers and devices.

This is a go port of [device-detector-js](https://github.com/etienne-martin/device-detector-js) (uses this library in v8 engine directly and exports logics) and relatively based on [Matomo device-detector](https://github.com/matomo-org/device-detector) rules

## Why this project

There is too many go user agent parsers, but none of them has powerful rules as [Matomo device-detector](https://github.com/matomo-org/device-detector) and there is only one golang port of matomo ([detector gamebtc/devicedetector](https://github.com/gamebtc/devicedetector)), but this one works heavy like 40ms. This is too much. This projects aim is using speed of the nodejs repo. And not implement the go version from scratch.

## Getting Started

### Installation

To use device-detector-go in your project, run:

```bash
go get github.com/umutbasal/device-detector-go
```

### Usage

#### GO import

```go
import (
 device detector "github.com/umutbasal/device-detector-go"
)
```

**Example** - user agent detection:

```go
import (
  devicedetector "github.com/umutbasal/device-detector-go"
)

func main() {
  userAgent := "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36"
  detector, err := devicedetector.NewDeviceDetector(devicedetector.DeviceDetectorOptions{})
  if err != nil {
    panic(err)
  }
  result, err := detector.Parse(user)
  if err != nil {
    panic(err)
  }

  fmt.Printf("%s\n", result)

}
```

Output:

```json
{
 "client": {
  "type": "browser",
  "name": "Chrome Mobile",
  "version": "41.0",
  "engine": "Blink",
  "engineVersion": ""
 },
 "os": {
  "name": "Android",
  "version": "6.0",
  "platform": ""
 },
 "device": {
  "type": "smartphone",
  "brand": "Google",
  "model": "Nexus 5X"
 },
 "bot": null
}
```

**Example** - bot detection:

```go
import (
 devicedetector "github.com/umutbasal/device-detector-go"
)

func main() {
 userAgent := "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36"
 detector, err := devicedetector.NewBotDetector(devicedetector.DeviceDetectorOptions{})
 if err != nil {
  panic(err)
 }
 result, err := detector.Parse(user)
 if err != nil {
  panic(err)
 }

 fmt.Printf("%s\n", result)
}
```

Output:

```json
{
 "name": "Googlebot",
 "category": "Search bot",
 "url": "http://www.google.com/bot.html",
 "producer": {
  "name": "Google Inc.",
  "url": "http://www.google.com"
 }
}
```

## What device-detector-go is able to detect

There is a list in original js [repo](https://github.com/etienne-martin/device-detector-js/#list-of-detected-operating-systems)

## Built with

- [Device-detector-js](https://github.com/etienne-martin/device-detector-js)
- [ESBuild](https://github.com/evanw/esbuild)
- [V8 Engine](https://github.com/rogchap/v8go)
- [Matomo device-detector](https://github.com/matomo-org/device-detector)

## Contributing

When contributing to this project, please first discuss the change you wish to make via issue, email, or any other method with the owners of this repository before making a change.

Update the [README.md](https://github.com/umutbasal/device-detector-go/blob/master/README.md) with details of changes to the library.

Execute `go test -v` and update the tests if needed.

## License

This is a free/libre library under license LGPL v3 or later.
