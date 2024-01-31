package api

type Location struct {
	Location            string `json:"location"`
	Txid                string `json:"txid"`
	Index               int    `json:"index"`
	ScriptHash          string `json:"scripthash"`
	Value               int    `json:"value"`
	Script              string `json:"script"`
	Address             string `json:"address"`
	AtomicalsAtLocation []any  `json:"atomicals_at_location"`
	TxNum               int    `json:"tx_num"`
}

type LocationInfo struct {
	Locations []Location `json:"locations"`
}

type BlockHeaderInfo struct {
	Version    int    `json:"version"`
	PrevHash   string `json:"prevHash"`
	MerkleRoot string `json:"merkleRoot"`
	Timestamp  int    `json:"timestamp"`
	Bits       int    `json:"bits"`
	Nonce      int    `json:"nonce"`
}

type Bitwork struct {
	Bitworkc string `json:"$bitworkc"`
	Bitworkr string `json:"$bitworkr"`
}

type MintInfo struct {
	CommitTxid               string                 `json:"commit_txid"`
	CommitIndex              int                    `json:"commit_index"`
	CommitLocation           string                 `json:"commit_location"`
	CommitTxNum              int                    `json:"commit_tx_num"`
	CommitHeight             int                    `json:"commit_height"`
	RevealLocationTxid       string                 `json:"reveal_location_txid"`
	RevealLocationIndex      int                    `json:"reveal_location_index"`
	RevealLocation           string                 `json:"reveal_location"`
	RevealLocationTxNum      int                    `json:"reveal_location_tx_num"`
	RevealLocationHeight     int                    `json:"reveal_location_height"`
	RevealLocationHeader     string                 `json:"reveal_location_header"`
	RevealLocationBlockhash  string                 `json:"reveal_location_blockhash"`
	RevealLocationScripthash string                 `json:"reveal_location_scripthash"`
	RevealLocationScript     string                 `json:"reveal_location_script"`
	RevealLocationValue      int                    `json:"reveal_location_value"`
	Args                     map[string]interface{} `json:"args"`
	Meta                     map[string]interface{} `json:"meta"`
	Ctx                      map[string]interface{} `json:"ctx"`
	Init                     map[string]interface{} `json:"init"`
	RevealLocationAddress    string                 `json:"reveal_location_address"`
	BlockheaderInfo          *BlockHeaderInfo       `json:"blockheader_info"`
	RequestRealm             string                 `json:"$request_realm"`
	RequestSubrealm          string                 `json:"$request_subrealm"`
	RequestContainer         string                 `json:"$request_container"`
	RequestTicker            string                 `json:"$request_ticker"`
	Pid                      string                 `json:"$pid"`
	Bitwork                  *Bitwork               `json:"$bitwork"`
}

type MintDataSummary struct {
	Fields map[string]interface{} `json:"fields"`
}

type StateInfo struct {
}

type RuleSet struct {
	Pattern string   `json:"pattern"`
	Outputs []Output `json:"outputs"`
}

type Output struct {
	V int    `json:"v"`
	S string `json:"s"`
}

type ApplicableRule struct {
	RuleSetTxid         string  `json:"rule_set_txid"`
	RuleSetHeight       int     `json:"rule_set_height"`
	RuleValidFromHeight int     `json:"rule_valid_from_height"`
	MatchedRule         RuleSet `json:"matched_rule"`
}

type SubrealmCandidate struct {
	TxNum                       int             `json:"tx_num"`
	AtomicalID                  string          `json:"atomical_id"`
	Txid                        string          `json:"txid"`
	CommitHeight                int             `json:"commit_height"`
	RevealLocationHeight        int             `json:"reveal_location_height"`
	Payment                     string          `json:"payment"`
	PaymentType                 string          `json:"payment_type"`
	MakePaymentFromHeight       int             `json:"make_payment_from_height"`
	PaymentDueNoLaterThanHeight string          `json:"payment_due_no_later_than_height"`
	ApplicableRule              *ApplicableRule `json:"applicable_rule"`
}

type RequestSubrealmStatus struct {
	Status                     string `json:"status"`
	VerifiedAtomicalID         string `json:"verified_atomical_id"`
	ClaimedByAtomicalID        string `json:"claimed_by_atomical_id"`
	PendingCandidateAtomicalID string `json:"pending_candidate_atomical_id"`
	PendingClaimedByAtomicalID string `json:"pending_claimed_by_atomical_id"`
	Note                       string `json:"note"`
}

type RequestNameStatus struct {
	Status                     string `json:"status"`
	VerifiedAtomicalID         string `json:"verified_atomical_id"`
	ClaimedByAtomicalID        string `json:"claimed_by_atomical_id"`
	PendingCandidateAtomicalID string `json:"pending_candidate_atomical_id"`
	Note                       string `json:"note"`
}

type NameCandidate struct {
	TxNum                int    `json:"tx_num"`
	AtomicalID           string `json:"atomical_id"`
	Txid                 string `json:"txid"`
	CommitHeight         int    `json:"commit_height"`
	RevealLocationHeight int    `json:"reveal_location_height"`
}

type AtomicalStatus struct {
	AtomicalID             string                  `json:"atomical_id"`
	AtomicalNumber         int                     `json:"atomical_number"`
	Type                   string                  `json:"type"`
	Subtype                string                  `json:"subtype"`
	LocationInfoObj        *LocationInfo           `json:"location_info_obj"`
	MintInfo               *MintInfo               `json:"mint_info"`
	MintData               *MintDataSummary        `json:"mint_data"`
	StateInfo              *StateInfo              `json:"state_info"`
	Relns                  *map[string]interface{} `json:"$relns"`
	Bitwork                *Bitwork                `json:"$bitwork"`
	RequestRealmStatus     *RequestNameStatus      `json:"$request_realm_status"`
	RealmCandidates        []NameCandidate         `json:"$realm_candidates"`
	RequestRealm           string                  `json:"$request_realm"`
	Realm                  string                  `json:"$realm"`
	FullRealmName          string                  `json:"$full_realm_name"`
	RequestFullRealmName   string                  `json:"$request_full_realm_name"`
	SubrealmCandidates     []SubrealmCandidate     `json:"$subrealm_candidates"`
	RequestSubrealmStatus  *RequestSubrealmStatus  `json:"$request_subrealm_status"`
	RequestSubrealm        string                  `json:"$request_subrealm"`
	Pid                    string                  `json:"$pid"`
	Subrealm               string                  `json:"$subrealm"`
	MaxSupply              int                     `json:"$max_supply"`
	MintHeight             int                     `json:"$mint_height"`
	MintAmount             int                     `json:"$mint_amount"`
	MaxMints               int                     `json:"$max_mints"`
	MintBitworkc           string                  `json:"$mint_bitworkc"`
	MintBitworkr           string                  `json:"$mint_bitworkr"`
	TickerCandidates       []NameCandidate         `json:"$ticker_candidates"`
	RequestTickerStatus    *RequestNameStatus      `json:"$request_ticker_status"`
	RequestTicker          string                  `json:"$request_ticker"`
	Ticker                 string                  `json:"$ticker"`
	RequestContainerStatus *RequestNameStatus      `json:"$request_container_status"`
	ContainerCandidates    []NameCandidate         `json:"$container_candidates"`
	RequestContainer       string                  `json:"$request_container"`
	Container              string                  `json:"$container"`
}
