package utils

import "testing"
import "reflect"

func expectMessageEqual(t *testing.T, message1 string, message2 string) {
	if message1 != message2 {
		t.Errorf("got %q, wanted %q", message1, message2)
	}
}

func expectFormatStringsEqual(t *testing.T, formatStrings1 []string, formatStrings2 []string) {
	if !reflect.DeepEqual(formatStrings1, formatStrings2) {
		t.Errorf("got %q, wanted %q", formatStrings1, formatStrings2)
	}
}

func TestSimpleMessage1(t *testing.T) {
    message, signalCliFormatStrings := ParseMarkdownMessage("*italic*")
	expectMessageEqual(t, message, "italic")
	expectFormatStringsEqual(t, signalCliFormatStrings, []string{"0:6:ITALIC"})
}

func TestSimpleMessage(t *testing.T) {
    message, signalCliFormatStrings := ParseMarkdownMessage("*This is a italic message*")
	expectMessageEqual(t, message, "This is a italic message")
	expectFormatStringsEqual(t, signalCliFormatStrings, []string{"0:24:ITALIC"})
}

func TestBoldAndItalicMessage(t *testing.T) {
    message, signalCliFormatStrings := ParseMarkdownMessage("This is a **bold** and *italic* message")
	expectMessageEqual(t, message, "This is a bold and italic message")
	expectFormatStringsEqual(t, signalCliFormatStrings, []string{"10:4:BOLD", "19:6:ITALIC"})
}

func TestTwoBoldFormattedStrings(t *testing.T) {
	message, signalCliFormatStrings := ParseMarkdownMessage("This is a **bold** and another **bold** message")
	expectMessageEqual(t, message, "This is a bold and another bold message")
	expectFormatStringsEqual(t, signalCliFormatStrings, []string{"10:4:BOLD", "27:4:BOLD"})
}

func TestStrikethrough(t *testing.T) {
	message, signalCliFormatStrings := ParseMarkdownMessage("This is a ~strikethrough~ and a **bold** message")
	expectMessageEqual(t, message, "This is a strikethrough and a bold message")
	expectFormatStringsEqual(t, signalCliFormatStrings, []string{"10:13:STRIKETHROUGH", "30:4:BOLD"})
}

func TestMonospace(t *testing.T) {
	message, signalCliFormatStrings := ParseMarkdownMessage("This is a `monospace` and a **bold** message")
	expectMessageEqual(t, message, "This is a monospace and a bold message")
	expectFormatStringsEqual(t, signalCliFormatStrings, []string{"10:9:MONOSPACE", "26:4:BOLD"})
}
