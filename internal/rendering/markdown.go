// Package rendering converts a subset of Markdown into GTK Pango markup.
//
// Supported syntax covers the kata README format:
//   - # Heading, ## Subheading
//   - **bold**
//   - `inline code`
//   - ``` fenced code blocks
//   - - bullet items
//   - --- horizontal rules
//   - bare URLs → wrapped in an anchor tag
//
// The output is safe for textbuffer.InsertMarkup().
package rendering

import (
	"html"
	"regexp"
	"strings"
)

var (
	reHeading1   = regexp.MustCompile(`^#{1}\s+(.+)$`)
	reHeading2   = regexp.MustCompile(`^#{2,6}\s+(.+)$`)
	reBold       = regexp.MustCompile(`\*\*(.+?)\*\*`)
	reInlineCode = regexp.MustCompile("`(.*?)`")
	reBullet     = regexp.MustCompile(`^\s*[-*]\s+(.+)$`)
	reHRule      = regexp.MustCompile(`^---+\s*$`)
	reFenceStart = regexp.MustCompile("^```")
	reURL        = regexp.MustCompile(`https?://[^\s<>"']+`)
)

// MarkdownToPango converts markdown text into Pango markup suitable for
// display in a GTK4 TextBuffer using InsertMarkup.
func MarkdownToPango(md string) string {
	lines := strings.Split(md, "\n")
	var out strings.Builder
	inCodeBlock := false
	inBulletList := false

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// Fenced code block toggle
		if reFenceStart.MatchString(line) {
			if inCodeBlock {
				inCodeBlock = false
				out.WriteString("</tt>\n")
			} else {
				inCodeBlock = true
				out.WriteString(`<tt>`)
				// Close any open bullet list
				if inBulletList {
					inBulletList = false
				}
			}
			continue
		}

		if inCodeBlock {
			out.WriteString(html.EscapeString(line))
			out.WriteString("\n")
			continue
		}

		// Close bullet list if we hit a non-bullet line
		if inBulletList && !reBullet.MatchString(line) && strings.TrimSpace(line) != "" {
			inBulletList = false
		}

		// Horizontal rule
		if reHRule.MatchString(line) {
			out.WriteString(`<span foreground="#6b7686">─────────────────────────────────────────────────────</span>` + "\n")
			continue
		}

		// Heading 1
		if m := reHeading1.FindStringSubmatch(line); m != nil {
			out.WriteString(`<span size="x-large" weight="800">` + inlinePango(m[1]) + `</span>` + "\n")
			continue
		}

		// Heading 2+
		if m := reHeading2.FindStringSubmatch(line); m != nil {
			out.WriteString(`<span size="large" weight="700">` + inlinePango(m[1]) + `</span>` + "\n")
			continue
		}

		// Bullet item
		if m := reBullet.FindStringSubmatch(line); m != nil {
			if !inBulletList {
				inBulletList = true
			}
			out.WriteString(`  • ` + inlinePango(m[1]) + "\n")
			continue
		}

		// Empty line
		if strings.TrimSpace(line) == "" {
			out.WriteString("\n")
			continue
		}

		// Regular line
		out.WriteString(inlinePango(line))
		out.WriteString("\n")
	}

	// Close unclosed code block
	if inCodeBlock {
		out.WriteString("</tt>\n")
	}

	return strings.TrimRight(out.String(), "\n")
}

// inlinePango applies inline formatting (bold, inline code, URLs) to a line.
func inlinePango(line string) string {
	// Escape first so README prose such as `<Name>` and `a & b` can never be
	// interpreted as Pango markup. Formatting is then applied to the escaped
	// representation.
	line = html.EscapeString(line)
	line = reBold.ReplaceAllStringFunc(line, func(match string) string {
		content := match[2 : len(match)-2]
		return `<b>` + content + `</b>`
	})

	line = reInlineCode.ReplaceAllStringFunc(line, func(match string) string {
		content := match[1 : len(match)-1]
		return `<tt foreground="#14b8a6">` + content + `</tt>`
	})

	line = reURL.ReplaceAllStringFunc(line, func(match string) string {
		return `<span foreground="#14b8a6">` + match + `</span>`
	})

	return line
}
