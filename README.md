# Humanize

Go package to convert machine-readables to human-readables.


## Commas

```go
AddCommas(n, decimals int) string
```

Returns a string-representation of the given int with commas inserted appropriately.

```
AddCommas(1234) => 1,234
```

## Numbers

```go
Number(n, decimals int) string
```

Simplify the given int to a human-friendly abbreviation with a specified number of decimal places

```
Number(1234, 2) => 1.23k
```

```go
Size(n, decimals int) string
```

Simplify the given int, which is assumed to be a byte count, to a human-friendly abbreviation with a specified number of decimal places. Uses 1024 instead of 1000 as the base.

```
Size(1024, 2) => 1.00K
```

## Tests

```go
IsASCII(str string, ratio float32) bool
```

Tests the given string for ASCII-wise readability, i.e. 32 <= c < 127 and compares the count of unreadable chars against the given ratio threshold.

## Hex

```go
HexDump(slice []byte) string
```

Renders a byte-slice to a string with hex-editor-like formatting.

e.g.

```
00000  74 3E C7 45 DE 41 E1 EE  2C 26 26 C4 C2 9E BC 4F  t>.E.A..  ,&&....O
00016  00 70 5A BA 0C 39 7E 86  40 22 20 31 A5 2E BC 69  .pZ..9~.  @" 1...i
00032  90 BF DE 9D 55 65 7E 58  71 40 7A AD 20 76 2E 73  ....Ue~X  q@z. v.s
00048  C8 AE 70 4B 86 E8 B2 3D  A3 BF 27 78 77 53 5C 14  ..pK...=  ..'xwS\.
00064  3E 07 FF 05 16 1C 60 54  81 E0 3E 62 FB CD 3A B4  >.....`T  ..>b..:.
00080  B6 78 7B 9B D5 EE 68 EE  09 27 DE 73 22 AC 09 FF  .x{...h.  .'.s"...
00096  45 F1 1D A5                                       E...
```


# License

http://unlicense.org