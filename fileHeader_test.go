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

// mockFileHeader creates a FileHeader
func mockFileHeader() FileHeader {
	fh := NewFileHeader()
	fh.StandardLevel = "35"
	fh.TestFileIndicator = "T"
	fh.ImmediateDestination = "231380104"
	fh.ImmediateOrigin = "121042882"
	fh.FileCreationDate = time.Now()
	fh.FileCreationTime = time.Now()
	fh.ResendIndicator = "N"
	fh.ImmediateDestinationName = "Citadel"
	fh.ImmediateOriginName = "Wells Fargo"
	fh.FileIDModifier = ""
	fh.CountryCode = "US"
	fh.UserField = ""
	fh.CompanionDocumentIndicator = ""
	return fh
}

// testMockFileHeader creates a FileHeader
func testMockFileHeader(t testing.TB) {
	fh := mockFileHeader()
	if err := fh.Validate(); err != nil {
		t.Error("mockFileHeader does not validate and will break other tests: ", err)
	}
	if fh.recordType != "01" {
		t.Error("recordType does not validate")
	}
	if fh.StandardLevel != "35" {
		t.Error("StandardLevel does not validate")
	}
	if fh.TestFileIndicator != "T" {
		t.Error("TestFileIndicator does not validate")
	}
	if fh.ResendIndicator != "N" {
		t.Error("ResendIndicator does not validate")
	}
	if fh.ImmediateDestination != "231380104" {
		t.Error("DestinationRoutingNumber does not validate")
	}
	if fh.ImmediateOrigin != "121042882" {
		t.Error("ECEInstitutionRoutingNumber does not validate")
	}
	if fh.ImmediateDestinationName != "Citadel" {
		t.Error("ImmediateDestinationName does not validate")
	}
	if fh.ImmediateOriginName != "Wells Fargo" {
		t.Error("ImmediateOriginName does not validate")
	}
	if fh.FileIDModifier != "" {
		t.Error("FileIDModifier does not validate")
	}
	if fh.CountryCode != "US" {
		t.Error("CountryCode does not validate")
	}
	if fh.UserField != "" {
		t.Error("UserField does not validate")
	}
	if fh.CompanionDocumentIndicator != "" {
		t.Error("CompanionDocumentIndicator does not validate")
	}
}

// TestMockFileHeader tests creating a FileHeader
func TestMockFileHeader(t *testing.T) {
	testMockFileHeader(t)
}

// BenchmarkMockFileHeader benchmarks creating a FileHeader
func BenchmarkMockFileHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockFileHeader(b)
	}
}

// parseFileHeader validates parsing a FileHeader
func parseFileHeader(t testing.TB) {
	var line = "0135T231380104121042882201809051523NCitadel           Wells Fargo        US     "
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseFileHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.File.Header

	if record.recordType != "01" {
		t.Errorf("RecordType Expected '01' got: %v", record.recordType)
	}
	if record.StandardLevelField() != "35" {
		t.Errorf("StandardLevel Expected '35' got: %v", record.StandardLevelField())
	}
	if record.TestFileIndicatorField() != "T" {
		t.Errorf("TestFileIndicator 'T' got: %v", record.TestFileIndicatorField())
	}
	if record.ImmediateDestinationField() != "231380104" {
		t.Errorf("ImmediateDestination Expected '231380104' got: %v", record.ImmediateDestinationField())
	}
	if record.ImmediateOriginField() != "121042882" {
		t.Errorf("ImmediateOrigin Expected '121042882' got: %v", record.ImmediateOriginField())
	}
	if record.FileCreationDateField() != "20180905" {
		t.Errorf("FileCreationDate Expected '20180905' got:'%v'", record.FileCreationDateField())
	}
	if record.FileCreationTimeField() != "1523" {
		t.Errorf("FileCreationTime Expected '1523' got:'%v'", record.FileCreationTimeField())
	}
	if record.ResendIndicatorField() != "N" {
		t.Errorf("ResendIndicator Expected 'N' got: %v", record.ResendIndicatorField())
	}
	if record.ImmediateDestinationNameField() != "Citadel           " {
		t.Errorf("ImmediateDestinationName Expected 'Citadel           ' got:'%v'", record.ImmediateDestinationNameField())
	}
	if record.ImmediateOriginNameField() != "Wells Fargo       " {
		t.Errorf("ImmediateOriginName Expected 'Wells Fargo       ' got: '%v'", record.ImmediateOriginNameField())
	}
	if record.FileIDModifierField() != " " {
		t.Errorf("FileIDModifier Expected ' ' got:'%v'", record.FileIDModifierField())
	}
	if record.CountryCodeField() != "US" {
		t.Errorf("CountryCode Expected 'US' got:'%v'", record.CountryCodeField())
	}
	if record.UserFieldField() != "    " {
		t.Errorf("UserField Expected '    ' got:'%v'", record.UserFieldField())
	}
	if record.CompanionDocumentIndicatorField() != " " {
		t.Errorf("CompanionDocumentIndicator Expected ' ' got:'%v'", record.CompanionDocumentIndicatorField())
	}
}

// TestParseFileHeader tests validating parsing a FileHeader
func TestParseFileHeader(t *testing.T) {
	parseFileHeader(t)
}

// BenchmarkParseFileHeader benchmarks validating parsing a FileHeader
func BenchmarkParseFileHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseFileHeader(b)
	}
}

// testFHString validates that a known parsed FileHeader can return to a string of the same value
func testFHString(t testing.TB) {
	var line = "0135T231380104121042882201809051523NCitadel           Wells Fargo        US     "
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseFileHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.File.Header

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestFHString tests validating that a known parsed FileHeader can return to a string of the same value
func TestFHString(t *testing.T) {
	testFHString(t)
}

// BenchmarkFHString benchmarks validating that a known parsed FileHeader
// can return to a string of the same value
func BenchmarkFHString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFHString(b)
	}
}

// TestFHRecordType validation
func TestFHRecordType(t *testing.T) {
	fh := mockFileHeader()
	fh.recordType = "00"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestStandardLevel validation
func TestStandardLevel(t *testing.T) {
	fh := mockFileHeader()
	fh.StandardLevel = "01"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "StandardLevel" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestTestFileIndicator validation
func TestTestFileIndicator(t *testing.T) {
	fh := mockFileHeader()
	fh.TestFileIndicator = "S"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TestFileIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestResendIndicator validation
func TestResendIndicator(t *testing.T) {
	fh := mockFileHeader()
	fh.ResendIndicator = "R"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ResendIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestImmediateDestinationName validation
func TestImmediateDestinationName(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateDestinationName = "????"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImmediateDestinationName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestImmediateOriginName validation
func TestImmediateOriginName(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateOriginName = "????"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImmediateOriginName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFileIDModifier validation
func TestFileIDModifier(t *testing.T) {
	fh := mockFileHeader()
	fh.FileIDModifier = "--"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "FileIDModifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCountryCode validation
func TestCountryCode(t *testing.T) {
	fh := mockFileHeader()
	fh.CompanionDocumentIndicator = "D"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CompanionDocumentIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestCACountryCode validation
func TestCACountryCode(t *testing.T) {
	fh := mockFileHeader()
	fh.CountryCode = "CA"
	fh.CompanionDocumentIndicator = "1"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CompanionDocumentIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUserField validation
func TestUserFieldI(t *testing.T) {
	fh := mockFileHeader()
	fh.UserField = "????"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "UserField" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionRecordType validates FieldInclusion
func TestFHFieldInclusionRecordType(t *testing.T) {
	fh := mockFileHeader()
	fh.recordType = ""
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionStandardLevel validates FieldInclusion
func TestFHFieldInclusionStandardLevel(t *testing.T) {
	fh := mockFileHeader()
	fh.StandardLevel = ""
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "StandardLevel" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionTestFileIndicator validates FieldInclusion
func TestFHFieldInclusionTestFileIndicator(t *testing.T) {
	fh := mockFileHeader()
	fh.TestFileIndicator = ""
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TestFileIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionResendIndicator validates FieldInclusion
func TestFHFieldInclusionResendIndicator(t *testing.T) {
	fh := mockFileHeader()
	fh.ResendIndicator = ""
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ResendIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionImmediateDestination validates FieldInclusion
func TestFHFieldInclusionImmediateDestination(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateDestination = ""
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImmediateDestination" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionImmediateDestinationZero validates FieldInclusion
func TestFHFieldInclusionImmediateDestinationZero(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateDestination = "000000000"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImmediateDestination" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionImmediateOrigin validates FieldInclusion
func TestFHFieldInclusionImmediateOrigin(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateOrigin = ""
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImmediateOrigin" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionImmediateOriginZero validates FieldInclusion
func TestFHFieldInclusionImmediateOriginZero(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateOrigin = "000000000"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImmediateOrigin" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionCreationDate validates FieldInclusion
func TestFHFieldInclusionCreationDate(t *testing.T) {
	fh := mockFileHeader()
	fh.FileCreationDate = time.Time{}
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "FileCreationDate" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionCreationTime validates FieldInclusion
func TestFHFieldInclusionCreationTime(t *testing.T) {
	fh := mockFileHeader()
	fh.FileCreationTime = time.Time{}
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "FileCreationTime" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFileHeaderRuneCountInString validates RuneCountInString
func TestFileHeaderRuneCountInString(t *testing.T) {
	fh := NewFileHeader()
	var line = "01"
	fh.Parse(line)

	if fh.ImmediateOrigin != "" {
		t.Error("Parsed with an invalid RuneCountInString")
	}
}
