/*
Package languages is the editor's language system.

A Language is a single self-contained definition that drives everything the
editor does with a source file:

  - Syntax highlighting (keywords, types, strings, comments, numbers, calls)
  - Auto-pairing and smart indentation (braces, brackets, quotes)
  - Syntax error indication (a Checker reports diagnostics as squiggles)
  - Workspace filenames (solution / learner tests per language)

Adding a language is a pure registration — no editor code changes:

  - Implement a Checker (optional): parse a lint/compile tool's output into
    []Diagnostic so the editor can underline errors.
  - Add the Language to builtins() in builtins.go (or Register it at startup).

Frameworks extend a language the same way: register a Language for the
framework's dialect (e.g. "typescript+react") reusing the base keywords, or
register a new extension that maps to an existing checker. The registry is
the single extension point; everything else keys off the Language value.

To add a new built-in language:

 1. Fill in keywords, types, comment syntax, quote chars, auto-pairs, indent.
 2. Give it a Checker if a syntax checker exists for it.
 3. Add it to builtins().

Katase declare their language via the "language" field in kata.json metadata;
an empty value means Go.
*/
package languages
