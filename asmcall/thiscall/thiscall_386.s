TEXT ·Call0(SB),4,$0-12
  MOVL a+4(FP), CX

  MOVL addr+0(FP), AX
  CALL AX
  MOVL AX, ret+8(FP)
  RET

TEXT ·Call1(SB),4,$4-16
  MOVL a+4(FP), CX
  MOVL b+8(FP), BX
  MOVL BX, 0(SP)

  MOVL addr+0(FP), AX
  CALL AX
  SUBL $4, SP
  MOVL AX, ret+12(FP)
  RET

TEXT ·Call2(SB),4,$8-20
  MOVL a+4(FP), CX
  MOVL b+8(FP), BX
  MOVL BX, 0(SP)
  MOVL c+12(FP), BX
  MOVL BX, 4(SP)

  MOVL addr+0(FP), AX
  CALL AX
  SUBL $8, SP
  MOVL AX, ret+16(FP)
  RET

TEXT ·Call3(SB),4,$12-24
  MOVL a+4(FP), CX
  MOVL b+8(FP), BX
  MOVL BX, 0(SP)
  MOVL c+12(FP), BX
  MOVL BX, 4(SP)
  MOVL d+16(FP), BX
  MOVL BX, 8(SP)

  MOVL addr+0(FP), AX
  CALL AX
  SUBL $12, SP
  MOVL AX, ret+20(FP)
  RET

TEXT ·Call4(SB),4,$16-28
  MOVL a+4(FP), CX
  MOVL b+8(FP), BX
  MOVL BX, 0(SP)
  MOVL c+12(FP), BX
  MOVL BX, 4(SP)
  MOVL d+16(FP), BX
  MOVL BX, 8(SP)
  MOVL e+20(FP), BX
  MOVL BX, 12(SP)

  MOVL addr+0(FP), AX
  CALL AX
  SUBL $16, SP
  MOVL AX, ret+24(FP)
  RET

TEXT ·Call5(SB),4,$20-32
  MOVL a+4(FP), CX
  MOVL b+8(FP), BX
  MOVL BX, 0(SP)
  MOVL c+12(FP), BX
  MOVL BX, 4(SP)
  MOVL d+16(FP), BX
  MOVL BX, 8(SP)
  MOVL e+20(FP), BX
  MOVL BX, 12(SP)
  MOVL f+24(FP), BX
  MOVL BX, 16(SP)

  MOVL addr+0(FP), AX
  CALL AX
  SUBL $20, SP
  MOVL AX, ret+28(FP)
  RET
