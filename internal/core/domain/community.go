package domain

import "time"

type CommunityResolver struct {
	ID         string      `json:"id"`
	MsgToken   string      `json:"msgToken"`
	Notify     bool        `json:"notify"`
	CRecharges []CRecharge `json:"cRecharges"`
	Resolvers  []Resolver  `json:"resolvers"`
}

type CRecharge struct {
	ResolvedAt time.Time `json:"resolvedAt"`
	Target     uint16    `json:"target"`
	Mount      uint8     `json:"mount"`
	Gain       uint8     `json:"gain"`
}

// CommunityRecharge real time data base (Firestore)
type CommunityRecharge struct {
	Company    string    `json:"company"`
	CreatedAt  time.Time `json:"createdAt"`
	ExecCode   string    `json:"execCode"`
	IDRecharge string    `json:"idRecharge"`
	Mount      int       `json:"mount"`
	State      int       `json:"state"`
}
