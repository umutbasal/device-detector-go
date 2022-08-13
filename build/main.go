package main

import (
	_ "embed"
	"log"

	esbuild "github.com/evanw/esbuild/pkg/api"
)

func main() {
	// ! NO INSTALLATION IS REQUIRED FOR NOW !
	// run npm install in the project directory
	// exec := exec.Command("npm", "install", "--legacy-peer-deps")
	// exec.Dir = "./device-detector-js"
	// exec.Stdout = os.Stdout
	// exec.Stderr = os.Stderr
	// if err := exec.Run(); err != nil {
	// 	log.Fatal(err)
	// }
	// run esbuild in the project directory
	bundleContent := esbuild.Build(esbuild.BuildOptions{
		EntryPoints: []string{"device-detector-js/src/index.ts"},
		Outfile:     "dist/device-detector.js",
		Bundle:      true,
		Write:       true,
		Target:      esbuild.ES2015,
		Format:      esbuild.FormatIIFE,
	})
	if len(bundleContent.OutputFiles) == 0 {
		log.Fatalf("no output files: %v", bundleContent.Errors)
	}
}
