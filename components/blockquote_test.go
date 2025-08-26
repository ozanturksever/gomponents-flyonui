package components

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestBlockquote_Basic(t *testing.T) {
	quote := Blockquote(g.Text("This is a test quote"))
	result := renderToHTML(quote)

	if !strings.Contains(result, "<blockquote") {
		t.Error("Expected blockquote element")
	}
	if !strings.Contains(result, "This is a test quote") {
		t.Error("Expected quote content")
	}
	if !strings.Contains(result, "blockquote") {
		t.Error("Expected blockquote class")
	}
}

func TestBlockquote_WithAuthor(t *testing.T) {
	quote := Blockquote(g.Text("Test quote")).WithAuthor("John Doe")
	result := renderToHTML(quote)

	if !strings.Contains(result, "Test quote") {
		t.Error("Expected quote content")
	}
	if !strings.Contains(result, "John Doe") {
		t.Error("Expected author name")
	}
	if !strings.Contains(result, "<cite") {
		t.Error("Expected cite element for author")
	}
}

func TestBlockquote_WithModifiers(t *testing.T) {
	quote := Blockquote(g.Text("Test quote")).With("custom-class", h.ID("test-id"))
	result := renderToHTML(quote)

	if !strings.Contains(result, "custom-class") {
		t.Error("Expected custom class")
	}
	if !strings.Contains(result, `id="test-id"`) {
		t.Error("Expected test ID")
	}
}

func TestBlockquote_Empty(t *testing.T) {
	quote := Blockquote(g.Text(""))
	result := renderToHTML(quote)

	if !strings.Contains(result, "<blockquote") {
		t.Error("Expected blockquote element even with empty content")
	}
}

func TestBlockquote_MultipleChildren(t *testing.T) {
	quote := Blockquote(
		h.P(g.Text("First paragraph")),
		h.P(g.Text("Second paragraph")),
	)
	result := renderToHTML(quote)

	if !strings.Contains(result, "First paragraph") {
		t.Error("Expected first paragraph")
	}
	if !strings.Contains(result, "Second paragraph") {
		t.Error("Expected second paragraph")
	}
	if strings.Count(result, "<p>") != 2 {
		t.Error("Expected two paragraph elements")
	}
}

func TestBlockquote_WithAuthorAndSource(t *testing.T) {
	quote := Blockquote(g.Text("Test quote")).WithAuthor("John Doe").WithSource("Famous Book")
	result := renderToHTML(quote)

	if !strings.Contains(result, "John Doe") {
		t.Error("Expected author name")
	}
	if !strings.Contains(result, "Famous Book") {
		t.Error("Expected source")
	}
	if !strings.Contains(result, "<cite") {
		t.Error("Expected cite element")
	}
}