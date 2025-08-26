package components

import (
	"strings"
	"testing"
	"time"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

// Helper function to render DatePickerComponent to string
func renderToStringDatePicker(component flyon.Component) string {
	var sb strings.Builder
	component.Render(&sb)
	return sb.String()
}

func TestNewDatePicker(t *testing.T) {
	datePicker := NewDatePicker()

	if datePicker.id != "" {
		t.Errorf("Expected empty ID, got %s", datePicker.id)
	}
	if datePicker.name != "" {
		t.Errorf("Expected empty name, got %s", datePicker.name)
	}
	if datePicker.placeholder != "" {
		t.Errorf("Expected empty placeholder, got %s", datePicker.placeholder)
	}
	if !datePicker.value.IsZero() {
		t.Errorf("Expected zero time value, got %v", datePicker.value)
	}
	if datePicker.disabled != false {
		t.Errorf("Expected disabled to be false, got %t", datePicker.disabled)
	}
	if datePicker.color != flyon.Primary {
		t.Errorf("Expected color to be Primary, got %v", datePicker.color)
	}
	if datePicker.size != flyon.SizeMedium {
		t.Errorf("Expected size to be SizeMedium, got %v", datePicker.size)
	}
	if datePicker.format != "2006-01-02" {
		t.Errorf("Expected format to be '2006-01-02', got %s", datePicker.format)
	}
	if len(datePicker.classes) != 0 {
		t.Errorf("Expected empty classes, got %v", datePicker.classes)
	}
	if len(datePicker.attributes) != 0 {
		t.Errorf("Expected empty attributes, got %v", datePicker.attributes)
	}
}

func TestDatePickerComponent_WithID(t *testing.T) {
	datePicker := NewDatePicker().WithID("test-id")

	if datePicker.id != "test-id" {
		t.Errorf("Expected ID to be 'test-id', got %s", datePicker.id)
	}
}

func TestDatePickerComponent_WithName(t *testing.T) {
	datePicker := NewDatePicker().WithName("test-name")

	if datePicker.name != "test-name" {
		t.Errorf("Expected name to be 'test-name', got %s", datePicker.name)
	}
}

func TestDatePickerComponent_WithPlaceholder(t *testing.T) {
	datePicker := NewDatePicker().WithPlaceholder("Select date...")

	if datePicker.placeholder != "Select date..." {
		t.Errorf("Expected placeholder to be 'Select date...', got %s", datePicker.placeholder)
	}
}

func TestDatePickerComponent_WithValue(t *testing.T) {
	testDate := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	datePicker := NewDatePicker().WithValue(testDate)

	if !datePicker.value.Equal(testDate) {
		t.Errorf("Expected value to be %v, got %v", testDate, datePicker.value)
	}
}

func TestDatePickerComponent_WithDisabled(t *testing.T) {
	datePicker := NewDatePicker().WithDisabled(true)

	if datePicker.disabled != true {
		t.Errorf("Expected disabled to be true, got %t", datePicker.disabled)
	}
}

func TestDatePickerComponent_WithColor(t *testing.T) {
	tests := []struct {
		name     string
		color    flyon.Color
		expected flyon.Color
	}{
		{"Primary color", flyon.Primary, flyon.Primary},
		{"Secondary color", flyon.Secondary, flyon.Secondary},
		{"Success color", flyon.Success, flyon.Success},
		{"Warning color", flyon.Warning, flyon.Warning},
		{"Error color", flyon.Error, flyon.Error},
		{"Info color", flyon.Info, flyon.Info},
		{"Neutral color", flyon.Neutral, flyon.Neutral},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datePicker := NewDatePicker().WithColor(tt.color)
			if datePicker.color != tt.expected {
				t.Errorf("Expected color to be %v, got %v", tt.expected, datePicker.color)
			}
		})
	}
}

func TestDatePickerComponent_WithSize(t *testing.T) {
	tests := []struct {
		name     string
		size     flyon.Size
		expected flyon.Size
	}{
		{"Extra small size", flyon.SizeXS, flyon.SizeXS},
		{"Small size", flyon.SizeSmall, flyon.SizeSmall},
		{"Medium size", flyon.SizeMedium, flyon.SizeMedium},
		{"Large size", flyon.SizeLarge, flyon.SizeLarge},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datePicker := NewDatePicker().WithSize(tt.size)
			if datePicker.size != tt.expected {
				t.Errorf("Expected size to be %v, got %v", tt.expected, datePicker.size)
			}
		})
	}
}

func TestDatePickerComponent_WithFormat(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected string
	}{
		{"ISO format", "2006-01-02", "2006-01-02"},
		{"US format", "01/02/2006", "01/02/2006"},
		{"European format", "02/01/2006", "02/01/2006"},
		{"Long format", "January 2, 2006", "January 2, 2006"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datePicker := NewDatePicker().WithFormat(tt.format)
			if datePicker.format != tt.expected {
				t.Errorf("Expected format to be %s, got %s", tt.expected, datePicker.format)
			}
		})
	}
}

func TestDatePickerComponent_WithMinDate(t *testing.T) {
	minDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	datePicker := NewDatePicker().WithMinDate(minDate)

	if !datePicker.minDate.Equal(minDate) {
		t.Errorf("Expected minDate to be %v, got %v", minDate, datePicker.minDate)
	}
}

func TestDatePickerComponent_WithMaxDate(t *testing.T) {
	maxDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	datePicker := NewDatePicker().WithMaxDate(maxDate)

	if !datePicker.maxDate.Equal(maxDate) {
		t.Errorf("Expected maxDate to be %v, got %v", maxDate, datePicker.maxDate)
	}
}

func TestDatePickerComponent_WithClasses(t *testing.T) {
	datePicker := NewDatePicker().WithClasses("custom-class", "another-class")
	html := renderToStringDatePicker(datePicker)

	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected HTML to contain 'custom-class', got %s", html)
	}
	if !strings.Contains(html, "another-class") {
		t.Errorf("Expected HTML to contain 'another-class', got %s", html)
	}
}

func TestDatePickerComponent_With(t *testing.T) {
	datePicker := NewDatePicker().WithAttribute("data-test", "value")
	html := renderToStringDatePicker(datePicker)

	if !strings.Contains(html, `data-test="value"`) {
		t.Errorf("Expected HTML to contain data-test attribute, got %s", html)
	}
}

func TestDatePickerComponent_Render(t *testing.T) {
	datePicker := NewDatePicker()
	html := renderToStringDatePicker(datePicker)

	if !strings.Contains(html, "input") {
		t.Errorf("Expected HTML to contain input element, got %s", html)
	}
	if !strings.Contains(html, "input-bordered") {
		t.Errorf("Expected HTML to contain input-bordered class, got %s", html)
	}
	if !strings.Contains(html, `type="date"`) {
		t.Errorf("Expected HTML to contain date input type, got %s", html)
	}
}

func TestDatePickerComponent_RenderWithValue(t *testing.T) {
	testDate := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	datePicker := NewDatePicker().WithValue(testDate)
	html := renderToStringDatePicker(datePicker)

	if !strings.Contains(html, "2024-01-15") {
		t.Errorf("Expected HTML to contain formatted date '2024-01-15', got %s", html)
	}
}

func TestDatePickerComponent_RenderWithMinMaxDates(t *testing.T) {
	minDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	maxDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	datePicker := NewDatePicker().WithMinDate(minDate).WithMaxDate(maxDate)
	html := renderToStringDatePicker(datePicker)

	if !strings.Contains(html, `min="2024-01-01"`) {
		t.Errorf("Expected HTML to contain min date attribute, got %s", html)
	}
	if !strings.Contains(html, `max="2024-12-31"`) {
		t.Errorf("Expected HTML to contain max date attribute, got %s", html)
	}
}

func TestDatePickerComponent_RenderWithAllAttributes(t *testing.T) {
	testDate := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	minDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	maxDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)

	datePicker := NewDatePicker().
		WithID("test-id").
		WithName("test-name").
		WithPlaceholder("Select date...").
		WithValue(testDate).
		WithDisabled(true).
		WithColor(flyon.Success).
		WithSize(flyon.SizeLarge).
		WithFormat("01/02/2006").
		WithMinDate(minDate).
		WithMaxDate(maxDate).
		WithClasses("custom-class").
		WithAttribute("data-test", "value")

	html := renderToStringDatePicker(datePicker)

	if !strings.Contains(html, `id="test-id"`) {
		t.Errorf("Expected HTML to contain id attribute, got %s", html)
	}
	if !strings.Contains(html, `name="test-name"`) {
		t.Errorf("Expected HTML to contain name attribute, got %s", html)
	}
	if !strings.Contains(html, "disabled") {
		t.Errorf("Expected HTML to contain disabled attribute, got %s", html)
	}
	if !strings.Contains(html, "input-success") {
		t.Errorf("Expected HTML to contain success color class, got %s", html)
	}
	if !strings.Contains(html, "input-lg") {
		t.Errorf("Expected HTML to contain large size class, got %s", html)
	}
	if !strings.Contains(html, "custom-class") {
		t.Errorf("Expected HTML to contain custom class, got %s", html)
	}
	if !strings.Contains(html, `data-test="value"`) {
		t.Errorf("Expected HTML to contain data-test attribute, got %s", html)
	}
	if !strings.Contains(html, `min="2024-01-01"`) {
		t.Errorf("Expected HTML to contain min date, got %s", html)
	}
	if !strings.Contains(html, `max="2024-12-31"`) {
		t.Errorf("Expected HTML to contain max date, got %s", html)
	}
}

func TestDatePickerComponent_Immutability(t *testing.T) {
	original := NewDatePicker()
	modified := original.WithID("new-id")

	if original.id == modified.id {
		t.Error("Expected original component to remain unchanged")
	}
	if original.id != "" {
		t.Errorf("Expected original ID to be empty, got %s", original.id)
	}
	if modified.id != "new-id" {
		t.Errorf("Expected modified ID to be 'new-id', got %s", modified.id)
	}
}