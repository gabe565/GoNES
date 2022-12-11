package ppu

import "image/color"

var SystemPalette = []color.RGBA{
	0x00: {0x58, 0x58, 0x58, 0xFF},
	0x01: {0x00, 0x23, 0x8C, 0xFF},
	0x02: {0x00, 0x13, 0x9B, 0xFF},
	0x03: {0x2D, 0x05, 0x85, 0xFF},
	0x04: {0x5D, 0x00, 0x52, 0xFF},
	0x05: {0x7A, 0x00, 0x17, 0xFF},
	0x06: {0x7A, 0x08, 0x00, 0xFF},
	0x07: {0x5F, 0x18, 0x00, 0xFF},
	0x08: {0x35, 0x2A, 0x00, 0xFF},
	0x09: {0x09, 0x39, 0x00, 0xFF},
	0x0A: {0x00, 0x3F, 0x00, 0xFF},
	0x0B: {0x00, 0x3C, 0x22, 0xFF},
	0x0C: {0x00, 0x32, 0x5D, 0xFF},
	0x0D: {0x00, 0x00, 0x00, 0xFF},
	0x0E: {0x00, 0x00, 0x00, 0xFF},
	0x0F: {0x00, 0x00, 0x00, 0xFF},
	0x10: {0xA1, 0xA1, 0xA1, 0xFF},
	0x11: {0x00, 0x53, 0xEE, 0xFF},
	0x12: {0x15, 0x3C, 0xFE, 0xFF},
	0x13: {0x60, 0x28, 0xE4, 0xFF},
	0x14: {0xA9, 0x1D, 0x98, 0xFF},
	0x15: {0xD4, 0x1E, 0x41, 0xFF},
	0x16: {0xD2, 0x2C, 0x00, 0xFF},
	0x17: {0xAA, 0x44, 0x00, 0xFF},
	0x18: {0x6C, 0x5E, 0x00, 0xFF},
	0x19: {0x2D, 0x73, 0x00, 0xFF},
	0x1A: {0x00, 0x7D, 0x06, 0xFF},
	0x1B: {0x00, 0x78, 0x52, 0xFF},
	0x1C: {0x00, 0x69, 0xA9, 0xFF},
	0x1D: {0x00, 0x00, 0x00, 0xFF},
	0x1E: {0x00, 0x00, 0x00, 0xFF},
	0x1F: {0x00, 0x00, 0x00, 0xFF},
	0x20: {0xFF, 0xFF, 0xFF, 0xFF},
	0x21: {0x1F, 0xA5, 0xFE, 0xFF},
	0x22: {0x5E, 0x89, 0xFE, 0xFF},
	0x23: {0xB5, 0x72, 0xFE, 0xFF},
	0x24: {0xFE, 0x65, 0xF6, 0xFF},
	0x25: {0xFE, 0x67, 0x90, 0xFF},
	0x26: {0xFE, 0x77, 0x3C, 0xFF},
	0x27: {0xFE, 0x93, 0x08, 0xFF},
	0x28: {0xC4, 0xB2, 0x00, 0xFF},
	0x29: {0x79, 0xCA, 0x10, 0xFF},
	0x2A: {0x3A, 0xD5, 0x4A, 0xFF},
	0x2B: {0x11, 0xD1, 0xA4, 0xFF},
	0x2C: {0x06, 0xBF, 0xFE, 0xFF},
	0x2D: {0x42, 0x42, 0x42, 0xFF},
	0x2E: {0x00, 0x00, 0x00, 0xFF},
	0x2F: {0x00, 0x00, 0x00, 0xFF},
	0x30: {0xFF, 0xFF, 0xFF, 0xFF},
	0x31: {0xA0, 0xD9, 0xFE, 0xFF},
	0x32: {0xBD, 0xCC, 0xFE, 0xFF},
	0x33: {0xE1, 0xC2, 0xFE, 0xFF},
	0x34: {0xFE, 0xBC, 0xFB, 0xFF},
	0x35: {0xFE, 0xBD, 0xD0, 0xFF},
	0x36: {0xFE, 0xC5, 0xA9, 0xFF},
	0x37: {0xFE, 0xD1, 0x8E, 0xFF},
	0x38: {0xE9, 0xDE, 0x86, 0xFF},
	0x39: {0xC7, 0xE9, 0x92, 0xFF},
	0x3A: {0xA8, 0xEE, 0xB0, 0xFF},
	0x3B: {0x95, 0xEC, 0xD9, 0xFF},
	0x3C: {0x91, 0xE4, 0xFE, 0xFF},
	0x3D: {0xAC, 0xAC, 0xAC, 0xFF},
	0x3E: {0x00, 0x00, 0x00, 0xFF},
	0x3F: {0x00, 0x00, 0x00, 0xFF},
}
