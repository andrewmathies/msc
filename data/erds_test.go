package data

import "testing"

func TestChecksValidation(t *testing.T) {
	e := &ERD{
		Name:    "Erd_BluetoothStatusProbe",
		Address: "0x6069",
		Fields: []string{
			"Erd_BluetoothStatusProbe_Connected?",
			"Erd_BluetoothStatusProbe_Temperature",
			"Erd_BluetoothStatusProbe_Synced?",
		},
	}

	err := e.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
