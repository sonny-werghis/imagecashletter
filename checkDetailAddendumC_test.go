// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"log"
	"strings"
	"testing"
	"time"
)

// mockCheckDetailAddendumC creates a CheckDetailAddendumC
func mockCheckDetailAddendumC() CheckDetailAddendumC {
	cdAddendumC := NewCheckDetailAddendumC()
	cdAddendumC.RecordNumber = 1
	cdAddendumC.EndorsingBankRoutingNumber = "121042882"
	cdAddendumC.BOFDEndorsementBusinessDate = time.Now()
	cdAddendumC.EndorsingBankItemSequenceNumber = "1              "
	cdAddendumC.TruncationIndicator = "Y"
	cdAddendumC.EndorsingBankConversionIndicator = "1"
	cdAddendumC.EndorsingBankCorrectionIndicator = 0
	cdAddendumC.ReturnReason = "A"
	cdAddendumC.UserField = ""
	cdAddendumC.EndorsingBankIdentifier = 0
	return cdAddendumC
}

// TestMockCheckDetailAddendumC creates a CheckDetailAddendumC
func TestMockCheckDetailAddendumC(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	if err := cdAddendumC.Validate(); err != nil {
		t.Error("mockCheckDetailAddendumC does not validate and will break other tests: ", err)
	}
	if cdAddendumC.recordType != "28" {
		t.Error("recordType does not validate")
	}
	if cdAddendumC.RecordNumber != 1 {
		t.Error("RecordNumber does not validate")
	}
	if cdAddendumC.EndorsingBankRoutingNumber != "121042882" {
		t.Error("EndorsingBankRoutingNumber does not validate")
	}
	if cdAddendumC.EndorsingBankItemSequenceNumber != "1              " {
		t.Error("EndorsingBankItemSequenceNumber does not validate")
	}
	if cdAddendumC.TruncationIndicator != "Y" {
		t.Error("TruncationIndicator does not validate")
	}
	if cdAddendumC.ReturnReason != "A" {
		t.Error("ReturnReason does not validate")
	}
	if cdAddendumC.EndorsingBankConversionIndicator != "1" {
		t.Error("EndorsingBankConversionIndicator does not validate")
	}
	if cdAddendumC.EndorsingBankCorrectionIndicator != 0 {
		t.Error("EndorsingBankCorrectionIndicator does not validate")
	}
	if cdAddendumC.UserField != "" {
		t.Error("UserField does not validate")
	}
	if cdAddendumC.EndorsingBankIdentifier != 0 {
		t.Error("EndorsingBankIdentifier does not validate")
	}
}

func TestParseCheckDetailAddendumCErr(t *testing.T) {
	var c CheckDetailAddendumC
	c.Parse("asdsakjahsfa")
	if c.RecordNumber != 0 {
		t.Errorf("c.RecordNumber=%d", c.RecordNumber)
	}
}

// TestParseCheckDetailAddendumC validates parsing a CheckDetailAddendumC
func TestParseCheckDetailAddendumC(t *testing.T) {
	var line = "2801121042882201809051              Y10A                   0                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	if err := r.parseCheckDetailAddendumC(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumC[0]

	if record.recordType != "28" {
		t.Errorf("RecordType Expected '28' got: %v", record.recordType)
	}
	if record.RecordNumberField() != "01" {
		t.Errorf("RecordNumber Expected '01' got: %v", record.RecordNumberField())
	}

	if record.EndorsingBankRoutingNumberField() != "121042882" {
		t.Errorf("EndorsingBankRoutingNumbeRoutingNumber Expected '121042882' got: %v",
			record.EndorsingBankRoutingNumberField())
	}
	if record.BOFDEndorsementBusinessDateField() != "20180905" {
		t.Errorf("BOFDEndorsementBusinessDate Expected '20180905' got: %v",
			record.BOFDEndorsementBusinessDateField())
	}
	if record.EndorsingBankItemSequenceNumberField() != "1              " {
		t.Errorf("EndorsingBankItemSequenceNumber Expected '1              ' got: %v",
			record.EndorsingBankItemSequenceNumberField())
	}
	if record.TruncationIndicatorField() != "Y" {
		t.Errorf("TruncationIndicator Expected 'Y' got: %v", record.TruncationIndicatorField())
	}
	if record.EndorsingBankConversionIndicatorField() != "1" {
		t.Errorf("EndorsingBankConversionIndicator  Expected '1' got: %v", record.EndorsingBankConversionIndicatorField())
	}
	if record.EndorsingBankCorrectionIndicatorField() != "0" {
		t.Errorf("EndorsingBankCorrectionIndicator Expected '0' got: %v", record.EndorsingBankCorrectionIndicatorField())
	}
	if record.ReturnReasonField() != "A" {
		t.Errorf("ReturnReason  Expected 'A' got: %v", record.ReturnReasonField())
	}
	if record.UserFieldField() != "                   " {
		t.Errorf("UserField Expected '                   ' got: %v", record.UserFieldField())
	}
	if record.reservedField() != "                    " {
		t.Errorf("reserved Expected '                    ' got: %v", record.reservedField())
	}
}

// testCDAddendumCString validates that a known parsed CheckDetailAddendumC can return to a string of the same value
func testCDAddendumCString(t testing.TB) {
	var line = "2801121042882201809051              Y10A                   0                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	clh := mockCashLetterHeader()
	r.addCurrentCashLetter(NewCashLetter(clh))
	bh := mockBundleHeader()
	b := NewBundle(bh)
	r.currentCashLetter.AddBundle(b)
	r.addCurrentBundle(b)
	cd := mockCheckDetail()
	r.currentCashLetter.currentBundle.AddCheckDetail(cd)

	if err := r.parseCheckDetailAddendumC(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumC[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestCDAddendumCString tests validating that a known parsed CheckDetailAddendumC can return to a string of the
// same value
func TestCDAddendumCString(t *testing.T) {
	testCDAddendumCString(t)
}

// BenchmarkCDAddendumCString benchmarks validating that a known parsed CheckDetailAddendumC
// can return to a string of the same value
func BenchmarkCDAddendumCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCDAddendumCString(b)
	}
}

// TestCDAddendumCRecordType validation
func TestCDAddendumCRecordType(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.recordType = "00"
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumCEndorsingBankRoutingNumber validation
func TestCDAddendumCEndorsingBankRoutingNumber(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankRoutingNumber = "Z"
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EndorsingBankRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumCTruncationIndicator validation
func TestCDAddendumCTruncationIndicator(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.TruncationIndicator = "A"
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TruncationIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumCEndorsingBankConversionIndicator validation
func TestCDAddendumCPayeeName(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankConversionIndicator = "10"
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EndorsingBankConversionIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumCEndorsingBankCorrectionIndicator validation
func TestCDAddendumCEndorsingBankCorrectionIndicator(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankCorrectionIndicator = 6
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EndorsingBankCorrectionIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumCBOFDReturnReason validation
func TestCDAddendumCBOFDReturnReason(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.ReturnReason = "????"
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnReason" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumCUserField validation
func TestCDAddendumCUserField(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.UserField = "????"
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserField" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumCEndorsingBankIdentifier validation
func TestCDAddendumCEndorsingBankIdentifier(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankIdentifier = 10
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EndorsingBankIdentifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// FieldInclusion

// TestCDAddendumCFIRecordType validation
func TestCDAddendumCFIRecordType(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.recordType = ""
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumCFIRecordNumber validation
func TestCDAddendumCFIRecordNumber(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.RecordNumber = 0
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "RecordNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumCFIEndorsingBankRoutingNumber validation
func TestCDAddendumCFIEndorsingBankRoutingNumber(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankRoutingNumber = ""
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EndorsingBankRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumCFIEndorsingBankRoutingNumberZero validation
func TestCDAddendumCFIEndorsingBankRoutingNumberZero(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankRoutingNumber = "000000000"
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EndorsingBankRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumCFIBOFDEndorsementBusinessDate validation
func TestCDAddendumCFIBOFDEndorsementBusinessDate(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.BOFDEndorsementBusinessDate = time.Time{}
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BOFDEndorsementBusinessDate" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumCFIEndorsingBankItemSequenceNumber validation
func TestCDAddendumCFIEndorsingBankItemSequenceNumber(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.EndorsingBankItemSequenceNumber = "               "
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EndorsingBankItemSequenceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumCFITruncationIndicator validation
func TestCDAddendumCFITruncationIndicator(t *testing.T) {
	cdAddendumC := mockCheckDetailAddendumC()
	cdAddendumC.TruncationIndicator = ""
	if err := cdAddendumC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TruncationIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
