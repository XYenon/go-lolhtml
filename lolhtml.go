// Package lolhtml provides the ability to parse and rewrite HTML on the fly,
// with a CSS-selector based API.
//
// It is a binding for the Rust crate lol_html.
// https://github.com/cloudflare/lol-html
//
// Please see /examples subdirectory for more detailed examples.
package lolhtml

/*
#cgo CFLAGS:-I${SRCDIR}/build/include
#cgo LDFLAGS:-llolhtml
#cgo !windows LDFLAGS:-lm
#cgo linux,amd64 LDFLAGS:-L${SRCDIR}/build/linux-x86_64
#cgo linux,arm64 LDFLAGS:-L${SRCDIR}/build/linux-aarch64
#cgo darwin,amd64 LDFLAGS:-L${SRCDIR}/build/macos-x86_64
#cgo darwin,arm64 LDFLAGS:-L${SRCDIR}/build/macos-aarch64
#include <stdlib.h>
#include "lol_html.h"
*/
import "C"
