package sizex

import "fmt"

type Byte int

const (
	B  Byte = 1
	Kb      = 1 << 10
	Mb      = Kb << 10
	Gb      = Mb << 10
	Tb      = Gb << 10
)

func (b Byte) String() string {
	switch {
	case b < Kb:
		return fmt.Sprintf("%db", b)
	case b >= Kb && b < Mb:
		return fmt.Sprintf("%dKb", b>>10)
	case b >= Mb && b < Gb:
		return fmt.Sprintf("%dMb", b>>20)
	case b >= Gb && b < Tb:
		return fmt.Sprintf("%dGb", b>>30)
	default:
		return fmt.Sprintf("%dTb", b>>40)
	}
}

func IsB(b Byte) bool {
	return b < Kb
}

func IsKb(b Byte) bool {
	return b >= Kb && b < Mb
}

func IsMb(b Byte) bool {
	return b >= Mb && b < Gb
}

func IsGb(b Byte) bool {
	return b >= Gb && b < Tb
}

func IsTb(b Byte) bool {
	return b >= Tb
}
