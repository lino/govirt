package govirt

// #cgo pkg-config: libvirt
// #include <stdlib.h>
// #include <libvirt/libvirt.h>
import "C"
import "unsafe"
import "errors"

// Represents a libvirt node.
// Exposes methods for controlling the hypervisor.
type Node struct {
	Url  string
	conn Connection
}

type Connection struct {
	Pointer C.virConnectPtr
}

type Domain struct {
}

type Network struct {
}

type StorageVol struct {
}

type StoragePool struct {
}

func (p *Node) Connect() error {
	cUrl := C.CString(p.Url)
	defer C.free(unsafe.Pointer(cUrl))
	conn := C.virConnectOpenAuth(cUrl, C.virConnectAuthPtrDefault, 0) // TODO: Implement alias override
	if conn == nil {
		return errors.New("Unable to establish connection")
	}
	p.conn = Connection{conn}

	return nil
}

func (p *Node) Disconnect() {
	if p.conn.Pointer != nil {
		C.virConnectClose(p.conn.Pointer)
	}
}

func Connect(url string) (Node, error) {
	var retval Node
	retval = Node{Url: url}
	err := retval.Connect()
	return retval, err
}
