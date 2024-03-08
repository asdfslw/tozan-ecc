package std

import (
	"fmt"
	"math/big"

	C "github.com/armfazh/tozan-ecc/curve"
	GF "github.com/armfazh/tozan-ecc/field"
)

// ID is an identifier of a std curve.
type ID string

const (
	SECP256K1 ID = "secp256k1"
	STARKNET  ID = "starknet"
	ED25519   ID = "ed25519"
	FOURQ     ID = "fourQ"
)

type params struct {
	model C.Model
	p     *big.Int
	m     int
	a, b  interface{}
	r     *big.Int
	h     int
	x, y  interface{}
}

// Curves is a list of toy curves.
var Curves []ID
var stdCurves map[ID]*params

func init() {
	Curves = make([]ID, 0, 4)
	stdCurves = make(map[ID]*params)

	{
		p := new(big.Int)
		a := new(big.Int)
		b := new(big.Int)
		r := new(big.Int)
		x := new(big.Int)
		y := new(big.Int)

		p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
		a.SetString("00", 16)
		b.SetString("07", 16)
		r.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
		x.SetString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 16)
		y.SetString("483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 16)
		SECP256K1.register(&params{model: C.Weierstrass, p: p, m: 1, a: a, b: b, r: r, h: 1, x: x, y: y})
	}
	{
		p := new(big.Int)
		a := new(big.Int)
		b := new(big.Int)
		r := new(big.Int)
		x := new(big.Int)
		y := new(big.Int)

		p.SetString("0800000000000011000000000000000000000000000000000000000000000001", 16)
		a.SetString("01", 16)
		b.SetString("3141592653589793238462643383279502884197169399375105820974944592307816406665", 10)
		r.SetString("0800000000000010ffffffffffffffffb781126dcae7b2321e66a241adc64d2f", 16)
		x.SetString("01ef15c18599971b7beced415a40f0c7deacfd9b0d1819e03d723d8bc943cfca", 16)
		y.SetString("005668060aa49730b7be4801df46ec62de53ecd11abe43a32873000c36e8dc1f", 16)
		STARKNET.register(&params{model: C.Weierstrass, p: p, m: 1, a: a, b: b, r: r, h: 1, x: x, y: y})
	}
	{
		p := new(big.Int)
		a := new(big.Int)
		b := new(big.Int)
		r := new(big.Int)
		x := new(big.Int)
		y := new(big.Int)

		p.SetString("7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed", 16)
		a.SetString("7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec", 16)
		b.SetString("52036cee2b6ffe738cc740797779e89800700a4d4141d8ab75eb4dca135978a3", 16)
		r.SetString("1000000000000000000000000000000014def9dea2f79cd65812631a5cf5d3ed", 16)
		x.SetString("216936d3cd6e53fec0a4e231fdd6dc5c692cc7609525a7b2c9562d608f25d51a", 16)
		y.SetString("6666666666666666666666666666666666666666666666666666666666666658", 16)
		ED25519.register(&params{model: C.TwistedEdwards, p: p, m: 1, a: a, b: b, r: r, h: 8, x: x, y: y})
	}
	{
		p := new(big.Int)
		a0 := new(big.Int)
		a1 := new(big.Int)
		b0 := new(big.Int)
		b1 := new(big.Int)
		r := new(big.Int)
		x0 := new(big.Int)
		x1 := new(big.Int)
		y0 := new(big.Int)
		y1 := new(big.Int)

		p.SetString("7fffffffffffffffffffffffffffffff", 16)
		a0.SetString("7ffffffffffffffffffffffffffffffe", 16)
		a1.SetString("00", 16)
		b0.SetString("e40000000000000142", 16)
		b1.SetString("5e472f846657e0fcb3821488f1fc0c8d", 16)
		r.SetString("0029cbc14e5e0a72f05397829cbc14e5dfbd004dfe0f79992fb2540ec7768ce7", 16)
		x0.SetString("1A3472237C2FB305286592AD7B3833AA", 16)
		x1.SetString("1E1F553F2878AA9C96869FB360AC77F6", 16)
		y0.SetString("0E3FEE9BA120785AB924A2462BCBB287", 16)
		y1.SetString("6E1C4AF8630E024249A7C344844C8B5C", 16)
		FOURQ.register(&params{model: C.TwistedEdwards, p: p, m: 2,
			a: []interface{}{a0, a1}, b: []interface{}{b0, b1},
			r: r, h: 8 * 49,
			x: []interface{}{x0, x1}, y: []interface{}{y0, y1},
		})
	}
}

func (id ID) register(p *params) { stdCurves[id] = p; Curves = append(Curves, id) }

// New returns an elliptic curve and a generator point.
func (id ID) New() (C.EllCurve, C.Point, error) {
	if v, ok := stdCurves[id]; ok {
		var F GF.Field
		if v.m == 1 {
			F = GF.NewFp(fmt.Sprintf("%v", v.p), v.p)
		} else if v.m == 2 {
			F = GF.NewFp2(fmt.Sprintf("%v", v.p), v.p)
		}
		E := v.model.New(string(id), F,
			F.Elt(v.a), F.Elt(v.b),
			v.r, big.NewInt(int64(v.h)))
		P := E.NewPoint(F.Elt(v.x), F.Elt(v.y))
		return E, P, nil
	}
	return nil, nil, fmt.Errorf("curve not supported")
}
