# Maps
learning how maps works in Golang

### String & Slice Utilities

A collection of Go helpers for parsing strings into structured maps and slices:


| Function | Input | Output | Description |
| :--- | :--- | :--- | :--- |
| **MapCSVToIndex** | `string` | `map[int]string` | Converts comma-separated values into an index-mapped collection. |
| **MapWordsToIndex** | `string` | `map[int]string` | Splits text by whitespace and maps each word to its position. |
| **MapWordLastOccurrence** | `string` | `map[string]int` | Tracks unique words and stores the index of their most recent occurrence. |
| **ChunkWordsIntoTriplets** | `string` | `[][]string` | Segments text into a grid-like slice with 3 words per group. |

**Usage Example:**
```go
// Groups text into chunks of three words
chunks := ChunkWordsIntoTriplets("The quick brown fox jumps over") 
// Output: [[The quick brown] [fox jumps over]]
```

