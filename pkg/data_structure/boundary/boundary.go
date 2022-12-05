package boundary

type Boundary struct {
	Lower int
	Upper int
}

func (b *Boundary) Encompasses(other *Boundary) bool {
	return b.Lower <= other.Lower && b.Upper >= other.Upper
}

func (b *Boundary) EncompassedBy(other *Boundary) bool {
	return other.Lower <= b.Lower && other.Upper >= b.Upper
}

func (b *Boundary) Overlaps(other *Boundary) bool {
	if b.EncompassedBy(other) || b.Encompasses(other) {
		return true
	}

	if b.Lower <= other.Lower {
		return other.Lower <= b.Upper
	} else {
		return other.Upper >= b.Lower
	}
}
