package models

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type Technology string

const (
	Wifi        Technology = "wifi"
	Uart        Technology = "uart"
	Jtag        Technology = "jtag"
	Bluetooth   Technology = "bluetooth"
	Lte         Technology = "lte"
	Rfid        Technology = "rfid"
	Nfc         Technology = "nfc"
	Antplus     Technology = "ant+"
	Lifi        Technology = "lifi"
	Zigbee      Technology = "zigbee"
	Zwave       Technology = "z-wave"
	Lteadvanced Technology = "lte-advanced"
	Lra         Technology = "lora"
	NbIot       Technology = "nb-iot"
	Sigfox      Technology = "sigfox"
	NbFi        Technology = "nb-fi"
	Http        Technology = "http"
	Https       Technology = "https"
	Coap        Technology = "coap"
	Mqtt        Technology = "mqtt"
	Amqp        Technology = "amqp"
	Xmpp        Technology = "xmpp"
)

func (t *Technology) Scan(value interface{}) error {
	if value == nil {
		*t = ""
		return nil
	}
	strValue, ok := value.(string)
	if !ok {
		return errors.New("failed to scan Technology")
	}
	*t = Technology(strValue)
	return nil
}

func (t Technology) Value() (driver.Value, error) {
	return string(t), nil
}

type StringSlice []Technology

func (s *StringSlice) Scan(value interface{}) error {
	if value == nil {
		*s = StringSlice{}
		return nil
	}
	strValue, ok := value.(string)
	if !ok {
		return errors.New("failed to scan StringSlice")
	}
	technologies := strings.Split(strValue, ",")
	for _, tech := range technologies {
		*s = append(*s, Technology(tech))
	}
	return nil
}

func (s StringSlice) Value() (driver.Value, error) {
	techStrings := make([]string, len(s))
	for i, tech := range s {
		techStrings[i] = string(tech)
	}
	return strings.Join(techStrings, ","), nil
}
