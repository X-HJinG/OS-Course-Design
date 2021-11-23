package mysys

import (
	"strconv"
	"strings"
)

const (
	_ = iota
	K = 1 << (10 * iota)
	M
	G
	T
)

type LogicSpace struct {
	Bits int
}

type LogicAddr struct {
	No   int
	Addr int
}

type PageSize struct {
	Size int
}

func NewLogicSpace(bits int) *LogicSpace {
	return &LogicSpace{
		Bits: bits,
	}
}

func NewLogicAddrForPaging(addr int) *LogicAddr {
	return &LogicAddr{
		Addr: addr,
	}
}

func NewLogicAddrForSection(no int, addr int) *LogicAddr {
	return &LogicAddr{
		No:   no,
		Addr: addr,
	}
}

func NewPageSize(size string) *PageSize {
	res := ParseInt(size)
	return &PageSize{
		Size: res,
	}
}

func ParseInt(str string) int {
	str = strings.ToUpper(str)
	num := make([]rune, 0)
	bits := make([]rune, 0)
	for _, v := range str {
		if v >= '0' && v <= '9' {
			num = append(num, v)
		}
		if v >= 'A' && v <= 'Z' || v >= 'a' && v <= 'z' {
			bits = append(bits, v)
		}
	}
	i, _ := strconv.Atoi(string(num))
	parse := func(str string) int {
		switch str {
		case "K":
			return K
		case "M":
			return M
		case "G":
			return G
		case "T":
			return T
		default:
			return 1
		}
	}
	return i * parse(string(bits))
}
