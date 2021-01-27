package types

import (
	"fmt"

	"github.com/umbracle/fastrlp"
)

// MarshalRLP implements the Marshaler interface
func (r *Receipt) MarshalRLP() []byte {
	return r.MarshalRLPTo(nil)
}

// MarshalRLPTo implements the Marshaler interface
func (r *Receipt) MarshalRLPTo(dst []byte) []byte {
	ar := fastrlp.DefaultArenaPool.Get()
	dst = r.MarshalRLPWith(ar).MarshalTo(dst)
	fastrlp.DefaultArenaPool.Put(ar)
	return dst
}

// MarshalRLPWith implements the Marshaler interface
func (r *Receipt) MarshalRLPWith(ar *fastrlp.Arena) *fastrlp.Value {
	vv := ar.NewArray()

	if r.Status != nil {
		// Field 'Status'
		vv.Set(ar.NewUint(uint64(*r.Status)))
	} else {
		// Field 'Root'
		vv.Set(ar.NewBytes(r.Root[:]))
	}

	// Field 'CumulativeGasUsed'
	vv.Set(ar.NewUint(r.CumulativeGasUsed))

	// Field 'LogsBloom'
	vv.Set(ar.NewBytes(r.LogsBloom[:]))

	// Field 'Logs'
	{
		if len(r.Logs) == 0 {
			vv.Set(ar.NewNullArray())
		} else {
			v0 := ar.NewArray()
			for _, item := range r.Logs {
				v0.Set(item.MarshalRLPWith(ar))
			}
			vv.Set(v0)
		}
	}

	return vv
}

// UnmarshalRLP implements the Unmarshaler interface
func (r *Receipt) UnmarshalRLP(buf []byte) error {
	pr := fastrlp.DefaultParserPool.Get()
	defer fastrlp.DefaultParserPool.Put(pr)

	vv, err := pr.Parse(buf)
	if err != nil {
		return err
	}
	if err := r.UnmarshalRLPFrom(pr, vv); err != nil {
		return err
	}
	return nil
}

// UnmarshalRLPFrom implements the Unmarshaler interface
func (r *Receipt) UnmarshalRLPFrom(p *fastrlp.Parser, v *fastrlp.Value) error {
	elems, err := v.GetElems()
	if err != nil {
		return err
	}
	if num := len(elems); num != 5 {
		return fmt.Errorf("not enough elements to decode transaction, expected 5 but found %d", num)
	}

	buf, err := elems[0].Bytes()
	if err != nil {
		return err
	}
	switch size := len(buf); size {
	case 32:
		// Field 'Root'
		copy(r.Root[:], buf[:])
	case 1:
		// Field 'Status'
		r.SetStatus(ReceiptStatus(buf[0]))
	default:
		return fmt.Errorf("bad root/status size %d", size)
	}

	// Field 'CumulativeGasUsed'
	if r.CumulativeGasUsed, err = elems[1].GetUint64(); err != nil {
		return err
	}

	// Field 'LogsBloom'
	if _, err = elems[2].GetBytes(r.LogsBloom[:0], 256); err != nil {
		return err
	}

	// Field 'Logs'
	{
		subElems, err := elems[3].GetElems()
		if err != nil {
			return err
		}
		r.Logs = make([]*Log, len(subElems))
		for indx, elem := range subElems {
			bb := &Log{}
			if err := bb.UnmarshalRLPFrom(p, elem); err != nil {
				return err
			}
			r.Logs[indx] = bb
		}
	}
	return nil
}
