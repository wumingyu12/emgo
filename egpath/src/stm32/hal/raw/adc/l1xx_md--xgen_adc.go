// +build l1xx_md

package adc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type ADC_Periph struct {
	SR    SR
	CR1   CR1
	CR2   CR2
	SMPR1 SMPR1
	SMPR2 SMPR2
	SMPR3 SMPR3
	JOFR1 JOFR1
	JOFR2 JOFR2
	JOFR3 JOFR3
	JOFR4 JOFR4
	HTR   HTR
	LTR   LTR
	SQR1  SQR1
	SQR2  SQR2
	SQR3  SQR3
	SQR4  SQR4
	SQR5  SQR5
	JSQR  JSQR
	JDR1  JDR1
	JDR2  JDR2
	JDR3  JDR3
	JDR4  JDR4
	DR    DR
	SMPR0 SMPR0
}

func (p *ADC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var ADC1 = (*ADC_Periph)(unsafe.Pointer(uintptr(mmap.ADC1_BASE)))

type SR_Bits uint32

func (b SR_Bits) Field(mask SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR_Bits) J(v int) SR_Bits {
	return SR_Bits(bits.Make32(v, uint32(mask)))
}

type SR struct{ mmio.U32 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U32.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U32.Store(uint32(b)) }

type SR_Mask struct{ mmio.UM32 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM32.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) AWD() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(AWD)}}
}

func (p *ADC_Periph) EOC() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(EOC)}}
}

func (p *ADC_Periph) JEOC() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(JEOC)}}
}

func (p *ADC_Periph) JSTRT() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(JSTRT)}}
}

func (p *ADC_Periph) STRT() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(STRT)}}
}

func (p *ADC_Periph) OVR() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(OVR)}}
}

func (p *ADC_Periph) ADONS() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(ADONS)}}
}

func (p *ADC_Periph) RCNR() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(RCNR)}}
}

func (p *ADC_Periph) JCNR() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(JCNR)}}
}

type CR1_Bits uint32

func (b CR1_Bits) Field(mask CR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1_Bits) J(v int) CR1_Bits {
	return CR1_Bits(bits.Make32(v, uint32(mask)))
}

type CR1 struct{ mmio.U32 }

func (r *CR1) Bits(mask CR1_Bits) CR1_Bits { return CR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR1) StoreBits(mask, b CR1_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR1) SetBits(mask CR1_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CR1) ClearBits(mask CR1_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CR1) Load() CR1_Bits              { return CR1_Bits(r.U32.Load()) }
func (r *CR1) Store(b CR1_Bits)            { r.U32.Store(uint32(b)) }

type CR1_Mask struct{ mmio.UM32 }

func (rm CR1_Mask) Load() CR1_Bits   { return CR1_Bits(rm.UM32.Load()) }
func (rm CR1_Mask) Store(b CR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) AWDCH() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(AWDCH)}}
}

func (p *ADC_Periph) EOCIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(EOCIE)}}
}

func (p *ADC_Periph) AWDIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(AWDIE)}}
}

func (p *ADC_Periph) JEOCIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(JEOCIE)}}
}

func (p *ADC_Periph) SCAN() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(SCAN)}}
}

func (p *ADC_Periph) AWDSGL() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(AWDSGL)}}
}

func (p *ADC_Periph) JAUTO() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(JAUTO)}}
}

func (p *ADC_Periph) DISCEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(DISCEN)}}
}

func (p *ADC_Periph) JDISCEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(JDISCEN)}}
}

func (p *ADC_Periph) DISCNUM() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(DISCNUM)}}
}

func (p *ADC_Periph) PDD() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(PDD)}}
}

func (p *ADC_Periph) PDI() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(PDI)}}
}

func (p *ADC_Periph) JAWDEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(JAWDEN)}}
}

func (p *ADC_Periph) AWDEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(AWDEN)}}
}

func (p *ADC_Periph) RES() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(RES)}}
}

func (p *ADC_Periph) OVRIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OVRIE)}}
}

type CR2_Bits uint32

func (b CR2_Bits) Field(mask CR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2_Bits) J(v int) CR2_Bits {
	return CR2_Bits(bits.Make32(v, uint32(mask)))
}

type CR2 struct{ mmio.U32 }

func (r *CR2) Bits(mask CR2_Bits) CR2_Bits { return CR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR2) StoreBits(mask, b CR2_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR2) SetBits(mask CR2_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CR2) ClearBits(mask CR2_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CR2) Load() CR2_Bits              { return CR2_Bits(r.U32.Load()) }
func (r *CR2) Store(b CR2_Bits)            { r.U32.Store(uint32(b)) }

type CR2_Mask struct{ mmio.UM32 }

func (rm CR2_Mask) Load() CR2_Bits   { return CR2_Bits(rm.UM32.Load()) }
func (rm CR2_Mask) Store(b CR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) ADON() CR2_Mask {
	return CR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(ADON)}}
}

func (p *ADC_Periph) CONT() CR2_Mask {
	return CR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CONT)}}
}

func (p *ADC_Periph) CFG() CR2_Mask {
	return CR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CFG)}}
}

func (p *ADC_Periph) DELS() CR2_Mask {
	return CR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(DELS)}}
}

func (p *ADC_Periph) DMA() CR2_Mask {
	return CR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(DMA)}}
}

func (p *ADC_Periph) DDS() CR2_Mask {
	return CR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(DDS)}}
}

func (p *ADC_Periph) EOCS() CR2_Mask {
	return CR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(EOCS)}}
}

func (p *ADC_Periph) ALIGN() CR2_Mask {
	return CR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(ALIGN)}}
}

func (p *ADC_Periph) JEXTSEL() CR2_Mask {
	return CR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(JEXTSEL)}}
}

func (p *ADC_Periph) JEXTEN() CR2_Mask {
	return CR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(JEXTEN)}}
}

func (p *ADC_Periph) JSWSTART() CR2_Mask {
	return CR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(JSWSTART)}}
}

func (p *ADC_Periph) EXTSEL() CR2_Mask {
	return CR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(EXTSEL)}}
}

func (p *ADC_Periph) EXTEN() CR2_Mask {
	return CR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(EXTEN)}}
}

func (p *ADC_Periph) SWSTART() CR2_Mask {
	return CR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(SWSTART)}}
}

type SMPR1_Bits uint32

func (b SMPR1_Bits) Field(mask SMPR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SMPR1_Bits) J(v int) SMPR1_Bits {
	return SMPR1_Bits(bits.Make32(v, uint32(mask)))
}

type SMPR1 struct{ mmio.U32 }

func (r *SMPR1) Bits(mask SMPR1_Bits) SMPR1_Bits { return SMPR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *SMPR1) StoreBits(mask, b SMPR1_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SMPR1) SetBits(mask SMPR1_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *SMPR1) ClearBits(mask SMPR1_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *SMPR1) Load() SMPR1_Bits                { return SMPR1_Bits(r.U32.Load()) }
func (r *SMPR1) Store(b SMPR1_Bits)              { r.U32.Store(uint32(b)) }

type SMPR1_Mask struct{ mmio.UM32 }

func (rm SMPR1_Mask) Load() SMPR1_Bits   { return SMPR1_Bits(rm.UM32.Load()) }
func (rm SMPR1_Mask) Store(b SMPR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SMP20() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(SMP20)}}
}

func (p *ADC_Periph) SMP21() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(SMP21)}}
}

func (p *ADC_Periph) SMP22() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(SMP22)}}
}

func (p *ADC_Periph) SMP23() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(SMP23)}}
}

func (p *ADC_Periph) SMP24() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(SMP24)}}
}

func (p *ADC_Periph) SMP25() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(SMP25)}}
}

func (p *ADC_Periph) SMP26() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(SMP26)}}
}

func (p *ADC_Periph) SMP27() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(SMP27)}}
}

func (p *ADC_Periph) SMP28() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(SMP28)}}
}

func (p *ADC_Periph) SMP29() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(SMP29)}}
}

type SMPR2_Bits uint32

func (b SMPR2_Bits) Field(mask SMPR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SMPR2_Bits) J(v int) SMPR2_Bits {
	return SMPR2_Bits(bits.Make32(v, uint32(mask)))
}

type SMPR2 struct{ mmio.U32 }

func (r *SMPR2) Bits(mask SMPR2_Bits) SMPR2_Bits { return SMPR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *SMPR2) StoreBits(mask, b SMPR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SMPR2) SetBits(mask SMPR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *SMPR2) ClearBits(mask SMPR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *SMPR2) Load() SMPR2_Bits                { return SMPR2_Bits(r.U32.Load()) }
func (r *SMPR2) Store(b SMPR2_Bits)              { r.U32.Store(uint32(b)) }

type SMPR2_Mask struct{ mmio.UM32 }

func (rm SMPR2_Mask) Load() SMPR2_Bits   { return SMPR2_Bits(rm.UM32.Load()) }
func (rm SMPR2_Mask) Store(b SMPR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SMP10() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(SMP10)}}
}

func (p *ADC_Periph) SMP11() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(SMP11)}}
}

func (p *ADC_Periph) SMP12() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(SMP12)}}
}

func (p *ADC_Periph) SMP13() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(SMP13)}}
}

func (p *ADC_Periph) SMP14() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(SMP14)}}
}

func (p *ADC_Periph) SMP15() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(SMP15)}}
}

func (p *ADC_Periph) SMP16() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(SMP16)}}
}

func (p *ADC_Periph) SMP17() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(SMP17)}}
}

func (p *ADC_Periph) SMP18() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(SMP18)}}
}

func (p *ADC_Periph) SMP19() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(SMP19)}}
}

type SMPR3_Bits uint32

func (b SMPR3_Bits) Field(mask SMPR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SMPR3_Bits) J(v int) SMPR3_Bits {
	return SMPR3_Bits(bits.Make32(v, uint32(mask)))
}

type SMPR3 struct{ mmio.U32 }

func (r *SMPR3) Bits(mask SMPR3_Bits) SMPR3_Bits { return SMPR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *SMPR3) StoreBits(mask, b SMPR3_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SMPR3) SetBits(mask SMPR3_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *SMPR3) ClearBits(mask SMPR3_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *SMPR3) Load() SMPR3_Bits                { return SMPR3_Bits(r.U32.Load()) }
func (r *SMPR3) Store(b SMPR3_Bits)              { r.U32.Store(uint32(b)) }

type SMPR3_Mask struct{ mmio.UM32 }

func (rm SMPR3_Mask) Load() SMPR3_Bits   { return SMPR3_Bits(rm.UM32.Load()) }
func (rm SMPR3_Mask) Store(b SMPR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SMP0() SMPR3_Mask {
	return SMPR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(SMP0)}}
}

func (p *ADC_Periph) SMP1() SMPR3_Mask {
	return SMPR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(SMP1)}}
}

func (p *ADC_Periph) SMP2() SMPR3_Mask {
	return SMPR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(SMP2)}}
}

func (p *ADC_Periph) SMP3() SMPR3_Mask {
	return SMPR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(SMP3)}}
}

func (p *ADC_Periph) SMP4() SMPR3_Mask {
	return SMPR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(SMP4)}}
}

func (p *ADC_Periph) SMP5() SMPR3_Mask {
	return SMPR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(SMP5)}}
}

func (p *ADC_Periph) SMP6() SMPR3_Mask {
	return SMPR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(SMP6)}}
}

func (p *ADC_Periph) SMP7() SMPR3_Mask {
	return SMPR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(SMP7)}}
}

func (p *ADC_Periph) SMP8() SMPR3_Mask {
	return SMPR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(SMP8)}}
}

func (p *ADC_Periph) SMP9() SMPR3_Mask {
	return SMPR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(SMP9)}}
}

type JOFR1_Bits uint32

func (b JOFR1_Bits) Field(mask JOFR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JOFR1_Bits) J(v int) JOFR1_Bits {
	return JOFR1_Bits(bits.Make32(v, uint32(mask)))
}

type JOFR1 struct{ mmio.U32 }

func (r *JOFR1) Bits(mask JOFR1_Bits) JOFR1_Bits { return JOFR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *JOFR1) StoreBits(mask, b JOFR1_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JOFR1) SetBits(mask JOFR1_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *JOFR1) ClearBits(mask JOFR1_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *JOFR1) Load() JOFR1_Bits                { return JOFR1_Bits(r.U32.Load()) }
func (r *JOFR1) Store(b JOFR1_Bits)              { r.U32.Store(uint32(b)) }

type JOFR1_Mask struct{ mmio.UM32 }

func (rm JOFR1_Mask) Load() JOFR1_Bits   { return JOFR1_Bits(rm.UM32.Load()) }
func (rm JOFR1_Mask) Store(b JOFR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JOFFSET1() JOFR1_Mask {
	return JOFR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(JOFFSET1)}}
}

type JOFR2_Bits uint32

func (b JOFR2_Bits) Field(mask JOFR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JOFR2_Bits) J(v int) JOFR2_Bits {
	return JOFR2_Bits(bits.Make32(v, uint32(mask)))
}

type JOFR2 struct{ mmio.U32 }

func (r *JOFR2) Bits(mask JOFR2_Bits) JOFR2_Bits { return JOFR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *JOFR2) StoreBits(mask, b JOFR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JOFR2) SetBits(mask JOFR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *JOFR2) ClearBits(mask JOFR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *JOFR2) Load() JOFR2_Bits                { return JOFR2_Bits(r.U32.Load()) }
func (r *JOFR2) Store(b JOFR2_Bits)              { r.U32.Store(uint32(b)) }

type JOFR2_Mask struct{ mmio.UM32 }

func (rm JOFR2_Mask) Load() JOFR2_Bits   { return JOFR2_Bits(rm.UM32.Load()) }
func (rm JOFR2_Mask) Store(b JOFR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JOFFSET2() JOFR2_Mask {
	return JOFR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(JOFFSET2)}}
}

type JOFR3_Bits uint32

func (b JOFR3_Bits) Field(mask JOFR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JOFR3_Bits) J(v int) JOFR3_Bits {
	return JOFR3_Bits(bits.Make32(v, uint32(mask)))
}

type JOFR3 struct{ mmio.U32 }

func (r *JOFR3) Bits(mask JOFR3_Bits) JOFR3_Bits { return JOFR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *JOFR3) StoreBits(mask, b JOFR3_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JOFR3) SetBits(mask JOFR3_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *JOFR3) ClearBits(mask JOFR3_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *JOFR3) Load() JOFR3_Bits                { return JOFR3_Bits(r.U32.Load()) }
func (r *JOFR3) Store(b JOFR3_Bits)              { r.U32.Store(uint32(b)) }

type JOFR3_Mask struct{ mmio.UM32 }

func (rm JOFR3_Mask) Load() JOFR3_Bits   { return JOFR3_Bits(rm.UM32.Load()) }
func (rm JOFR3_Mask) Store(b JOFR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JOFFSET3() JOFR3_Mask {
	return JOFR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 32)), uint32(JOFFSET3)}}
}

type JOFR4_Bits uint32

func (b JOFR4_Bits) Field(mask JOFR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JOFR4_Bits) J(v int) JOFR4_Bits {
	return JOFR4_Bits(bits.Make32(v, uint32(mask)))
}

type JOFR4 struct{ mmio.U32 }

func (r *JOFR4) Bits(mask JOFR4_Bits) JOFR4_Bits { return JOFR4_Bits(r.U32.Bits(uint32(mask))) }
func (r *JOFR4) StoreBits(mask, b JOFR4_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JOFR4) SetBits(mask JOFR4_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *JOFR4) ClearBits(mask JOFR4_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *JOFR4) Load() JOFR4_Bits                { return JOFR4_Bits(r.U32.Load()) }
func (r *JOFR4) Store(b JOFR4_Bits)              { r.U32.Store(uint32(b)) }

type JOFR4_Mask struct{ mmio.UM32 }

func (rm JOFR4_Mask) Load() JOFR4_Bits   { return JOFR4_Bits(rm.UM32.Load()) }
func (rm JOFR4_Mask) Store(b JOFR4_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JOFFSET4() JOFR4_Mask {
	return JOFR4_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(JOFFSET4)}}
}

type HTR_Bits uint32

func (b HTR_Bits) Field(mask HTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HTR_Bits) J(v int) HTR_Bits {
	return HTR_Bits(bits.Make32(v, uint32(mask)))
}

type HTR struct{ mmio.U32 }

func (r *HTR) Bits(mask HTR_Bits) HTR_Bits { return HTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *HTR) StoreBits(mask, b HTR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HTR) SetBits(mask HTR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *HTR) ClearBits(mask HTR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *HTR) Load() HTR_Bits              { return HTR_Bits(r.U32.Load()) }
func (r *HTR) Store(b HTR_Bits)            { r.U32.Store(uint32(b)) }

type HTR_Mask struct{ mmio.UM32 }

func (rm HTR_Mask) Load() HTR_Bits   { return HTR_Bits(rm.UM32.Load()) }
func (rm HTR_Mask) Store(b HTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) HT() HTR_Mask {
	return HTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 40)), uint32(HT)}}
}

type LTR_Bits uint32

func (b LTR_Bits) Field(mask LTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LTR_Bits) J(v int) LTR_Bits {
	return LTR_Bits(bits.Make32(v, uint32(mask)))
}

type LTR struct{ mmio.U32 }

func (r *LTR) Bits(mask LTR_Bits) LTR_Bits { return LTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *LTR) StoreBits(mask, b LTR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *LTR) SetBits(mask LTR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *LTR) ClearBits(mask LTR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *LTR) Load() LTR_Bits              { return LTR_Bits(r.U32.Load()) }
func (r *LTR) Store(b LTR_Bits)            { r.U32.Store(uint32(b)) }

type LTR_Mask struct{ mmio.UM32 }

func (rm LTR_Mask) Load() LTR_Bits   { return LTR_Bits(rm.UM32.Load()) }
func (rm LTR_Mask) Store(b LTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) LT() LTR_Mask {
	return LTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 44)), uint32(LT)}}
}

type SQR1_Bits uint32

func (b SQR1_Bits) Field(mask SQR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR1_Bits) J(v int) SQR1_Bits {
	return SQR1_Bits(bits.Make32(v, uint32(mask)))
}

type SQR1 struct{ mmio.U32 }

func (r *SQR1) Bits(mask SQR1_Bits) SQR1_Bits { return SQR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *SQR1) StoreBits(mask, b SQR1_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SQR1) SetBits(mask SQR1_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *SQR1) ClearBits(mask SQR1_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *SQR1) Load() SQR1_Bits               { return SQR1_Bits(r.U32.Load()) }
func (r *SQR1) Store(b SQR1_Bits)             { r.U32.Store(uint32(b)) }

type SQR1_Mask struct{ mmio.UM32 }

func (rm SQR1_Mask) Load() SQR1_Bits   { return SQR1_Bits(rm.UM32.Load()) }
func (rm SQR1_Mask) Store(b SQR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) L() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint32(L)}}
}

func (p *ADC_Periph) SQ28() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint32(SQ28)}}
}

func (p *ADC_Periph) SQ27() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint32(SQ27)}}
}

func (p *ADC_Periph) SQ26() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint32(SQ26)}}
}

func (p *ADC_Periph) SQ25() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint32(SQ25)}}
}

type SQR2_Bits uint32

func (b SQR2_Bits) Field(mask SQR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR2_Bits) J(v int) SQR2_Bits {
	return SQR2_Bits(bits.Make32(v, uint32(mask)))
}

type SQR2 struct{ mmio.U32 }

func (r *SQR2) Bits(mask SQR2_Bits) SQR2_Bits { return SQR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *SQR2) StoreBits(mask, b SQR2_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SQR2) SetBits(mask SQR2_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *SQR2) ClearBits(mask SQR2_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *SQR2) Load() SQR2_Bits               { return SQR2_Bits(r.U32.Load()) }
func (r *SQR2) Store(b SQR2_Bits)             { r.U32.Store(uint32(b)) }

type SQR2_Mask struct{ mmio.UM32 }

func (rm SQR2_Mask) Load() SQR2_Bits   { return SQR2_Bits(rm.UM32.Load()) }
func (rm SQR2_Mask) Store(b SQR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ19() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(SQ19)}}
}

func (p *ADC_Periph) SQ20() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(SQ20)}}
}

func (p *ADC_Periph) SQ21() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(SQ21)}}
}

func (p *ADC_Periph) SQ22() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(SQ22)}}
}

func (p *ADC_Periph) SQ23() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(SQ23)}}
}

func (p *ADC_Periph) SQ24() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(SQ24)}}
}

type SQR3_Bits uint32

func (b SQR3_Bits) Field(mask SQR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR3_Bits) J(v int) SQR3_Bits {
	return SQR3_Bits(bits.Make32(v, uint32(mask)))
}

type SQR3 struct{ mmio.U32 }

func (r *SQR3) Bits(mask SQR3_Bits) SQR3_Bits { return SQR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *SQR3) StoreBits(mask, b SQR3_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SQR3) SetBits(mask SQR3_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *SQR3) ClearBits(mask SQR3_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *SQR3) Load() SQR3_Bits               { return SQR3_Bits(r.U32.Load()) }
func (r *SQR3) Store(b SQR3_Bits)             { r.U32.Store(uint32(b)) }

type SQR3_Mask struct{ mmio.UM32 }

func (rm SQR3_Mask) Load() SQR3_Bits   { return SQR3_Bits(rm.UM32.Load()) }
func (rm SQR3_Mask) Store(b SQR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ13() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(SQ13)}}
}

func (p *ADC_Periph) SQ14() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(SQ14)}}
}

func (p *ADC_Periph) SQ15() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(SQ15)}}
}

func (p *ADC_Periph) SQ16() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(SQ16)}}
}

func (p *ADC_Periph) SQ17() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(SQ17)}}
}

func (p *ADC_Periph) SQ18() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(SQ18)}}
}

type SQR4_Bits uint32

func (b SQR4_Bits) Field(mask SQR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR4_Bits) J(v int) SQR4_Bits {
	return SQR4_Bits(bits.Make32(v, uint32(mask)))
}

type SQR4 struct{ mmio.U32 }

func (r *SQR4) Bits(mask SQR4_Bits) SQR4_Bits { return SQR4_Bits(r.U32.Bits(uint32(mask))) }
func (r *SQR4) StoreBits(mask, b SQR4_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SQR4) SetBits(mask SQR4_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *SQR4) ClearBits(mask SQR4_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *SQR4) Load() SQR4_Bits               { return SQR4_Bits(r.U32.Load()) }
func (r *SQR4) Store(b SQR4_Bits)             { r.U32.Store(uint32(b)) }

type SQR4_Mask struct{ mmio.UM32 }

func (rm SQR4_Mask) Load() SQR4_Bits   { return SQR4_Bits(rm.UM32.Load()) }
func (rm SQR4_Mask) Store(b SQR4_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ7() SQR4_Mask {
	return SQR4_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(SQ7)}}
}

func (p *ADC_Periph) SQ8() SQR4_Mask {
	return SQR4_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(SQ8)}}
}

func (p *ADC_Periph) SQ9() SQR4_Mask {
	return SQR4_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(SQ9)}}
}

func (p *ADC_Periph) SQ10() SQR4_Mask {
	return SQR4_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(SQ10)}}
}

func (p *ADC_Periph) SQ11() SQR4_Mask {
	return SQR4_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(SQ11)}}
}

func (p *ADC_Periph) SQ12() SQR4_Mask {
	return SQR4_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(SQ12)}}
}

type SQR5_Bits uint32

func (b SQR5_Bits) Field(mask SQR5_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR5_Bits) J(v int) SQR5_Bits {
	return SQR5_Bits(bits.Make32(v, uint32(mask)))
}

type SQR5 struct{ mmio.U32 }

func (r *SQR5) Bits(mask SQR5_Bits) SQR5_Bits { return SQR5_Bits(r.U32.Bits(uint32(mask))) }
func (r *SQR5) StoreBits(mask, b SQR5_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SQR5) SetBits(mask SQR5_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *SQR5) ClearBits(mask SQR5_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *SQR5) Load() SQR5_Bits               { return SQR5_Bits(r.U32.Load()) }
func (r *SQR5) Store(b SQR5_Bits)             { r.U32.Store(uint32(b)) }

type SQR5_Mask struct{ mmio.UM32 }

func (rm SQR5_Mask) Load() SQR5_Bits   { return SQR5_Bits(rm.UM32.Load()) }
func (rm SQR5_Mask) Store(b SQR5_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ1() SQR5_Mask {
	return SQR5_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(SQ1)}}
}

func (p *ADC_Periph) SQ2() SQR5_Mask {
	return SQR5_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(SQ2)}}
}

func (p *ADC_Periph) SQ3() SQR5_Mask {
	return SQR5_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(SQ3)}}
}

func (p *ADC_Periph) SQ4() SQR5_Mask {
	return SQR5_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(SQ4)}}
}

func (p *ADC_Periph) SQ5() SQR5_Mask {
	return SQR5_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(SQ5)}}
}

func (p *ADC_Periph) SQ6() SQR5_Mask {
	return SQR5_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(SQ6)}}
}

type JSQR_Bits uint32

func (b JSQR_Bits) Field(mask JSQR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JSQR_Bits) J(v int) JSQR_Bits {
	return JSQR_Bits(bits.Make32(v, uint32(mask)))
}

type JSQR struct{ mmio.U32 }

func (r *JSQR) Bits(mask JSQR_Bits) JSQR_Bits { return JSQR_Bits(r.U32.Bits(uint32(mask))) }
func (r *JSQR) StoreBits(mask, b JSQR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JSQR) SetBits(mask JSQR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *JSQR) ClearBits(mask JSQR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *JSQR) Load() JSQR_Bits               { return JSQR_Bits(r.U32.Load()) }
func (r *JSQR) Store(b JSQR_Bits)             { r.U32.Store(uint32(b)) }

type JSQR_Mask struct{ mmio.UM32 }

func (rm JSQR_Mask) Load() JSQR_Bits   { return JSQR_Bits(rm.UM32.Load()) }
func (rm JSQR_Mask) Store(b JSQR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JSQ1() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 68)), uint32(JSQ1)}}
}

func (p *ADC_Periph) JSQ2() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 68)), uint32(JSQ2)}}
}

func (p *ADC_Periph) JSQ3() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 68)), uint32(JSQ3)}}
}

func (p *ADC_Periph) JSQ4() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 68)), uint32(JSQ4)}}
}

func (p *ADC_Periph) JL() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 68)), uint32(JL)}}
}

type JDR1_Bits uint32

func (b JDR1_Bits) Field(mask JDR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JDR1_Bits) J(v int) JDR1_Bits {
	return JDR1_Bits(bits.Make32(v, uint32(mask)))
}

type JDR1 struct{ mmio.U32 }

func (r *JDR1) Bits(mask JDR1_Bits) JDR1_Bits { return JDR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *JDR1) StoreBits(mask, b JDR1_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JDR1) SetBits(mask JDR1_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *JDR1) ClearBits(mask JDR1_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *JDR1) Load() JDR1_Bits               { return JDR1_Bits(r.U32.Load()) }
func (r *JDR1) Store(b JDR1_Bits)             { r.U32.Store(uint32(b)) }

type JDR1_Mask struct{ mmio.UM32 }

func (rm JDR1_Mask) Load() JDR1_Bits   { return JDR1_Bits(rm.UM32.Load()) }
func (rm JDR1_Mask) Store(b JDR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JDATA() JDR1_Mask {
	return JDR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 72)), uint32(JDATA)}}
}

type JDR2_Bits uint32

func (b JDR2_Bits) Field(mask JDR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JDR2_Bits) J(v int) JDR2_Bits {
	return JDR2_Bits(bits.Make32(v, uint32(mask)))
}

type JDR2 struct{ mmio.U32 }

func (r *JDR2) Bits(mask JDR2_Bits) JDR2_Bits { return JDR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *JDR2) StoreBits(mask, b JDR2_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JDR2) SetBits(mask JDR2_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *JDR2) ClearBits(mask JDR2_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *JDR2) Load() JDR2_Bits               { return JDR2_Bits(r.U32.Load()) }
func (r *JDR2) Store(b JDR2_Bits)             { r.U32.Store(uint32(b)) }

type JDR2_Mask struct{ mmio.UM32 }

func (rm JDR2_Mask) Load() JDR2_Bits   { return JDR2_Bits(rm.UM32.Load()) }
func (rm JDR2_Mask) Store(b JDR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JDATA() JDR2_Mask {
	return JDR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 76)), uint32(JDATA)}}
}

type JDR3_Bits uint32

func (b JDR3_Bits) Field(mask JDR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JDR3_Bits) J(v int) JDR3_Bits {
	return JDR3_Bits(bits.Make32(v, uint32(mask)))
}

type JDR3 struct{ mmio.U32 }

func (r *JDR3) Bits(mask JDR3_Bits) JDR3_Bits { return JDR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *JDR3) StoreBits(mask, b JDR3_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JDR3) SetBits(mask JDR3_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *JDR3) ClearBits(mask JDR3_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *JDR3) Load() JDR3_Bits               { return JDR3_Bits(r.U32.Load()) }
func (r *JDR3) Store(b JDR3_Bits)             { r.U32.Store(uint32(b)) }

type JDR3_Mask struct{ mmio.UM32 }

func (rm JDR3_Mask) Load() JDR3_Bits   { return JDR3_Bits(rm.UM32.Load()) }
func (rm JDR3_Mask) Store(b JDR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JDATA() JDR3_Mask {
	return JDR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 80)), uint32(JDATA)}}
}

type JDR4_Bits uint32

func (b JDR4_Bits) Field(mask JDR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JDR4_Bits) J(v int) JDR4_Bits {
	return JDR4_Bits(bits.Make32(v, uint32(mask)))
}

type JDR4 struct{ mmio.U32 }

func (r *JDR4) Bits(mask JDR4_Bits) JDR4_Bits { return JDR4_Bits(r.U32.Bits(uint32(mask))) }
func (r *JDR4) StoreBits(mask, b JDR4_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JDR4) SetBits(mask JDR4_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *JDR4) ClearBits(mask JDR4_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *JDR4) Load() JDR4_Bits               { return JDR4_Bits(r.U32.Load()) }
func (r *JDR4) Store(b JDR4_Bits)             { r.U32.Store(uint32(b)) }

type JDR4_Mask struct{ mmio.UM32 }

func (rm JDR4_Mask) Load() JDR4_Bits   { return JDR4_Bits(rm.UM32.Load()) }
func (rm JDR4_Mask) Store(b JDR4_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JDATA() JDR4_Mask {
	return JDR4_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 84)), uint32(JDATA)}}
}

type DR_Bits uint32

func (b DR_Bits) Field(mask DR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR_Bits) J(v int) DR_Bits {
	return DR_Bits(bits.Make32(v, uint32(mask)))
}

type DR struct{ mmio.U32 }

func (r *DR) Bits(mask DR_Bits) DR_Bits { return DR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DR) StoreBits(mask, b DR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DR) SetBits(mask DR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DR) ClearBits(mask DR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DR) Load() DR_Bits             { return DR_Bits(r.U32.Load()) }
func (r *DR) Store(b DR_Bits)           { r.U32.Store(uint32(b)) }

type DR_Mask struct{ mmio.UM32 }

func (rm DR_Mask) Load() DR_Bits   { return DR_Bits(rm.UM32.Load()) }
func (rm DR_Mask) Store(b DR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) DATA() DR_Mask {
	return DR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 88)), uint32(DATA)}}
}

type SMPR0_Bits uint32

func (b SMPR0_Bits) Field(mask SMPR0_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SMPR0_Bits) J(v int) SMPR0_Bits {
	return SMPR0_Bits(bits.Make32(v, uint32(mask)))
}

type SMPR0 struct{ mmio.U32 }

func (r *SMPR0) Bits(mask SMPR0_Bits) SMPR0_Bits { return SMPR0_Bits(r.U32.Bits(uint32(mask))) }
func (r *SMPR0) StoreBits(mask, b SMPR0_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SMPR0) SetBits(mask SMPR0_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *SMPR0) ClearBits(mask SMPR0_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *SMPR0) Load() SMPR0_Bits                { return SMPR0_Bits(r.U32.Load()) }
func (r *SMPR0) Store(b SMPR0_Bits)              { r.U32.Store(uint32(b)) }

type SMPR0_Mask struct{ mmio.UM32 }

func (rm SMPR0_Mask) Load() SMPR0_Bits   { return SMPR0_Bits(rm.UM32.Load()) }
func (rm SMPR0_Mask) Store(b SMPR0_Bits) { rm.UM32.Store(uint32(b)) }
