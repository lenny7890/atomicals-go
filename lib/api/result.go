package api

type ElectrumResponse[T any] struct {
	Global *GlobalVal `json:"global,omitempty"`
	Result T          `json:"result"`
}

type GlobalVal struct {
	Coin                 string            `json:"coin"`
	Network              string            `json:"network"`
	Height               int64             `json:"height"`
	BlockTip             string            `json:"block_tip"`
	ServerTime           string            `json:"server_time"`
	AtomicalsBlockTip    string            `json:"atomicals_block_tip"`
	AtomicalCount        int64             `json:"atomical_count"`
	AtomicalsBlockHashes map[string]string `json:"atomicals_block_hashes"`
}

type GetByTickerResult struct {
	Status              string      `json:"status"`
	CandidateAtomicalID string      `json:"candidate_atomical_id"`
	AtomicalID          string      `json:"atomical_id"`
	Candidates          []Candidate `json:"candidates"`
	Type                string      `json:"type"`
}

type Candidate struct {
	TxNum                int64  `json:"tx_num"`
	AtomicalID           string `json:"atomical_id"`
	TxID                 string `json:"txid"`
	CommitHeight         int64  `json:"commit_height"`
	RevealLocationHeight int64  `json:"reveal_location_height"`
}

type FtInfo struct {
	AtomicalID     string `json:"atomical_id"`
	AtomicalNumber int64  `json:"atomical_number"`
	AtomicalRef    string `json:"atomical_ref"`
	Type           string `json:"type"`
	Confirmed      bool   `json:"confirmed"`
	MintInfo       struct {
		CommitTxid               string `json:"commit_txid"`
		CommitIndex              int64  `json:"commit_index"`
		CommitLocation           string `json:"commit_location"`
		CommitTxNum              int64  `json:"commit_tx_num"`
		CommitHeight             int64  `json:"commit_height"`
		RevealLocationTxid       string `json:"reveal_location_txid"`
		RevealLocationIndex      int64  `json:"reveal_location_index"`
		RevealLocation           string `json:"reveal_location"`
		RevealLocationTxNum      int64  `json:"reveal_location_tx_num"`
		RevealLocationHeight     int64  `json:"reveal_location_height"`
		RevealLocationHeader     string `json:"reveal_location_header"`
		RevealLocationBlockhash  string `json:"reveal_location_blockhash"`
		RevealLocationScripthash string `json:"reveal_location_scripthash"`
		RevealLocationScript     string `json:"reveal_location_script"`
		RevealLocationValue      int64  `json:"reveal_location_value"`
		Args                     struct {
			MintAmount    int64  `json:"mint_amount"`
			MintHeight    int64  `json:"mint_height"`
			MaxMints      int64  `json:"max_mints"`
			MintBitworkc  string `json:"mint_bitworkc"`
			RequestTicker string `json:"request_ticker"`
			Bitworkc      string `json:"bitworkc"`
			Nonce         int64  `json:"nonce"`
			Time          int64  `json:"time"`
		} `json:"args"`
		Meta struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Legal       struct {
				Terms string `json:"terms"`
			} `json:"legal"`
		} `json:"meta"`
		Ctx struct {
		} `json:"ctx"`
		MintBitworkc  string `json:"$mint_bitworkc"`
		RequestTicker string `json:"$request_ticker"`
		Bitwork       struct {
			Bitworkc string      `json:"bitworkc"`
			Bitworkr interface{} `json:"bitworkr"`
		} `json:"$bitwork"`
	} `json:"mint_info"`
	Subtype      string `json:"subtype"`
	MintMode     string `json:"$mint_mode"`
	MaxSupply    int64  `json:"$max_supply"`
	MintHeight   int64  `json:"$mint_height"`
	MintAmount   int64  `json:"$mint_amount"`
	MaxMints     int64  `json:"$max_mints"`
	MintBitworkc string `json:"$mint_bitworkc"`
	Bitwork      struct {
		Bitworkc string      `json:"bitworkc"`
		Bitworkr interface{} `json:"bitworkr"`
	} `json:"$bitwork"`
	TickerCandidates []struct {
		TxNum                int64  `json:"tx_num"`
		AtomicalID           string `json:"atomical_id"`
		Txid                 string `json:"txid"`
		CommitHeight         int64  `json:"commit_height"`
		RevealLocationHeight int64  `json:"reveal_location_height"`
	} `json:"$ticker_candidates"`
	RequestTickerStatus struct {
		Status             string `json:"status"`
		VerifiedAtomicalID string `json:"verified_atomical_id"`
		Note               string `json:"note"`
	} `json:"$request_ticker_status"`
	RequestTicker string `json:"$request_ticker"`
	Ticker        string `json:"$ticker"`
	MintData      struct {
		Fields struct {
			ImagePng struct {
				Ct string `json:"$ct"`
				B  struct {
					Len  int64 `json:"$len"`
					Auto bool  `json:"$auto"`
				} `json:"$b"`
			} `json:"image.png"`
			Args struct {
				MintAmount    int64  `json:"mint_amount"`
				MintHeight    int64  `json:"mint_height"`
				MaxMints      int64  `json:"max_mints"`
				MintBitworkc  string `json:"mint_bitworkc"`
				RequestTicker string `json:"request_ticker"`
				Bitworkc      string `json:"bitworkc"`
				Nonce         int64  `json:"nonce"`
				Time          int64  `json:"time"`
			} `json:"args"`
			Meta struct {
				Name        string `json:"name"`
				Description string `json:"description"`
				Legal       struct {
					Terms string `json:"terms"`
				} `json:"legal"`
			} `json:"meta"`
		} `json:"fields"`
	} `json:"mint_data"`
	DftInfo struct {
		MintCount int64 `json:"mint_count"`
	} `json:"dft_info"`
	LocationSummary struct {
		UniqueHolders     int64 `json:"unique_holders"`
		CirculatingSupply int64 `json:"circulating_supply"`
	} `json:"location_summary"`
}
