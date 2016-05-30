#include "textflag.h"

#define HANG CALL runtime·iamhang(SB)

TEXT _rt0_amd64_uefi(SB),NOSPLIT,$-8
    MOVQ CX, _efi_image_handle(SB)
    MOVQ DX, _efi_services(SB)

    HANG

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
    MOVQ    $_efi_services(SB), BX
	MOVQ	88(BX), BX // BootServices
    MOVQ    24+192(BX), AX // Exit
    
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

GLOBL _efi_image_handle(SB),NOPTR,$8
GLOBL _efi_services(SB),NOPTR,$8

DATA oktext<>+0x00(SB)/4, $0x004b004f
DATA oktext<>+0x04(SB)/4, $0x00000000
GLOBL oktext<>(SB),NOPTR,$8

TEXT runtime·iamhere(SB),NOSPLIT,$0
    MOVQ _efi_services(SB), AX
    MOVQ 64(AX), CX // ConOut
    MOVQ 8(CX), AX
    MOVQ $oktext<>(SB), DX
    CALL AX
    RET

TEXT runtime·iamhang(SB),NOSPLIT,$0
    CALL runtime·iamhere(SB)
    CALL runtime·iamhere(SB)
j:
    JMP j
