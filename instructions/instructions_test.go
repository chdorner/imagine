package instructions

import (
	"net/url"
	"testing"
)

func TestParseInstrunctions(t *testing.T) {
	queryString := "origin=http%3A%2F%2Fexample.com%2Ftest.jpg" +
		"&action=crop" +
		"&format=png" +
		"&width=200" +
		"&height=100"
	query, err := url.ParseQuery(queryString)
	if err != nil {
		t.Fatal(err)
	}

	instructions, err := ParseInstructions(query)
	if err != nil {
		t.Fatal(err)
	}

	if instructions.Origin != "http://example.com/test.jpg" {
		t.Fatalf("origin is not as expected, got: %s", instructions.Origin)
	}

	if instructions.Action != "crop" {
		t.Fatalf("action is not as expected, got: %s", instructions.Action)
	}

	if instructions.Format != "png" {
		t.Fatalf("format is not as expected, got: %s", instructions.Format)
	}

	if instructions.Width != 200 {
		t.Fatalf("width is not as expected, got: %d", instructions.Width)
	}

	if instructions.Height != 100 {
		t.Fatalf("height is not as expected, got: %d", instructions.Height)
	}
}

func TestParseInstructionsCropWithoutWidth(t *testing.T) {
	queryString := "origin=http%3A%2F%2Fexample.com%2Ftest.jpg" +
		"&action=crop" +
		"&format=png" +
		"&height=200"
	query, err := url.ParseQuery(queryString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ParseInstructions(query)
	if err == nil {
		t.Fatal("instructions should not be valid")
	}
}

func TestParseInstructionsCropWithoutHeight(t *testing.T) {
	queryString := "origin=http%3A%2F%2Fexample.com%2Ftest.jpg" +
		"&action=crop" +
		"&format=png" +
		"&width=200"
	query, err := url.ParseQuery(queryString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ParseInstructions(query)
	if err == nil {
		t.Fatal("instructions should not be valid")
	}
}

func TestParseInstructionsShrink(t *testing.T) {
	queryString := "origin=http%3A%2F%2Fexample.com%2Ftest.jpg" +
		"&action=shrink-w" +
		"&format=png" +
		"&width=200"
	query, err := url.ParseQuery(queryString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ParseInstructions(query)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseInstructionsShrinkWithoutWidth(t *testing.T) {
	queryString := "origin=http%3A%2F%2Fexample.com%2Ftest.jpg" +
		"&action=shrink-w" +
		"&format=png"
	query, err := url.ParseQuery(queryString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ParseInstructions(query)
	if err == nil {
		t.Fatal("instructions should not be valid")
	}
}

func TestParseInstructionsWithoutFormat(t *testing.T) {
	queryString := "origin=http%3A%2F%2Fexample.com%2Ftest.jpg" +
		"&action=crop" +
		"&width=200" +
		"&height=200"
	query, err := url.ParseQuery(queryString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ParseInstructions(query)
	if err == nil {
		t.Fatal("instructions should not be valid")
	}
}

func TestParseInstructionsWithInvalidFormat(t *testing.T) {
	queryString := "origin=http%3A%2F%2Fexample.com%2Ftest.jpg" +
		"&action=crop" +
		"&width=200" +
		"&height=200" +
		"&format=gif"
	query, err := url.ParseQuery(queryString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ParseInstructions(query)
	if err == nil {
		t.Fatal("instructions should not be valid")
	}
}

func TestParseInstructionsWithInvalidOrigin(t *testing.T) {
	queryString := "origin=%2Fimage.png" +
		"&action=crop" +
		"&width=200" +
		"&height=200" +
		"&format=png"
	query, err := url.ParseQuery(queryString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ParseInstructions(query)
	if err == nil {
		t.Fatal("instructions should not be valid")
	}
}
