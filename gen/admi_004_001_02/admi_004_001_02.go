// Code generated by xsdgen. DO NOT EDIT.

package admi_004_001_02

import (
	"github.com/moov-io/rtp20022/pkg/dt"
)

type Document struct {
	SysEvtNtfctn SystemEventNotificationV02 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 SysEvtNtfctn"`
}

type Event2 struct {
	EvtCd    Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtCd"`
	EvtParam []*Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtParam,omitempty"`
	EvtDesc  *Max1000Text         `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtDesc,omitempty"`
	EvtTm    *dt.ISODateTime      `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtTm,omitempty"`
}

// May be no more than 1000 items long
type Max1000Text string

// May be no more than 35 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

type SystemEventNotificationV02 struct {
	EvtInf Event2 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtInf"`
}
