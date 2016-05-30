// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"unsafe"
)

type sigset struct{}

type mOS struct{}

//go:nosplit
func getRandomData(r []byte) {
	// TODO: use UEFI RNG Protocol to obtain random bits
}

//go:nosplit
func semacreate(mp *m) {
	// TODO: figure out what sems mean in this whacky universe
}

//go:nosplit
func semasleep(ns int64) int32 {
	// TODO: figure out what sems mean in this whacky universe
	return 0 // BIG FAT LIE
}

//go:nosplit
func semawakeup(mp *m) {
	// TODO: figure out what sems mean in this whacky universe
}

//go:nosplit
func sysAlloc(n uintptr, sysStat *uint64) unsafe.Pointer {
	// Allocate memory from UEFI Memory protocol.
	return nil
}

//go:nosplit
func sysFree(v unsafe.Pointer, n uintptr, sysStat *uint64) {
	// Return memory to UEFI Memory protocol.
}

func sysReserve(v unsafe.Pointer, n uintptr, reserved *bool) unsafe.Pointer {
	// ???
	return nil
}

func sysMap(v unsafe.Pointer, n uintptr, reserved bool, sysStat *uint64) {
	// ???
}

func sysFault(v unsafe.Pointer, n uintptr) {
	// ???
}

func sysUsed(v unsafe.Pointer, n uintptr) {
	// ???
}

func sysUnused(v unsafe.Pointer, n uintptr) {
	// ???
}

func sigpanic() {
	// TODO: AFAICT, UEFI will never get here, there are no signals in
	// the UEFI environment and processor faults will just panic the
	// universe without telling us.
	throw("fault")
}

func gogetenv(s string) string {
	// TODO: maybe look up the key in UEFI nvram? Probably not.
	return ""
}

func osyield() {
	// Do nothing, UEFI has no multiprocessing.
}

func signame(sig uint32) string {
	return ""
}

func crash() {
	// Try to return to the UEFI environment somehow. For now, ignore.
}

var _cgo_setenv unsafe.Pointer   // pointer to C function
var _cgo_unsetenv unsafe.Pointer // pointer to C function

//go:nosplit
func msigsave(mp *m) {}

//go:nosplit
func msigrestore(sigmask sigset) {}

//go:nosplit
func sigblock() {}

func goenvs() {}

func mpreinit(mp *m) {
}

func minit() {}

func initsig(preinit bool) {}

func unminit() {}

// May run with m.p==nil, so write barriers are not allowed.
//go:nowritebarrier
func newosproc(mp *m, stk unsafe.Pointer) {
	// This is probably a no-op under UEFI, because everything is
	// single-threaded and polley.
	throw("todo")
}

func resetcpuprofiler(hz int32) {
	// TODO: Enable profiling interrupts.
	getg().m.profilehz = hz
}

const _NSIG = 0

func sigenable(sig uint32) {
}

func sigdisable(sig uint32) {
}

func sigignore(sig uint32) {
}

func osinit() {
	ncpu = 1
}

//go:nosplit
func ok()

//go:cgo_export_static efi_main
func efi_main()
