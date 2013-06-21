package govirt

// #cgo pkg-config: libvirt
// #include <libvirt/libvirt.h>
import "C"

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("nothing to see here yet")

	os.Exit(0)

}
