package render

import "encoding/xml"

type VSDXSchema struct {
	XMLName              xml.Name `xml:"schema"`
	Text                 string   `xml:",chardata"`
	Xsd                  string   `xml:"xsd,attr"`
	TargetNamespace      string   `xml:"targetNamespace,attr"`
	Xmlns                string   `xml:"xmlns,attr"`
	R                    string   `xml:"r,attr"`
	ElementFormDefault   string   `xml:"elementFormDefault,attr"`
	AttributeFormDefault string   `xml:"attributeFormDefault,attr"`
	Import               struct {
		Text           string `xml:",chardata"`
		Namespace      string `xml:"namespace,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"import"`
	Annotation struct {
		Text          string `xml:",chardata"`
		Documentation string `xml:"documentation"`
	} `xml:"annotation"`
	Element []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
		Type string `xml:"type,attr"`
	} `xml:"element"`
	ComplexType []struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name,attr"`
		Abstract string `xml:"abstract,attr"`
		Mixed    string `xml:"mixed,attr"`
		Sequence struct {
			Text      string `xml:",chardata"`
			MinOccurs string `xml:"minOccurs,attr"`
			MaxOccurs string `xml:"maxOccurs,attr"`
			Element   []struct {
				Text      string `xml:",chardata"`
				Name      string `xml:"name,attr"`
				Type      string `xml:"type,attr"`
				MinOccurs string `xml:"minOccurs,attr"`
				MaxOccurs string `xml:"maxOccurs,attr"`
			} `xml:"element"`
			Any struct {
				Text            string `xml:",chardata"`
				MinOccurs       string `xml:"minOccurs,attr"`
				MaxOccurs       string `xml:"maxOccurs,attr"`
				Namespace       string `xml:"namespace,attr"`
				ProcessContents string `xml:"processContents,attr"`
			} `xml:"any"`
		} `xml:"sequence"`
		AnyAttribute struct {
			Text            string `xml:",chardata"`
			Namespace       string `xml:"namespace,attr"`
			ProcessContents string `xml:"processContents,attr"`
		} `xml:"anyAttribute"`
		Attribute []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
			Type string `xml:"type,attr"`
			Use  string `xml:"use,attr"`
			Ref  string `xml:"ref,attr"`
		} `xml:"attribute"`
		ComplexContent struct {
			Text      string `xml:",chardata"`
			Extension struct {
				Text      string `xml:",chardata"`
				Base      string `xml:"base,attr"`
				Attribute []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					Type string `xml:"type,attr"`
					Use  string `xml:"use,attr"`
				} `xml:"attribute"`
				Sequence struct {
					Text    string `xml:",chardata"`
					Element []struct {
						Text      string `xml:",chardata"`
						Name      string `xml:"name,attr"`
						Type      string `xml:"type,attr"`
						MinOccurs string `xml:"minOccurs,attr"`
						MaxOccurs string `xml:"maxOccurs,attr"`
					} `xml:"element"`
				} `xml:"sequence"`
			} `xml:"extension"`
		} `xml:"complexContent"`
		Choice struct {
			Text      string `xml:",chardata"`
			MinOccurs string `xml:"minOccurs,attr"`
			MaxOccurs string `xml:"maxOccurs,attr"`
			Element   []struct {
				Text      string `xml:",chardata"`
				Name      string `xml:"name,attr"`
				Type      string `xml:"type,attr"`
				MinOccurs string `xml:"minOccurs,attr"`
				MaxOccurs string `xml:"maxOccurs,attr"`
			} `xml:"element"`
		} `xml:"choice"`
		SimpleContent struct {
			Text      string `xml:",chardata"`
			Extension struct {
				Text      string `xml:",chardata"`
				Base      string `xml:"base,attr"`
				Attribute []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					Type string `xml:"type,attr"`
					Use  string `xml:"use,attr"`
				} `xml:"attribute"`
			} `xml:"extension"`
		} `xml:"simpleContent"`
		All struct {
			Text    string `xml:",chardata"`
			Element []struct {
				Text      string `xml:",chardata"`
				Name      string `xml:"name,attr"`
				Type      string `xml:"type,attr"`
				MinOccurs string `xml:"minOccurs,attr"`
				MaxOccurs string `xml:"maxOccurs,attr"`
			} `xml:"element"`
		} `xml:"all"`
	} `xml:"complexType"`
}
