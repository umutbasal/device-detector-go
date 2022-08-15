package devicedetector

import (
	"io/ioutil"
	"log"
	"strings"

	v8 "rogchap.com/v8go"
)

// TODO: wont work with concurrent reqs design for concurrency

// Parser is a device detector parser.
type Parser struct {
	options DeviceDetectorOptions
	ctx     *v8.Context
	Parse   func(userAgent string) (result interface{}, err error)
}

// initContext initializes a v8 context from compiled device-detector-js library.
func initContext() (*v8.Context, error) {
	file, err := ioutil.ReadFile("build/dist/device-detector.js")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	src := string(file)
	src = strings.TrimPrefix(src, `"use strict";`+"\n")
	src = strings.TrimPrefix(src, "(() => {\n")
	src = strings.TrimSuffix(src, "\n})();\n")

	// Todo: take these from New_ functions nad run with native v8 apis. or export all funcs to global var with regexp replace.
	src = src + "\n" + `DeviceDetector = new require_src()`
	src = src + "\n" + `BotDetector = new require_bot()`
	// src = src + "\n" + `ClientParser_ = new init_client()`
	// src = src + "\n" + `DeviceParser = new init_device()`
	// src = src + "\n" + `OperatingSystemParser_ = new init_operating_system2()`
	// src = src + "\n" + `VendorFragmentParser_ = new init_vendor_fragment()`

	ctx := v8.NewContext()
	_, e := ctx.RunScript(src+"\n", "h.js")
	if e != nil {
		return nil, e
	}
	return ctx, nil
}

// TODO: clear duplicated code.
// TODO: use types.

// NewDeviceDetector creates a new device detector.
func NewDeviceDetector(options DeviceDetectorOptions) (parser Parser, err error) {
	ctx, err := initContext()
	if err != nil {
		return
	}
	class, _ := ctx.Global().Get("DeviceDetector")
	classTmp, _ := class.Object().AsFunction()
	instance, err := classTmp.NewInstance(v8.Undefined(v8.NewIsolate()))
	return Parser{
		options: options,
		ctx:     ctx,
		Parse: func(userAgent string) (result interface{}, err error) {
			ctx.Global().Set("userAgent", userAgent)
			parse, _ := instance.Object().Get("parse")
			res, _ := parse.AsFunction()
			ua, _ := ctx.Global().Get("userAgent")
			r, _ := res.Call(v8.Undefined(ctx.Isolate()), ua)
			return r.Object().MarshalJSON()
		},
	}, err
}

// NewBotDetector creates a new bot detector with options.
func NewBotDetector(options DeviceDetectorOptions) (parser Parser, err error) {
	ctx, err := initContext()
	if err != nil {
		return
	}
	class, _ := ctx.Global().Get("BotDetector")
	classTmp, _ := class.Object().AsFunction()
	instance, err := classTmp.NewInstance(v8.Undefined(v8.NewIsolate()))
	return Parser{
		options: options,
		ctx:     ctx,
		Parse: func(userAgent string) (result interface{}, err error) {
			ctx.Global().Set("userAgent", userAgent)
			parse, _ := instance.Object().Get("parse")
			res, _ := parse.AsFunction()
			ua, _ := ctx.Global().Get("userAgent")
			r, _ := res.Call(v8.Undefined(ctx.Isolate()), ua)
			return r.Object().MarshalJSON()
		},
	}, err
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
