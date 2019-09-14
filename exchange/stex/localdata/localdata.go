package localdata

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"runtime"

	"github.com/NzKSO/goExchange/exchange/stex/model"
)

const currencyPairsInfoFile = "currencyPairsInfo.json"

type currencyPairMap map[string]*model.CurrencyPair

func (m *currencyPairMap) UnmarshalJSON(data []byte) error {
	if *m == nil {
		*m = make(currencyPairMap)
	}

	dec := json.NewDecoder(bytes.NewReader(data))

	_, err := dec.Token()
	if err != nil {
		return err
	}

	for dec.More() {
		var pair model.CurrencyPair
		if err = dec.Decode(&pair); err != nil {
			return err
		}
		(*m)[pair.Symbol] = &pair
	}

	_, err = dec.Token()
	if err != nil {
		return err
	}

	return nil
}

// AllCurrencyPairs contain all information related to currency pair in stex
var AllCurrencyPairs currencyPairMap

func init() {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("no caller information")
	}

	data, err := ioutil.ReadFile(filepath.Join(filepath.Dir(file), currencyPairsInfoFile))
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, &AllCurrencyPairs); err != nil {
		panic(err)
	}
}
