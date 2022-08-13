package devicedetector

import (
	"io/ioutil"
	"log"
	"strings"

	v8 "rogchap.com/v8go"
)

type Parser struct {
	options DeviceDetectorOptions
	ctx     *v8.Context
	Parse   func(userAgent string) (result interface{}, err error)
}

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

func NewBotDetector(options DeviceDetectorOptions) (parser Parser) {
	panic("not implemented")
}

func NewDeviceParser(options DeviceDetectorOptions) (parser Parser) {
	panic("not implemented")
}

func NewOperatingSystemParser() (parser Parser) {
	panic("not implemented")
}

func NewVendorFragmentParser() (parser Parser) {
	panic("not implemented")
}
