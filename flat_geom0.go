package geom

type geom0 struct {
	layout     Layout
	stride     int
	flatCoords []float64
	srid       int
}

func (g *geom0) Bounds() *Bounds {
	return NewBounds(g.layout).extendFlatCoords(g.flatCoords, 0, len(g.flatCoords), g.stride)
}

func (g *geom0) Coords() Coord {
	return inflate0(g.flatCoords, 0, len(g.flatCoords), g.stride)
}

func (g *geom0) Ends() []int {
	return nil
}

func (g *geom0) Endss() [][]int {
	return nil
}

func (g *geom0) FlatCoords() []float64 {
	return g.flatCoords
}

func (g *geom0) Layout() Layout {
	return g.layout
}

func (g *geom0) NumCoords() int {
	return 1
}

func (g *geom0) Reserve(n int) {
	if cap(g.flatCoords) < n*g.stride {
		fcs := make([]float64, len(g.flatCoords), n*g.stride)
		copy(fcs, g.flatCoords)
		g.flatCoords = fcs
	}
}

func (g *geom0) SRID() int {
	return g.srid
}

func (g *geom0) setCoords(coords0 []float64) error {
	var err error
	g.flatCoords, err = deflate0(nil, coords0, g.stride)
	return err
}

func (g *geom0) Stride() int {
	return g.stride
}

// TransformXY replaces the X and Y ordinates of every coordinate in g with
// f(X, Y).
func (g *geom0) TransformXY(f func(float64, float64) (float64, float64)) error {
	if g.stride < 2 {
		return ErrStrideMismatch{Got: g.stride, Want: 2}
	}
	for i := 0; i < len(g.flatCoords); i += g.stride {
		g.flatCoords[i], g.flatCoords[i+1] = f(g.flatCoords[i], g.flatCoords[i+1])
	}
	return nil
}

// TransformXYZ replaces the X, Y, and Z ordinates of every coordinate in g
// with f(X, Y, Z).
func (g *geom0) TransformXYZ(f func(float64, float64, float64) (float64, float64, float64)) error {
	if g.stride < 3 {
		return ErrStrideMismatch{Got: g.stride, Want: 3}
	}
	for i := 0; i < len(g.flatCoords); i += g.stride {
		g.flatCoords[i], g.flatCoords[i+1], g.flatCoords[i+2] = f(g.flatCoords[i], g.flatCoords[i+1], g.flatCoords[i+2])
	}
	return nil
}

func (g *geom0) verify() error {
	if g.stride != g.layout.Stride() {
		return errStrideLayoutMismatch
	}
	if g.stride == 0 {
		if len(g.flatCoords) != 0 {
			return errNonEmptyFlatCoords
		}
		return nil
	}
	if len(g.flatCoords) != g.stride {
		return errLengthStrideMismatch
	}
	return nil
}
