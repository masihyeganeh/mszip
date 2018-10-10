# MSZip Decompressor

Decompresses MSZip compressed contents

## How to use

```go
import (
    "github.com/masihyeganeh/mszip"
)

compressed := bytes.NewReader([]byte("COMPRESSED DATA"))
decompressedSize := 1234

decompressor := mszip.New()
decompressed, err := decompressor.Decompress(compressed, decompressedSize)
```