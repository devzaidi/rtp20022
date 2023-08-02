// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 with prefix 'rt'
package camt_056_001_08

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification6) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "rt:FinInstnId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification6TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "rt:FinInstnId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CancellationReason33Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Cd, xml.StartElement{Name: xml.Name{Local: "rt:Cd"}})
	e.EncodeElement(v.Prtry, xml.StartElement{Name: xml.Name{Local: "rt:Prtry"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CancellationReason33ChoiceTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Cd, xml.StartElement{Name: xml.Name{Local: "rt:Cd"}})
	e.EncodeElement(v.Prtry, xml.StartElement{Name: xml.Name{Local: "rt:Prtry"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Case5) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "rt:Id"}})
	e.EncodeElement(v.Cretr, xml.StartElement{Name: xml.Name{Local: "rt:Cretr"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Case5TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "rt:Id"}})
	e.EncodeElement(v.Cretr, xml.StartElement{Name: xml.Name{Local: "rt:Cretr"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CaseAssignment5) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "rt:Id"}})
	e.EncodeElement(v.Assgnr, xml.StartElement{Name: xml.Name{Local: "rt:Assgnr"}})
	e.EncodeElement(v.Assgne, xml.StartElement{Name: xml.Name{Local: "rt:Assgne"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "rt:CreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CaseAssignment5TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "rt:Id"}})
	e.EncodeElement(v.Assgnr, xml.StartElement{Name: xml.Name{Local: "rt:Assgnr"}})
	e.EncodeElement(v.Assgne, xml.StartElement{Name: xml.Name{Local: "rt:Assgne"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "rt:CreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MmbId, xml.StartElement{Name: xml.Name{Local: "rt:MmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification2TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MmbId, xml.StartElement{Name: xml.Name{Local: "rt:MmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Contact4) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.PhneNb, xml.StartElement{Name: xml.Name{Local: "rt:PhneNb"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DateAndPlaceOfBirth1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.BirthDt, xml.StartElement{Name: xml.Name{Local: "rt:BirthDt"}})
	e.EncodeElement(v.CityOfBirth, xml.StartElement{Name: xml.Name{Local: "rt:CityOfBirth"}})
	e.EncodeElement(v.CtryOfBirth, xml.StartElement{Name: xml.Name{Local: "rt:CtryOfBirth"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FIToFIPmtCxlReq, xml.StartElement{Name: xml.Name{Local: "rt:FIToFIPmtCxlReq"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification18) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.ClrSysMmbId, xml.StartElement{Name: xml.Name{Local: "rt:ClrSysMmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification18TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.ClrSysMmbId, xml.StartElement{Name: xml.Name{Local: "rt:ClrSysMmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FIToFIPaymentCancellationRequestV08) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Assgnmt, xml.StartElement{Name: xml.Name{Local: "rt:Assgnmt"}})
	e.EncodeElement(v.Case, xml.StartElement{Name: xml.Name{Local: "rt:Case"}})
	e.EncodeElement(v.Undrlyg, xml.StartElement{Name: xml.Name{Local: "rt:Undrlyg"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FIToFIPaymentCancellationRequestV08TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Assgnmt, xml.StartElement{Name: xml.Name{Local: "rt:Assgnmt"}})
	e.EncodeElement(v.Case, xml.StartElement{Name: xml.Name{Local: "rt:Case"}})
	e.EncodeElement(v.Undrlyg, xml.StartElement{Name: xml.Name{Local: "rt:Undrlyg"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GenericOrganisationIdentification1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "rt:Id"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GenericOrganisationIdentification1TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "rt:Id"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OrganisationIdentification29) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.LEI, xml.StartElement{Name: xml.Name{Local: "rt:LEI"}})
	e.EncodeElement(v.Othr, xml.StartElement{Name: xml.Name{Local: "rt:Othr"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OrganisationIdentification29TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Othr, xml.StartElement{Name: xml.Name{Local: "rt:Othr"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OrganisationIdentification29TCH2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.LEI, xml.StartElement{Name: xml.Name{Local: "rt:LEI"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OriginalGroupHeader15) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlMsgId, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlMsgId"}})
	e.EncodeElement(v.OrgnlMsgNmId, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlMsgNmId"}})
	e.EncodeElement(v.OrgnlCreDtTm, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlCreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OriginalTransactionReference28) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Dbtr, xml.StartElement{Name: xml.Name{Local: "rt:Dbtr"}})
	e.EncodeElement(v.Cdtr, xml.StartElement{Name: xml.Name{Local: "rt:Cdtr"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OriginalTransactionReference28TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Dbtr, xml.StartElement{Name: xml.Name{Local: "rt:Dbtr"}})
	e.EncodeElement(v.Cdtr, xml.StartElement{Name: xml.Name{Local: "rt:Cdtr"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Party38Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgId, xml.StartElement{Name: xml.Name{Local: "rt:OrgId"}})
	e.EncodeElement(v.PrvtId, xml.StartElement{Name: xml.Name{Local: "rt:PrvtId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Party38ChoiceTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgId, xml.StartElement{Name: xml.Name{Local: "rt:OrgId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Party38ChoiceTCH2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgId, xml.StartElement{Name: xml.Name{Local: "rt:OrgId"}})
	e.EncodeElement(v.PrvtId, xml.StartElement{Name: xml.Name{Local: "rt:PrvtId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Party40Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Pty, xml.StartElement{Name: xml.Name{Local: "rt:Pty"}})
	e.EncodeElement(v.Agt, xml.StartElement{Name: xml.Name{Local: "rt:Agt"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Party40ChoiceTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Agt, xml.StartElement{Name: xml.Name{Local: "rt:Agt"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Party40ChoiceTCH2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Pty, xml.StartElement{Name: xml.Name{Local: "rt:Pty"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Party40ChoiceTCH3) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Pty, xml.StartElement{Name: xml.Name{Local: "rt:Pty"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PartyIdentification135) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Nm, xml.StartElement{Name: xml.Name{Local: "rt:Nm"}})
	e.EncodeElement(v.PstlAdr, xml.StartElement{Name: xml.Name{Local: "rt:PstlAdr"}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "rt:Id"}})
	e.EncodeElement(v.CtctDtls, xml.StartElement{Name: xml.Name{Local: "rt:CtctDtls"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PartyIdentification135TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Nm, xml.StartElement{Name: xml.Name{Local: "rt:Nm"}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "rt:Id"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PartyIdentification135TCH2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Nm, xml.StartElement{Name: xml.Name{Local: "rt:Nm"}})
	e.EncodeElement(v.PstlAdr, xml.StartElement{Name: xml.Name{Local: "rt:PstlAdr"}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "rt:Id"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PartyIdentification135TCH3) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Nm, xml.StartElement{Name: xml.Name{Local: "rt:Nm"}})
	e.EncodeElement(v.PstlAdr, xml.StartElement{Name: xml.Name{Local: "rt:PstlAdr"}})
	e.EncodeElement(v.Id, xml.StartElement{Name: xml.Name{Local: "rt:Id"}})
	e.EncodeElement(v.CtctDtls, xml.StartElement{Name: xml.Name{Local: "rt:CtctDtls"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentCancellationReason5) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Orgtr, xml.StartElement{Name: xml.Name{Local: "rt:Orgtr"}})
	e.EncodeElement(v.Rsn, xml.StartElement{Name: xml.Name{Local: "rt:Rsn"}})
	e.EncodeElement(v.AddtlInf, xml.StartElement{Name: xml.Name{Local: "rt:AddtlInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentCancellationReason5TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Orgtr, xml.StartElement{Name: xml.Name{Local: "rt:Orgtr"}})
	e.EncodeElement(v.Rsn, xml.StartElement{Name: xml.Name{Local: "rt:Rsn"}})
	e.EncodeElement(v.AddtlInf, xml.StartElement{Name: xml.Name{Local: "rt:AddtlInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentTransaction106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlInstrId, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlInstrId"}})
	e.EncodeElement(v.OrgnlEndToEndId, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlEndToEndId"}})
	e.EncodeElement(v.OrgnlTxId, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlTxId"}})
	e.EncodeElement(v.OrgnlUETR, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlUETR"}})
	e.EncodeElement(v.OrgnlClrSysRef, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlClrSysRef"}})
	e.EncodeElement(v.OrgnlIntrBkSttlmAmt, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlIntrBkSttlmAmt"}})
	e.EncodeElement(v.OrgnlIntrBkSttlmDt, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlIntrBkSttlmDt"}})
	e.EncodeElement(v.CxlRsnInf, xml.StartElement{Name: xml.Name{Local: "rt:CxlRsnInf"}})
	e.EncodeElement(v.OrgnlTxRef, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlTxRef"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentTransaction106TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlInstrId, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlInstrId"}})
	e.EncodeElement(v.OrgnlEndToEndId, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlEndToEndId"}})
	e.EncodeElement(v.OrgnlTxId, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlTxId"}})
	e.EncodeElement(v.OrgnlUETR, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlUETR"}})
	e.EncodeElement(v.OrgnlClrSysRef, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlClrSysRef"}})
	e.EncodeElement(v.OrgnlIntrBkSttlmAmt, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlIntrBkSttlmAmt"}})
	e.EncodeElement(v.OrgnlIntrBkSttlmDt, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlIntrBkSttlmDt"}})
	e.EncodeElement(v.CxlRsnInf, xml.StartElement{Name: xml.Name{Local: "rt:CxlRsnInf"}})
	e.EncodeElement(v.OrgnlTxRef, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlTxRef"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PersonIdentification13) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.DtAndPlcOfBirth, xml.StartElement{Name: xml.Name{Local: "rt:DtAndPlcOfBirth"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PersonIdentification13TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.DtAndPlcOfBirth, xml.StartElement{Name: xml.Name{Local: "rt:DtAndPlcOfBirth"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PostalAddress24) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.StrtNm, xml.StartElement{Name: xml.Name{Local: "rt:StrtNm"}})
	e.EncodeElement(v.BldgNb, xml.StartElement{Name: xml.Name{Local: "rt:BldgNb"}})
	e.EncodeElement(v.PstCd, xml.StartElement{Name: xml.Name{Local: "rt:PstCd"}})
	e.EncodeElement(v.TwnNm, xml.StartElement{Name: xml.Name{Local: "rt:TwnNm"}})
	e.EncodeElement(v.CtrySubDvsn, xml.StartElement{Name: xml.Name{Local: "rt:CtrySubDvsn"}})
	e.EncodeElement(v.Ctry, xml.StartElement{Name: xml.Name{Local: "rt:Ctry"}})
	e.EncodeElement(v.AdrLine, xml.StartElement{Name: xml.Name{Local: "rt:AdrLine"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PostalAddress24TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.StrtNm, xml.StartElement{Name: xml.Name{Local: "rt:StrtNm"}})
	e.EncodeElement(v.BldgNb, xml.StartElement{Name: xml.Name{Local: "rt:BldgNb"}})
	e.EncodeElement(v.PstCd, xml.StartElement{Name: xml.Name{Local: "rt:PstCd"}})
	e.EncodeElement(v.TwnNm, xml.StartElement{Name: xml.Name{Local: "rt:TwnNm"}})
	e.EncodeElement(v.CtrySubDvsn, xml.StartElement{Name: xml.Name{Local: "rt:CtrySubDvsn"}})
	e.EncodeElement(v.Ctry, xml.StartElement{Name: xml.Name{Local: "rt:Ctry"}})
	e.EncodeElement(v.AdrLine, xml.StartElement{Name: xml.Name{Local: "rt:AdrLine"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v UnderlyingTransaction23) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlGrpInfAndCxl, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlGrpInfAndCxl"}})
	e.EncodeElement(v.TxInf, xml.StartElement{Name: xml.Name{Local: "rt:TxInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v UnderlyingTransaction23TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlGrpInfAndCxl, xml.StartElement{Name: xml.Name{Local: "rt:OrgnlGrpInfAndCxl"}})
	e.EncodeElement(v.TxInf, xml.StartElement{Name: xml.Name{Local: "rt:TxInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

func (a ActiveOrHistoricCurrencyAndAmountSimpleType) MarshalText() ([]byte, error) {
	return rtp.Amount(a).MarshalText()
}
