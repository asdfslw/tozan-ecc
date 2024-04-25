package std

import (
	"fmt"
	"math/big"

	C "github.com/asdfslw/tozan-ecc/curve"
	GF "github.com/asdfslw/tozan-ecc/field"
)

// ID is an identifier of a std curve.
type ID string

const (
	SECP256K1    ID = "secp256k1"
	SECP256R1    ID = "secp256r1"
	STARKNET     ID = "starknet"
	PALLAS       ID = "Pallas"
	ED25519      ID = "ed25519"
	ZCASH_JUBJUB ID = "zcash_jubjub"
	FOURQ        ID = "fourQ"
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
	Curves = make([]ID, 0, 7)
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

		p.SetString("FFFFFFFF00000001000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFF", 16)
		a.SetString("FFFFFFFF00000001000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFC", 16)
		b.SetString("5AC635D8AA3A93E7B3EBBD55769886BC651D06B0CC53B0F63BCE3C3E27D2604B", 16)
		r.SetString("FFFFFFFF00000000FFFFFFFFFFFFFFFFBCE6FAADA7179E84F3B9CAC2FC632551", 16)
		x.SetString("6B17D1F2E12C4247F8BCE6E563A440F277037D812DEB33A0F4A13945D898C296", 16)
		y.SetString("4FE342E2FE1A7F9B8EE7EB4A7C0F9E162BCE33576B315ECECBB6406837BF51F5", 16)
		SECP256R1.register(&params{model: C.Weierstrass, p: p, m: 1, a: a, b: b, r: r, h: 1, x: x, y: y})
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

		p.SetString("40000000000000000000000000000000224698fc094cf91b992d30ed00000001", 16)
		a.SetString("00", 16)
		b.SetString("05", 10)
		r.SetString("40000000000000000000000000000000224698fc0994a8dd8c46eb2100000001", 16)
		x.SetString("01", 16)
		y.SetString("1b74b5a30a12937c53dfa9f06378ee548f655bd4333d477119cf7a23caed2abb", 16)
		PALLAS.register(&params{model: C.Weierstrass, p: p, m: 1, a: a, b: b, r: r, h: 1, x: x, y: y})
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
		a := new(big.Int)
		b := new(big.Int)
		r := new(big.Int)
		x := new(big.Int)
		y := new(big.Int)

		p.SetString("73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001", 16)
		a.SetString("73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000000", 16)
		b.SetString("2a9318e74bfa2b48f5fd9207e6bd7fd4292d7f6d37579d2601065fd6d6343eb1", 16)
		r.SetString("0e7db4ea6533afa906673b0101343b00a6682093ccc81082d0970e5ed6f72cb7", 16)
		x.SetString("0926d4f32059c712d418a7ff26753b6ad5b9a7d3ef8e282747bf46920a95a753", 16)
		y.SetString("57a1019e6de9b67553bb37d0c21cfd056d65674dcedbddbc305632adaaf2b530", 16)
		ZCASH_JUBJUB.register(&params{model: C.TwistedEdwards, p: p, m: 1, a: a, b: b, r: r, h: 8, x: x, y: y})
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
