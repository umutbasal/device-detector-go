package devicedetector

import (
	"io/ioutil"
	"log"
	"strings"

	v8 "rogchap.com/v8go"
)

// Parser is a device detector parser.
type Parser struct {
	options DeviceDetectorOptions
	ctx     *v8.Context
	Parse   func(userAgent string) (result interface{}, err error)
}

// initContext initializes a v8 context from compiled device-detector-js library.
func initContext(source string) (*v8.Context, error) {
	file, err := ioutil.ReadFile("build/dist/device-detector.js")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	src := string(file)
	src = strings.TrimPrefix(src, `"use strict";`+"\n")
	src = strings.TrimPrefix(src, "(() => {\n")
	src = strings.TrimSuffix(src, "\n})();\n")
	src = src + "\n" + `DeviceDetector = new require_src()`

	ctx := v8.NewContext()
	_, e := ctx.RunScript(src+"\n", "h.js")
	if e != nil {
		return nil, e
	}
	if source != "" {
		_, e := ctx.RunScript(source, "h.js")
		if e != nil {
			return nil, e
		}
	}
	return ctx, nil
}

// NewDeviceDetector creates a new device detector.
func NewDeviceDetector(options DeviceDetectorOptions) (parser Parser, err error) {
	ctx, err := initContext(`let d = new DeviceDetector();`)
	return Parser{
		options: options,
		ctx:     ctx,
		Parse: func(userAgent string) (result interface{}, err error) {
			script := "\n" + `JSON.stringify(d.parse("` + userAgent + `"))`
			r, e := ctx.RunScript(script, "h.js")
			if e != nil {
				return nil, e
			}
			return r.String(), nil
		},
	}, err
}

// NewBotDetector creates a new bot detector with options.
func NewBotDetector(options DeviceDetectorOptions) (parser Parser) {
	panic("not implemented")
}

// NewDeviceParser creates a new device parser with options
func NewDeviceParser(options DeviceDetectorOptions) (parser Parser) {
	panic("not implemented")
}

// NewOperatingSystemParser creates a new operating system parser.
func NewOperatingSystemParser() (parser Parser) {
	panic("not implemented")
}

// NewVendorFragmentParser creates a new vendor fragment parser.
func NewVendorFragmentParser() (parser Parser) {
	panic("not implemented")
}
