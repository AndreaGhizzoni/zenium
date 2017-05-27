package samples

import (
	"crypto/rand"
	"github.com/AndreaGhizzoni/zenium/structures"
	"github.com/AndreaGhizzoni/zenium/util"
	"math/big"
)

// This is the generator of random numbers.
type SGenerator struct {
	min, max *big.Int
}

// NewSecureGenerator returns a new instance of samples.SGenerator type.
// This generator uses math/crypt to generate *big.Int numbers between min
// and max. error is returned if min > max or if min == nil || max == nil.
func NewSecureGenerator(min, max *big.Int) (*SGenerator, error) {
	if err := util.CheckBoundsIfNotNil(min, max); err != nil {
		return nil, err
	}
	return &SGenerator{min, max}, nil
}

func (this *SGenerator) generateInt() (*big.Int, error) {
	width := big.NewInt(0).Sub(this.max, this.min)
	randomInWidth, err := rand.Int(rand.Reader, width)
	if err != nil {
		return nil, err
	}

	return randomInWidth.Add(randomInWidth, this.min), nil
}

func (this *SGenerator) generateSlice(len *big.Int) ([]*big.Int, error) {
	randomSlice := []*big.Int{}
	for i := big.NewInt(0); i.Cmp(len) == -1; i.Add(i, util.One) {
		if random, err := this.generateInt(); err != nil {
			return nil, err
		} else {
			randomSlice = append(randomSlice, random)
		}
	}

	return randomSlice, nil
}

func (this *SGenerator) generateBound(width *big.Int) (*structures.Bound, error) {
	lowerBound, err := this.generateInt()
	if err != nil {
		return nil, err
	}
	upperBound := big.NewInt(0)

	lowerBoundPlusWidth := big.NewInt(0).Add(lowerBound, width)
	if lowerBoundPlusWidth.Cmp(this.max) == 1 { // lowerBound + width > max
		upperBound = this.max
		lowerBound.Sub(upperBound, width) // lowerBound = upperBound - width
	} else {
		upperBound.Add(lowerBound, width) // upperBound = lowerBound + width
	}
	return structures.NewBound(lowerBound, upperBound), nil
}

// Int generate a random *big.Int according to samples.SGenerator instanced.
// error is returned if generation fails.
func (this *SGenerator) Int() (*big.Int, error) {
	return this.generateInt()
}

// Slice generate a slice of length len. error is returned if len == nil or
// if single *big.Int generation fails.
func (this *SGenerator) Slice(len *big.Int) ([]*big.Int, error) {
	if err := util.IsNilOrLessThenOne(len, "Slice length"); err != nil {
		return nil, err
	}

	return this.generateSlice(len)
}

// Matrix generate a matrix with rows and columns given according to
// samples.SGenerator instanced. error is returned if: rows == nil,
// columns == nil, rows >= 1, columns >= 1 or if single *bit.Int generation
// fails.
func (this *SGenerator) Matrix(rows, columns *big.Int) ([][]*big.Int, error) {
	if err := util.IsNilOrLessThenOne(rows, "Matrix rows"); err != nil {
		return nil, err
	}

	if err := util.IsNilOrLessThenOne(columns, "Matrix columns"); err != nil {
		return nil, err
	}

	matrix := [][]*big.Int{}
	for i := big.NewInt(0); i.Cmp(rows) == -1; i.Add(i, util.One) {
		random, err := this.generateSlice(columns)
		if err != nil {
			return nil, err
		}
		matrix = append(matrix, random)
	}

	return matrix, nil
}

// This function generate a slice of random structures.Bound. width is the
// fixed with of all the bounds. amount is the number of bounds that will be
// generated. error is returned if: width == nil, width >= 1, width can not
// be placed between min and max or if single *bit.Int generation fails.
func (this *SGenerator) Bounds(width, amount *big.Int) ([]*structures.Bound, error) {
	if err := util.IsNilOrLessThenOne(width, "Bound width"); err != nil {
		return nil, err
	}

	err := util.IsWidthContainedInBounds(this.min, this.max, width)
	if err != nil {
		return nil, err
	}

	bounds := []*structures.Bound{}
	for i := big.NewInt(0); i.Cmp(amount) == -1; i.Add(i, util.One) {
		if bound, err := this.generateBound(width); err != nil {
			return nil, err
		} else {
			bounds = append(bounds, bound)
		}
	}
	return bounds, nil
}
