module cli-color-calc

go 1.25.3

replace pkg/repl => ./pkg/repl

replace pkg/colorspace => ./pkg/colorspace

require pkg/repl v0.0.0-00010101000000-000000000000

require pkg/colorspace v0.0.0-00010101000000-000000000000 // indirect
