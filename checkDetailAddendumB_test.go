// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"log"
	"strings"
	"testing"
)

// mockCheckDetailAddendumB creates a CheckDetailAddendumB
func mockCheckDetailAddendumB() CheckDetailAddendumB {
	cdAddendumB := NewCheckDetailAddendumB()
	cdAddendumB.ImageReferenceKeyIndicator = 1
	cdAddendumB.MicrofilmArchiveSequenceNumber = "1A             "
	cdAddendumB.LengthImageReferenceKey = "0034"
	cdAddendumB.ImageReferenceKey = "0"
	cdAddendumB.Description = "CD Addendum B"
	cdAddendumB.UserField = ""
	return cdAddendumB
}

func TestCheckDetailAddendumBParseErr(t *testing.T) {
	var c CheckDetailAddendumB
	c.Parse("asdhfakjfsa")
	if c.ImageReferenceKeyIndicator != 0 {
		t.Errorf("c.ImageReferenceKeyIndicator=%d", c.ImageReferenceKeyIndicator)
	}
	c.Parse("2711A             00340                                 CD Addendum B")
	if c.ImageReferenceKeyIndicator != 1 {
		t.Errorf("c.ImageReferenceKeyIndicator=%d", c.ImageReferenceKeyIndicator)
	}
	if c.ImageReferenceKey != "" {
		t.Errorf("c.ImageReferenceKey=%s", c.ImageReferenceKey)
	}
}

// TestMockCheckDetailAddendumB creates a CheckDetailAddendumB
func TestMockCheckDetailAddendumB(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	if err := cdAddendumB.Validate(); err != nil {
		t.Error("MockCheckDetailAddendumB does not validate and will break other tests: ", err)
	}
	if cdAddendumB.recordType != "27" {
		t.Error("recordType does not validate")
	}
	if cdAddendumB.ImageReferenceKeyIndicator != 1 {
		t.Error("ImageReferenceKeyIndicator does not validate")
	}
	if cdAddendumB.MicrofilmArchiveSequenceNumber != "1A             " {
		t.Error("MicrofilmArchiveSequenceNumber does not validate")
	}
	if cdAddendumB.LengthImageReferenceKey != "0034" {
		t.Error("LengthImageReferenceKey does not validate")
	}
	if cdAddendumB.ImageReferenceKey != "0" {
		t.Error("ImageReferenceKey does not validate")
	}
	if cdAddendumB.Description != "CD Addendum B" {
		t.Error("Description does not validate")
	}
	if cdAddendumB.UserField != "" {
		t.Error("UserField does not validate")
	}
}

// TestParseCheckDetailAddendumB validates parsing a CheckDetailAddendumB
func TestParseCheckDetailAddendumB(t *testing.T) {
	var line = "2711A             00340                                 CD Addendum B           "
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

	if err := r.parseCheckDetailAddendumB(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumB[0]
	if record.recordType != "27" {
		t.Errorf("RecordType Expected '27' got: %v", record.recordType)
	}
	if record.ImageReferenceKeyIndicatorField() != "1" {
		t.Errorf("ImageReferenceKeyIndicator Expected '1' got: %v",
			record.ImageReferenceKeyIndicatorField())
	}
	if record.MicrofilmArchiveSequenceNumberField() != "1A             " {
		t.Errorf("MicrofilmArchiveSequenceNumber Expected '1A             ' got: %v",
			record.MicrofilmArchiveSequenceNumberField())
	}
	if record.LengthImageReferenceKeyField() != "0034" {
		t.Errorf("ImageReferenceKeyLength Expected '0034' got: %v", record.LengthImageReferenceKeyField())
	}
	if record.ImageReferenceKeyField() != "0                                 " {
		t.Errorf("ImageReferenceKey Expected '0                                 ' got: %v",
			record.ImageReferenceKeyField())
	}
	if record.DescriptionField() != "CD Addendum B  " {
		t.Errorf("Description Expected 'CD Addendum B  ' got: %v", record.DescriptionField())
	}
	if record.UserFieldField() != "    " {
		t.Errorf("UserField Expected '    ' got: %v", record.UserFieldField())
	}
	if record.reservedField() != "     " {
		t.Errorf("reserved Expected '     ' got: %v", record.reservedField())
	}
}

// testCDAddendumBString validates that a known parsed CheckDetailAddendumB can return to a string of the same value
func testCDAddendumBString(t testing.TB) {
	var line = "2711A             00340                                 CD Addendum B           "
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

	if err := r.parseCheckDetailAddendumB(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentCashLetter.currentBundle.GetChecks()[0].CheckDetailAddendumB[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestCDAddendumB String tests validating that a known parsed CheckDetailAddendumB can return to a string of the
// same value
func TestCDAddendumBString(t *testing.T) {
	testCDAddendumBString(t)
}

// BenchmarkCDAddendumB String benchmarks validating that a known parsed CheckDetailAddendumB
// can return to a string of the same value
func BenchmarkCDAddendumBString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCDAddendumBString(b)
	}
}

// TestCDAddendumBRecordType validation
func TestCDAddendumBRecordType(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.recordType = "00"
	if err := cdAddendumB.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumBImageReferenceKeyIndicator validation
func TestCDAddendumBImageReferenceKeyIndicator(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.ImageReferenceKeyIndicator = 5
	if err := cdAddendumB.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImageReferenceKeyIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumBImageReferenceKey validation
func TestCDAddendumBImageReferenceKey(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.ImageReferenceKey = "????"
	if err := cdAddendumB.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImageReferenceKey" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumBDescription validation
func TestCDAddendumBDescription(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.Description = "????"
	if err := cdAddendumB.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "Description" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumBUserField validation
func TestCDAddendumBUserField(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.UserField = "????"
	if err := cdAddendumB.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserField" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// Field Inclusion

// TestCDAddendumBFIRecordType validation
func TestCDAddendumBFIRecordType(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.recordType = ""
	if err := cdAddendumB.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCDAddendumBFIMicrofilmArchiveSequenceNumber validation
func TestCDAddendumBFIMicrofilmArchiveSequenceNumber(t *testing.T) {
	cdAddendumB := mockCheckDetailAddendumB()
	cdAddendumB.MicrofilmArchiveSequenceNumber = "               "
	if err := cdAddendumB.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "MicrofilmArchiveSequenceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// End FieldInclusion

// TestNBSMFieldTrim validation
func TestNBSMFieldTrim(t *testing.T) {
	rdAddendumB := mockReturnDetailAddendumB()
	rdAddendumB.AuxiliaryOnUs = "12345678901234567890"
	if len(rdAddendumB.AuxiliaryOnUsField()) > 15 {
		t.Error("AuxiliaryOnUs field is greater than max")
	}

}
