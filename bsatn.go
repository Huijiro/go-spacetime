package main

import (
	"bytes"
	"encoding/binary"
	"math"
)

type SumTypeVariant struct {
	name           *[]uint32
	algebraic_type AlgebraicValue
}

type SumValue struct {
	variants []SumTypeVariant
}

type ProductTypeVariant struct {
	name           []uint32
	algebraic_type AlgebraicValue
}

type ProductValue struct {
	elements []ProductTypeVariant
}

type ArrayValue struct {
	elem_ty []AlgebraicValue
}

type AlgebraicValue interface {
	SumValue | ProductValue | ArrayValue |
		string |
		bool |
		uint8 | uint16 | uint32 | uint64 |
		int8 | int16 | int32 | int64 |
		float32 | float64 |
		any
}

func BSATN(algebraic AlgebraicValue) []byte {
	var buffer bytes.Buffer

	switch v := algebraic.(type) {
	case uint8:
		buffer.WriteByte(v)
	case uint16, uint32, uint64, int8, int16, int32, int64:
		binary.Write(&buffer, binary.LittleEndian, v)
	case float32, float64:
		switch v := v.(type) {
		case float32:
			value := BSATN(math.Float32bits(v))
			buffer.Write(value)
		case float64:
			value := BSATN(math.Float64bits(v))
			buffer.Write(value)
		}
	case bool:
		if v {
			buffer.WriteByte(1)
		} else {
			buffer.WriteByte(0)
		}
	case string:
		value := append(BSATN(uint32(len(v))), BSATN([]byte(v))...)
		buffer.Write(value)
	case ProductValue:
		for _, element := range v.elements {
			buffer.Write(BSATN(element.algebraic_type))
		}
	default:
		panic("unsupported type")
	}
	return buffer.Bytes()
}
