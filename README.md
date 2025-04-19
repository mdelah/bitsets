# bitsets

[![PkgGoDev](https://pkg.go.dev/badge/github.com/mdelah/bitset?tab=doc)](https://pkg.go.dev/github.com/mdelah/bitset?tab=doc)

### Description

This small package provides bitset types for Go, leaning on `math/bits` under the hood. Example use cases might be
marking certain indices of a small slice, or keeping a selection of values from an enumeration.

There are a few similar other packages out there, this particular one has the following feature set:

- Friendly to stack allocation, zero value is good to go
- Separate types for fixed-width versions of capacity 8, 16, 32, 64 (`bit8` through `bit64`), 1024 and 8192 (`kbit` and `kbit8`), plus a variable-width one (`vbit`)
- Fluent methods for usual logic operations (`Not`, `And`, `Or`, etc), each with a self-assigning version (`AssignNot`, `AssignAnd`, `AssignOr`, etc)
- Loop over the values (or contiguous sub-ranges) with range-func
- No dependencies

The various types provide a similar interface, so you can swap between with relative ease if needs change.

### License

MIT
