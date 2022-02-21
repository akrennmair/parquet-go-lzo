# parquet-go-lzo

<p align="center">
<a href="https://pkg.go.dev/github.com/akrennmair/parquet-go-lzo"><img src="https://pkg.go.dev/badge/github.com/akrennmair/parquet-go-lzo.svg" alt="Go Reference"></a>
<a href="https://github.com/akrennmair/parquet-go-lzo/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-Apache%202-blue"></a>
</p>

This library implements the LZO compression algorithm for [github.com/fraugster/parquet-go](github.com/fraugster/parquet-go). By default,
`parquet-go` library only supports `GZIP` and `SNAPPY` as compression algorithms to minimize the list
of dependencies.

All you need to do is import this package into your program and the compression
algorithm will be automatically available in `parquet-go`.

```go
import (
    _ "github.com/akrennmair/parquet-go-lzo" // registers the LZO block compressor with parquet-go
)
```

## License

See the file `LICENSE` for further license information.

Please note that this package is built using `github.com/cyberdelia/lzo` which in turn
uses the original [LZO implementation which is licensed as GPLv2+](http://www.oberhumer.com/opensource/lzo/). Please be aware of the licensing implication this can have if you intend to use this in closed-source products that you intend to distribute.

## Author

Andreas Krennmair <ak@synflood.at>
