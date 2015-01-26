// This example shows how to use USART as 1-wire master.
package main

import (
	"delay"
	"fmt"
	"rtos"

	"stm32/f4/gpio"
	"stm32/f4/irqs"
	"stm32/f4/periph"
	"stm32/f4/setup"
	"stm32/f4/usarts"
	"stm32/onedrv"
	"stm32/serial"
	"stm32/usart"

	"onewire"
)

const (
	Green = 12 + iota
	Orange
	Red
	Blue
)

var (
	leds = gpio.D

	ow  = usarts.USART6
	one = serial.New(ow, 8, 8)

	us   = usarts.USART2
	term = serial.New(us, 80, 8)
)

func init() {
	setup.Performance168(8)

	// LEDS

	periph.AHB1ClockEnable(periph.GPIOD)
	periph.AHB1Reset(periph.GPIOD)

	leds.SetMode(Green, gpio.Out)
	leds.SetMode(Orange, gpio.Out)
	leds.SetMode(Red, gpio.Out)
	leds.SetMode(Blue, gpio.Out)

	// USART

	periph.AHB1ClockEnable(periph.GPIOA)
	periph.AHB1Reset(periph.GPIOA)
	periph.APB1ClockEnable(periph.USART2)
	periph.APB1Reset(periph.USART2)

	port, tx, rx := gpio.A, 2, 3

	port.SetMode(tx, gpio.Alt)
	port.SetOutType(tx, gpio.PushPull)
	port.SetPull(tx, gpio.PullUp)
	port.SetOutSpeed(tx, gpio.Fast)
	port.SetAltFunc(tx, gpio.USART2)
	port.SetMode(rx, gpio.Alt)
	port.SetAltFunc(rx, gpio.USART2)

	us.SetBaudRate(115200, setup.APB1Clk)
	us.SetWordLen(usart.Bits8)
	us.SetParity(usart.None)
	us.SetStopBits(usart.Stop1b)
	us.SetMode(usart.Tx | usart.Rx)
	us.EnableIRQs(usart.RxNotEmptyIRQ)
	us.Enable()

	rtos.IRQ(irqs.USART2).UseHandler(termISR)
	rtos.IRQ(irqs.USART2).Enable()

	term.SetUnix(true)

	// 1-wire

	periph.AHB1ClockEnable(periph.GPIOC)
	periph.AHB1Reset(periph.GPIOC)
	periph.APB2ClockEnable(periph.USART6)
	periph.APB2Reset(periph.USART6)

	port, tx = gpio.C, 6

	port.SetMode(tx, gpio.Alt)
	port.SetOutType(tx, gpio.OpenDrain)
	port.SetOutSpeed(tx, gpio.Fast)
	port.SetAltFunc(tx, gpio.USART6)

	ow.SetWordLen(usart.Bits8)
	ow.SetParity(usart.None)
	ow.SetStopBits(usart.Stop1b)
	ow.SetMode(usart.Tx | usart.Rx)
	ow.SetHalfDuplex(true)
	ow.EnableIRQs(usart.RxNotEmptyIRQ)
	ow.Enable()

	rtos.IRQ(irqs.USART6).UseHandler(oneISR)
	rtos.IRQ(irqs.USART6).Enable()
}

func termISR() {
	term.IRQ()
}

func oneISR() {
	one.IRQ()
}

func blink(c, d int) {
	leds.SetBit(c)
	if d > 0 {
		delay.Millisec(d)
	} else {
		delay.Loop(-1e4 * d)
	}
	leds.ClearBit(c)
}

func checkErr(err error) {
	if err == nil {
		return
	}
	term.WriteString(err.Error())
	term.WriteByte('\n')
	for {
		blink(Red, 100)
		delay.Millisec(100)
	}
}

func ok() {
	term.WriteString("OK\n")
	blink(Green, 50)
}

func main() {
	drv := onedrv.UARTDriver{Serial: one, Clock: setup.APB2Clk}
	m := onewire.Master{Driver: &drv}
	dtypes := []onewire.Type{onewire.DS18S20, onewire.DS18B20, onewire.DS1822}

	term.WriteString("\nConfigure all DS18B20, DS1822 to 10bit resolution.\n")
	checkErr(m.SkipROM())
	checkErr(m.WriteScratchpad(127, -128, onewire.T10bit))

	// This algorithm doesn't work in case of parasite power mode.
	for {
		term.WriteString(
			"\nSending ConvertT command on the bus (SkipROM addressing).",
		)
		checkErr(m.SkipROM())
		checkErr(m.ConvertT())
		term.WriteString("\nWaiting until all devices finish the conversion")
		for {
			delay.Millisec(50)
			term.WriteByte('.')
			b, err := m.ReadBit()
			checkErr(err)
			if b != 0 {
				break
			}
		}
		term.WriteString("\nSearching for temperature sensors:\n")
		for _, typ := range dtypes {
			s := onewire.NewSearch(typ, false)
			for m.SearchNext(&s) {
				d := s.Dev()
				fmt.Fprint(term, d, fmt.Str(": "))
				checkErr(m.MatchROM(d))
				s, err := m.ReadScratchpad()
				checkErr(err)
				t, err := s.Temp16(typ)
				checkErr(err)
				fmt.Fprint(
					term,
					fmt.Int(t/16), fmt.Rune('.'), fmt.Int(t>>3&1*5),
					fmt.Str(" C\n"),
				)
			}
			checkErr(s.Err())
		}
		term.WriteString("Done.\n")
		delay.Millisec(4e3)
	}
}