// +build darwin freebsd netbsd openbsd windows

package main

import (
	"fmt"
)

func watch(site *Site) {
	fmt.Printf("Listening for changes to %s is NOT supported on this platform\n", site.Src)
}
