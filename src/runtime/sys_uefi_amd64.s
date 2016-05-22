#include "textflag.h"

TEXT _rt0_amd64_uefi(SB),NOSPLIT,$-8
    // Print the OK text to prove we made it here.
    MOVQ 64(DX), CX // ConOut
    MOVQ 8(CX), BX
    MOVQ $oktext<>(SB), DX
    CALL BX

hang:
    JMP hang

  	LEAQ	8(SP), SI // argv
	MOVQ	0(SP), DI // argc
	MOVQ	$main(SB), AX
	JMP	AX

// func now() (sec int64, nsec int32)
TEXT time·now(SB),NOSPLIT,$8-12
	CALL	runtime·nanotime(SB)
	MOVQ	0(SP), AX

	// generated code for
	//	func f(x uint64) (uint64, uint64) { return x/1000000000, x%100000000 }
	// adapted to reduce duplication
	MOVQ	AX, CX
	MOVQ	$1360296554856532783, AX
	MULQ	CX
	ADDQ	CX, DX
	RCRQ	$1, DX
	SHRQ	$29, DX
	MOVQ	DX, sec+0(FP)
	IMULQ	$1000000000, DX
	SUBQ	DX, CX
	MOVL	CX, nsec+8(FP)
	RET

TEXT main(SB),NOSPLIT,$-8
	MOVQ	$runtime·rt0_go(SB), AX
	JMP	AX

TEXT runtime·settls(SB),NOSPLIT,$0
    // Need to set FS segment config correctly
	RET

TEXT runtime·exit(SB),NOSPLIT,$0-4
	RET

TEXT runtime·nanotime(SB),NOSPLIT,$0-8
    // TODO, time is currently immemorial.
	MOVQ	$0, ret+0(FP)
	RET

TEXT runtime·write(SB),NOSPLIT,$0-20
 	MOVQ	$0, CX
 	RET

TEXT runtime·usleep(SB),NOSPLIT,$0
 	RET

DATA oktext<>+0x0(SB)/2, $0x004f
DATA oktext<>+0x2(SB)/2, $0x004b
DATA oktext<>+0x4(SB)/2, $0
GLOBL oktext<>(SB),NOPTR,$6
