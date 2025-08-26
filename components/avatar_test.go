package components

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/ozanturksever/gomponents-flyonui/flyon"
)

func TestAvatar_BasicRendering(t *testing.T) {
	t.Run("renders basic avatar with default classes", func(t *testing.T) {
		avatar := NewAvatar(g.Text("JD"))
		html := renderToHTML(avatar)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		avatarEl := findElement(doc, "div")
		if avatarEl == nil {
			t.Fatal("Expected div element not found")
		}
		
		classAttr := getAttribute(avatarEl, "class")
		if !hasClass(classAttr, "avatar") {
			t.Errorf("Expected 'avatar' class, got classes: %s", classAttr)
		}
		
		if !strings.Contains(html, "JD") {
			t.Errorf("Expected 'JD' text, got: %s", html)
		}
	})
	
	t.Run("renders avatar with image", func(t *testing.T) {
		avatar := NewAvatar(
			h.Img(h.Src("https://example.com/avatar.jpg"), h.Alt("User Avatar")),
		)
		html := renderToHTML(avatar)
		
		if !strings.Contains(html, `<img src="https://example.com/avatar.jpg" alt="User Avatar">`) {
			t.Errorf("Expected img element with src and alt, got: %s", html)
		}
	})
	
	t.Run("renders avatar with placeholder", func(t *testing.T) {
		avatar := NewAvatar(
			h.Div(
				h.Class("bg-neutral text-neutral-content"),
				h.Span(g.Text("JD")),
			),
		)
		html := renderToHTML(avatar)
		
		if !strings.Contains(html, "JD") {
			t.Errorf("Expected 'JD' text in placeholder, got: %s", html)
		}
	})
}

func TestAvatar_ColorModifiers(t *testing.T) {
	tests := []struct {
		name          string
		color         flyon.Color
		expectedClass string
	}{
		{"Primary avatar", flyon.Primary, "avatar-primary"},
		{"Secondary avatar", flyon.Secondary, "avatar-secondary"},
		{"Success avatar", flyon.Success, "avatar-success"},
		{"Warning avatar", flyon.Warning, "avatar-warning"},
		{"Error avatar", flyon.Error, "avatar-error"},
		{"Info avatar", flyon.Info, "avatar-info"},
		{"Neutral avatar", flyon.Neutral, "avatar-neutral"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avatar := NewAvatar(g.Text("JD")).With(tt.color)
			html := renderToHTML(avatar)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			avatarEl := findElement(doc, "div")
			if avatarEl == nil {
				t.Fatal("Expected div element not found")
			}
			
			classAttr := getAttribute(avatarEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestAvatar_SizeModifiers(t *testing.T) {
	tests := []struct {
		name          string
		size          flyon.Size
		expectedClass string
	}{
		{"Extra small avatar", flyon.SizeXS, "avatar-xs"},
		{"Small avatar", flyon.SizeSmall, "avatar-sm"},
		{"Medium avatar", flyon.SizeMedium, "avatar-md"},
		{"Large avatar", flyon.SizeLarge, "avatar-lg"},
		{"Extra large avatar", flyon.SizeXL, "avatar-xl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avatar := NewAvatar(g.Text("JD")).With(tt.size)
			html := renderToHTML(avatar)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			avatarEl := findElement(doc, "div")
			if avatarEl == nil {
				t.Fatal("Expected div element not found")
			}
			
			classAttr := getAttribute(avatarEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestAvatar_VariantModifiers(t *testing.T) {
	tests := []struct {
		name          string
		variant       flyon.Variant
		expectedClass string
	}{
		{"Solid avatar", flyon.VariantSolid, "avatar-solid"},
		{"Outline avatar", flyon.VariantOutline, "avatar-outline"},
		{"Ghost avatar", flyon.VariantGhost, "avatar-ghost"},
		{"Soft avatar", flyon.VariantSoft, "avatar-soft"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avatar := NewAvatar(g.Text("JD")).With(tt.variant)
			html := renderToHTML(avatar)
			
			doc, err := parseHTML(html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}
			
			avatarEl := findElement(doc, "div")
			if avatarEl == nil {
				t.Fatal("Expected div element not found")
			}
			
			classAttr := getAttribute(avatarEl, "class")
			if !hasClass(classAttr, tt.expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", tt.expectedClass, classAttr)
			}
		})
	}
}

func TestAvatar_CombinedModifiers(t *testing.T) {
	t.Run("combines color, size, and variant modifiers", func(t *testing.T) {
		avatar := NewAvatar(g.Text("JD")).With(
			flyon.Primary,
			flyon.SizeLarge,
			flyon.VariantOutline,
		)
		html := renderToHTML(avatar)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		avatarEl := findElement(doc, "div")
		if avatarEl == nil {
			t.Fatal("Expected div element not found")
		}
		
		classAttr := getAttribute(avatarEl, "class")
		expectedClasses := []string{"avatar", "avatar-primary", "avatar-lg", "avatar-outline"}
		
		for _, expectedClass := range expectedClasses {
			if !hasClass(classAttr, expectedClass) {
				t.Errorf("Expected '%s' class, got classes: %s", expectedClass, classAttr)
			}
		}
	})
}

func TestAvatar_HTMLAttributes(t *testing.T) {
	t.Run("accepts HTML attributes", func(t *testing.T) {
		avatar := NewAvatar(
			h.ID("user-avatar"),
			h.DataAttr("testid", "avatar-component"),
			h.Title("User Profile Picture"),
			g.Text("JD"),
		)
		html := renderToHTML(avatar)
		
		doc, err := parseHTML(html)
		if err != nil {
			t.Fatalf("Failed to parse HTML: %v", err)
		}
		
		avatarEl := findElement(doc, "div")
		if avatarEl == nil {
			t.Fatal("Expected div element not found")
		}
		
		if id := getAttribute(avatarEl, "id"); id != "user-avatar" {
			t.Errorf("Expected id='user-avatar', got id='%s'", id)
		}
		
		if testid := getAttribute(avatarEl, "data-testid"); testid != "avatar-component" {
			t.Errorf("Expected data-testid='avatar-component', got data-testid='%s'", testid)
		}
		
		if title := getAttribute(avatarEl, "title"); title != "User Profile Picture" {
			t.Errorf("Expected title='User Profile Picture', got title='%s'", title)
		}
	})
}

func TestAvatar_ComponentInterface(t *testing.T) {
	t.Run("implements Component interface", func(t *testing.T) {
		avatar := NewAvatar(g.Text("JD"))
		
		// Test that it implements flyon.Component
		var _ flyon.Component = avatar
		
		// Test that it implements gomponents.Node
		var _ g.Node = avatar
		
		// Test that With returns a new instance
		newAvatar := avatar.With(flyon.Primary)
		if avatar == newAvatar {
			t.Error("With() should return a new instance, not modify the original")
		}
	})
}