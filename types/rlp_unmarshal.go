package types

import (
	"fmt"

	"github.com/umbracle/fastrlp"
)

type RLPUnmarshaler interface {
	RawRLPUnmarshaler

	UnmarshalRLP(input []byte) error
}

type RawRLPUnmarshaler interface {
	UnmarshalRLPFrom(p *fastrlp.Parser, v *fastrlp.Value) error
}

func UnmarshalRlp(obj RawRLPUnmarshaler, input []byte) error {
	pr := fastrlp.DefaultParserPool.Get()

	v, err := pr.Parse(input)
	if err != nil {
		fastrlp.DefaultParserPool.Put(pr)
		return err
	}
	if err := obj.UnmarshalRLPFrom(pr, v); err != nil {
		fastrlp.DefaultParserPool.Put(pr)
		return err
	}

	fastrlp.DefaultParserPool.Put(pr)
	return nil
}

func (b *Block) UnmarshalRLP(input []byte) error {
	return UnmarshalRlp(b, input)
}

func (b *Block) UnmarshalRLPFrom(p *fastrlp.Parser, v *fastrlp.Value) error {
	elems, err := v.GetElems()
	if err != nil {
		return err
	}
	if num := len(elems); num != 3 {
		return fmt.Errorf("not enough elements to decode block, expected 3 but found %d", num)
	}

	// header
	b.Header = &Header{}
	if err := b.Header.UnmarshalRLPFrom(p, elems[0]); err != nil {
		return err
	}

	// transactions
	txns, err := elems[1].GetElems()
	if err != nil {
		return err
	}
	for _, txn := range txns {
		bTxn := &Transaction{}
		if err := bTxn.UnmarshalRLPFrom(p, txn); err != nil {
			return err
		}
		b.Transactions = append(b.Transactions, bTxn)
	}

	// uncles
	uncles, err := elems[2].GetElems()
	if err != nil {
		return err
	}
	for _, uncle := range uncles {
		bUncle := &Header{}
		if err := bUncle.UnmarshalRLPFrom(p, uncle); err != nil {
			return err
		}
		b.Uncles = append(b.Uncles, bUncle)
	}

	return nil
}

func (b *Body) UnmarshalRLP(input []byte) error {
	return UnmarshalRlp(b, input)
}

func (b *Body) UnmarshalRLPFrom(p *fastrlp.Parser, v *fastrlp.Value) error {
	tuple, err := v.GetElems()
	if err != nil {
		return err
	}
	if len(tuple) != 2 {
		return fmt.Errorf("Two elements expected")
	}

	// transactions
	txns, err := tuple[0].GetElems()
	if err != nil {
		return err
	}
	for _, txn := range txns {
		bTxn := &Transaction{}
		if err := bTxn.UnmarshalRLPFrom(p, txn); err != nil {
			return err
		}
		b.Transactions = append(b.Transactions, bTxn)
	}

	// uncles
	uncles, err := tuple[1].GetElems()
	if err != nil {
		return err
	}
	for _, uncle := range uncles {
		bUncle := &Header{}
		if err := bUncle.UnmarshalRLPFrom(p, uncle); err != nil {
			return err
		}
		b.Uncles = append(b.Uncles, bUncle)
	}

	return nil
}

func (h *Header) UnmarshalRLP(input []byte) error {
	return UnmarshalRlp(h, input)
}

func (h *Header) UnmarshalRLPFrom(p *fastrlp.Parser, v *fastrlp.Value) error {
	elems, err := v.GetElems()
	if err != nil {
		return err
	}
	if num := len(elems); num != 15 {
		return fmt.Errorf("not enough elements to decode header, expected 15 but found %d", num)
	}

	p.Hash(h.Hash[:0], v)

	// parentHash
	if err = elems[0].GetHash(h.ParentHash[:]); err != nil {
		return err
	}
	// sha3uncles
	if err = elems[1].GetHash(h.Sha3Uncles[:]); err != nil {
		return err
	}
	// miner
	if err = elems[2].GetAddr(h.Miner[:]); err != nil {
		return err
	}
	// stateroot
	if err = elems[3].GetHash(h.StateRoot[:]); err != nil {
		return err
	}
	// txroot
	if err = elems[4].GetHash(h.TxRoot[:]); err != nil {
		return err
	}
	// receiptroot
	if err = elems[5].GetHash(h.ReceiptsRoot[:]); err != nil {
		return err
	}
	// logsBloom
	if _, err = elems[6].GetBytes(h.LogsBloom[:0], 256); err != nil {
		return err
	}
	// difficulty
	if h.Difficulty, err = elems[7].GetUint64(); err != nil {
		return err
	}
	// number
	if h.Number, err = elems[8].GetUint64(); err != nil {
		return err
	}
	// gasLimit
	if h.GasLimit, err = elems[9].GetUint64(); err != nil {
		return err
	}
	// gasused
	if h.GasUsed, err = elems[10].GetUint64(); err != nil {
		return err
	}
	// timestamp
	if h.Timestamp, err = elems[11].GetUint64(); err != nil {
		return err
	}
	// extraData
	if h.ExtraData, err = elems[12].GetBytes(h.ExtraData[:0]); err != nil {
		return err
	}
	// mixHash
	if err = elems[13].GetHash(h.MixHash[:0]); err != nil {
		return err
	}
	// nonce
	nonce, err := elems[14].GetUint64()
	if err != nil {
		return err
	}
	h.SetNonce(nonce)
	return err
}

func (r *Receipts) UnmarshalRLP(input []byte) error {
	return UnmarshalRlp(r, input)
}

func (r *Receipts) UnmarshalRLPFrom(p *fastrlp.Parser, v *fastrlp.Value) error {
	elems, err := v.GetElems()
	if err != nil {
		return err
	}
	for _, elem := range elems {
		rr := &Receipt{}
		if err := rr.UnmarshalRLPFrom(p, elem); err != nil {
			return err
		}
		(*r) = append(*r, rr)
	}
	return nil
}

func (r *Receipt) UnmarshalRLP(input []byte) error {
	return UnmarshalRlp(r, input)
}

// UnmarshalRLP unmarshals a Receipt in RLP format
func (r *Receipt) UnmarshalRLPFrom(p *fastrlp.Parser, v *fastrlp.Value) error {
	elems, err := v.GetElems()
	if err != nil {
		return err
	}
	if len(elems) != 4 {
		return fmt.Errorf("expected 4 elements")
	}

	// root or status
	buf, err := elems[0].Bytes()
	if err != nil {
		return err
	}
	switch size := len(buf); size {
	case 32:
		// root
		copy(r.Root[:], buf[:])
	case 1:
		// status
		r.SetStatus(ReceiptStatus(buf[0]))
	default:
		return fmt.Errorf("bad root/status size %d", size)
	}

	// cumulativeGasUsed
	if r.CumulativeGasUsed, err = elems[1].GetUint64(); err != nil {
		return err
	}
	// logsBloom
	if _, err = elems[2].GetBytes(r.LogsBloom[:0], 256); err != nil {
		return err
	}

	// logs
	logsElems, err := v.Get(3).GetElems()
	if err != nil {
		return err
	}
	for _, elem := range logsElems {
		log := &Log{}
		if err := log.UnmarshalRLPFrom(p, elem); err != nil {
			return err
		}
		r.Logs = append(r.Logs, log)
	}
	return nil
}

func (l *Log) UnmarshalRLPFrom(p *fastrlp.Parser, v *fastrlp.Value) error {
	elems, err := v.GetElems()
	if err != nil {
		return err
	}
	if len(elems) != 3 {
		return fmt.Errorf("bad elems")
	}

	// address
	if err := elems[0].GetAddr(l.Address[:]); err != nil {
		return err
	}
	// topics
	topicElems, err := elems[1].GetElems()
	if err != nil {
		return err
	}
	l.Topics = make([]Hash, len(topicElems))
	for indx, topic := range topicElems {
		if err := topic.GetHash(l.Topics[indx][:]); err != nil {
			return err
		}
	}
	// data
	if l.Data, err = elems[2].GetBytes(l.Data[:0]); err != nil {
		return err
	}
	return nil
}

func (t *Transaction) UnmarshalRLP(input []byte) error {
	return UnmarshalRlp(t, input)
}

// UnmarshalRLP unmarshals a Transaction in RLP format
func (t *Transaction) UnmarshalRLPFrom(p *fastrlp.Parser, v *fastrlp.Value) error {
	elems, err := v.GetElems()
	if err != nil {
		return err
	}
	if num := len(elems); num != 9 {
		return fmt.Errorf("not enough elements to decode transaction, expected 9 but found %d", num)
	}

	p.Hash(t.Hash[:0], v)

	// nonce
	if t.Nonce, err = elems[0].GetUint64(); err != nil {
		return err
	}
	// gasPrice
	if t.GasPrice, err = elems[1].GetBytes(t.GasPrice[:0]); err != nil {
		return err
	}
	// gas
	if t.Gas, err = elems[2].GetUint64(); err != nil {
		return err
	}
	// to
	vv, err := v.Get(3).Bytes()
	if len(vv) == 20 {
		// address
		addr := BytesToAddress(vv)
		t.To = &addr
	} else {
		// reset To
		t.To = nil
	}
	// value
	if t.Value, err = elems[4].GetBytes(t.Value[:0]); err != nil {
		return err
	}
	// input
	if t.Input, err = elems[5].GetBytes(t.Input[:0]); err != nil {
		return err
	}
	// v
	vv, err = v.Get(6).Bytes()
	if err != nil {
		return err
	}
	if len(vv) != 1 {
		return fmt.Errorf("only one byte expected")
	}
	t.V = byte(vv[0])
	// R
	if t.R, err = elems[7].GetBytes(t.R[:0]); err != nil {
		return err
	}
	// S
	if t.S, err = elems[8].GetBytes(t.S[:0]); err != nil {
		return err
	}
	return nil
}
