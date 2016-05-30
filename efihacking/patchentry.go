package main

import (
	"debug/pe"
	"encoding/binary"
	"fmt"
	"os"
)

func exitf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	path := os.Args[1]
	p, err := pe.Open(path)
	if err != nil {
		exitf("opening %s: %s", path, err)
	}

	var entrypoint uint32
	if oh, ok := p.OptionalHeader.(*pe.OptionalHeader64); !ok {
		exitf("%s is not a PE32+ binary", path)
	} else {
		entrypoint = oh.AddressOfEntryPoint
	}

	var fileOffset uint32
	for _, s := range p.Sections {
		if entrypoint < s.VirtualAddress || entrypoint > (s.VirtualAddress+s.VirtualSize) {
			fmt.Printf("Entrypoint 0x%x not in section %s\n", entrypoint, s.Name)
			continue
		}

		sectionOffset := entrypoint - s.VirtualAddress
		fileOffset = s.Offset + sectionOffset
		fmt.Printf("Entrypoint 0x%x at offset %d in section %s\n", entrypoint, sectionOffset, s.Name)
		fmt.Printf("Entrypoint at file offset %d\n", fileOffset)
		break
	}

	if fileOffset == 0 {
		exitf("Entrypoint not found in any section of %s", path)
	}
	f, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		exitf("Opening %s for patching: %s", path, err)
	}
	defer f.Close()

	var origInstr uint16
	if _, err = f.Seek(int64(fileOffset), 0); err != nil {
		exitf("Seeking in %s: %s", path, err)
	}
	if err = binary.Read(f, binary.LittleEndian, &origInstr); err != nil {
		exitf("Reading bytes before patch: %s", err)
	}

	if _, err = f.Seek(int64(fileOffset), 0); err != nil {
		exitf("Seeking in %s: %s", path, err)
	}
	if _, err = f.Write([]byte{0xeb, 0xfe}); err != nil {
		exitf("Patching %s: %s", path, err)
	}

	fmt.Printf("Entrypoint patched to loop for ever. Restore from GDB with:\n")
	fmt.Printf("set {unsigned int}$pc = 0x%x\n", origInstr)
}
