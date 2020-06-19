package newtype

import "github.com/shopspring/decimal"

// Decimal specifies the decimal type for this package.
type Decimal struct {
	decimal.Decimal
}

// FromDB implementing Conversion interface for xorm.
func (d *Decimal) FromDB(value []byte) (err error) {
	err = d.Scan(value)
	return
}

// ToDB implementing Conversion interface for xorm.
func (d Decimal) ToDB() (val []byte, err error) {
	valDecimal, _ := d.Value()
	val = []byte(valDecimal.(string))
	return
}
