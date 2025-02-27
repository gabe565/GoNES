package cpu

import (
	"fmt"
	"log/slog"
	"runtime"
	"strings"

	"gabe565.com/gones/internal/log"
	"gabe565.com/gones/internal/memory"
)

func (c *CPU) Trace() string {
	if runtime.GOOS == "js" {
		return "DISABLED"
	}

	code := c.ReadMem(c.ProgramCounter)
	op := opcodes[code]
	if op == nil {
		return ""
	}

	begin := c.ProgramCounter
	hexDump := make([]byte, 1, 3)
	hexDump[0] = code
	var valAddr uint16
	var val byte
	switch op.Mode {
	case Implicit, Immediate, Accumulator, Relative:
		//
	default:
		valAddr, _ = c.getAbsoluteAddress(op.Mode, begin+1)
		val = c.bus.(memory.ReadSafe).ReadMemSafe(valAddr)
	}

	var trace string
	switch op.Len {
	case 1:
		if op.Mode == Accumulator {
			trace += "A "
		}
	case 2:
		addr := c.ReadMem(begin + 1)
		hexDump = append(hexDump, addr)

		switch op.Mode {
		case Immediate:
			trace += fmt.Sprintf("#$%02X", addr)
		case ZeroPage:
			trace += fmt.Sprintf("$%02X = %02X", valAddr, val)
		case ZeroPageX:
			trace += fmt.Sprintf("$%02X,X @ %02X = %02X", addr, valAddr, val)
		case ZeroPageY:
			trace += fmt.Sprintf("$%02X,Y @ %02X = %02X", addr, valAddr, val)
		case IndirectX:
			trace += fmt.Sprintf("($%02X,X) @ %02X = %04X = %02X", addr, addr+c.RegisterX, valAddr, val)
		case IndirectY:
			trace += fmt.Sprintf("($%02X),Y = %04X @ %04X = %02X", addr, valAddr-uint16(c.RegisterY), valAddr, val)
		case Implicit, Relative:
			// assuming local jumps: BNE, BVS, etc
			addr := uint16(int8(addr)) + begin + 2
			trace += fmt.Sprintf("$%04X", addr)
		default:
			slog.Error("Invalid addressing mode has len 2",
				"mode", op.Mode,
				"code", log.HexVal(code),
			)
		}
	case 3:
		addr := c.ReadMem16(begin + 1)
		hexDump = append(hexDump, uint8(addr&0xFF), uint8(addr>>8))

		switch op.Mode {
		case Indirect:
			trace += fmt.Sprintf("($%04X) = %04X", addr, valAddr)
		case Implicit, Relative:
			trace += fmt.Sprintf("$%04X", addr)
		case Absolute:
			if op.Code() == 0x4C { // JMP
				trace += fmt.Sprintf("$%04X", valAddr)
			} else {
				trace += fmt.Sprintf("$%04X = %02X", valAddr, val)
			}
		case AbsoluteX:
			trace += fmt.Sprintf("$%04X,X @ %04X = %02X", addr, valAddr, val)
		case AbsoluteY:
			trace += fmt.Sprintf("$%04X,Y @ %04X = %02X", addr, valAddr, val)
		default:
			slog.Error("Invalid addressing mode has len 3",
				"mode", op.Mode,
				"code", log.HexVal(code),
			)
		}
	}

	var hexStr string
	for _, v := range hexDump {
		hexStr += fmt.Sprintf("%02X ", v)
	}
	hexStr = strings.TrimSpace(hexStr)

	undocumented := ' '
	if op.Undocumented {
		undocumented = '*'
	}
	final := fmt.Sprintf(
		"%04X  %-8s %c%3s %-27s A:%02X X:%02X Y:%02X P:%02X SP:%02X",
		begin,
		hexStr,
		undocumented,
		op.Name,
		trace,
		c.Accumulator,
		c.RegisterX,
		c.RegisterY,
		c.Status.Get(),
		c.StackPointer,
	)
	return final
}
