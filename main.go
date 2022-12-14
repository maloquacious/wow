/*
 * wars of warp - an implementation of warpwar
 *
 * Copyright (c) 2022 Michael D Henderson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

// Package main implements the engine for wars of warp.
package main

import (
	"github.com/mdhender/wow/cli"
	"log"
	"os"
	"time"
)

func main() {
	// default log format to UTC
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	//log.SetFlags(log.Lshortfile)

	defer func(started time.Time) {
		elapsed := time.Now().Sub(started)
		log.Printf("wow: total time %v\n", elapsed)
	}(time.Now())

	// print working directory
	if cwd, err := os.Getwd(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("wow: running in %s\n", cwd)
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	return cli.Execute()
}
