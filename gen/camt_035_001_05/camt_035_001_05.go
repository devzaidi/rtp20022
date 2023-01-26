// Code generated by xsdgen. DO NOT EDIT.

package camt_035_001_05

import (
	"encoding/xml"
	"github.com/moov-io/rtp20022/pkg/dt"
)

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 FinInstnId"`
}

type Case5 struct {
	Id    Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Id"`
	Cretr Party40Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Cretr"`
}

type CaseAssignment5 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Id"`
	Assgnr  Party40Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Assgnr"`
	Assgne  Party40Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Assgne"`
	CreDtTm dt.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 CreDtTm"`
}

type ClearingSystemMemberIdentification2 struct {
	MmbId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 MmbId"`
}

type Document struct {
	PrtryFrmtInvstgtn ProprietaryFormatInvestigationV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 PrtryFrmtInvstgtn"`
}

type FinancialInstitutionIdentification18 struct {
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 ClrSysMmbId"`
}

// May be no more than 140 items long
type Max140Text string

// May be no more than 35 items long
type Max35Text string

type Party40Choice struct {
	Pty *PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Agt,omitempty"`
}

type PartyIdentification135 struct {
	Nm Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Nm"`
}

type ProprietaryData6reduced struct {
	Ustrd    *Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Ustrd,omitempty"`
	OrigRefs TransactionReferences8reduced `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 OrigRefs"`
}

type ProprietaryData7TCH struct {
	Tp   Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Tp"`
	Data ProprietaryData6reduced `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Data"`
}

type ProprietaryFormatInvestigationV05 struct {
	Attr []xml.Attr `xml:",attr"`

	Assgnmt   CaseAssignment5     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Assgnmt"`
	Case      Case5               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 Case"`
	PrtryData ProprietaryData7TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 PrtryData"`
}

type TransactionReferences8reduced struct {
	InstrId    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 InstrId"`
	EndToEndId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 EndToEndId"`
	TxId       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 TxId"`
}
