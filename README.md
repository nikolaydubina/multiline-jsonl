> This package was build for programmatic access of multiline JSON in Go.  
> If you need CLI for JSON, I highly recommend `jq`.

[![Go Reference](https://pkg.go.dev/badge/github.com/nikolaydubina/multiline-jsonl.svg)](https://pkg.go.dev/github.com/nikolaydubina/multiline-jsonl)

```bash
$ go install github.com/nikolaydubina/multiline-jsonl@latest
```

For example, you want to parse input multiline JSONs
```bash
$ echo '{
    "from":"github.com/nikolaydubina/jsonl-graph/graph",
    "to":"bufio"
}

{"from":"my-id-1","to":"my-id-2"}
{"from":"my-id-5","to":"my-id-10", "amount": 123}
{"from":"my-id-5","to":"my-id-10", "amount": {"amount": 123, "currency": "KRW"}}

{
    "id": "my-id",
    "number": 123,
    "nested": {
        "title": "big title",
        "nested-level-2": {
            "subtitle": "some other thing",
            "count": 123
        }
    }
}

' | ./multiline-jsonl
```

Outputs shortened version
```jsonl
{"from":"github.com/nikolaydubina/jsonl-graph/graph","to":"bufio"}
{"from":"my-id-1","to":"my-id-2"}
{"amount":123,"from":"my-id-5","to":"my-id-10"}
{"amount":{"amount":123,"currency":"KRW"},"from":"my-id-5","to":"my-id-10"}
{"id":"my-id","nested":{"nested-level-2":{"count":123,"subtitle":"some other thing"},"title":"big title"},"number":123}
```

And with `-expand` flag
```json
{
    "from": "github.com/nikolaydubina/jsonl-graph/graph",
    "to": "bufio"
}
{
    "from": "my-id-1",
    "to": "my-id-2"
}
{
    "amount": 123,
    "from": "my-id-5",
    "to": "my-id-10"
}
{
    "amount": {
        "amount": 123,
        "currency": "KRW"
    },
    "from": "my-id-5",
    "to": "my-id-10"
}
{
    "id": "my-id",
    "nested": {
        "nested-level-2": {
            "count": 123,
            "subtitle": "some other thing"
        },
        "title": "big title"
    },
    "number": 123
}
```

Here is example from https://github.com/nikolaydubina/jsonl-graph

```go
func NewGraphFromJSONL(r io.Reader) (Graph, error) {
	g := NewGraph()

	scanner := bufio.NewScanner(r)
	scanner.Split(multilinejsonl.SplitMultilineJSONL)

	for scanner.Scan() {
		decoder := json.NewDecoder(bytes.NewReader(scanner.Bytes()))
		decoder.UseNumber()

		var nodeEdge orNodeDataEdgeData
		if err := decoder.Decode(&nodeEdge); err != nil {
			continue
		}

		node, edge, err := nodeEdge.cast()
		if err != nil {
			return g, fmt.Errorf("can not cast: %w", err)
		}

		switch {
		case node != nil:
			g.AddNode(*node)
		case edge != nil:
			g.AddEdge(*edge)
		}
	}

	return g, scanner.Err()
}
```

## Features

- [x] No reflection
- [x] Simple Code
- [x] CLI
- [x] 84% coverage

## Reference:
- https://github.com/wlredeye/jsonlines - reflection, no scanner.Split, no multiline
- https://github.com/neilotoole/sq - no scanner.Split, no multiline
- https://github.com/emersion/go-jsonld - custom json tags, no multiline
- https://github.com/qiangyt/jsonlines2json - no scanner, no multiline
- https://github.com/tylerstillwater/jsonl - just a wrapper
- https://github.com/Meromen/JsonlParser - no multiline
- https://github.com/go-ap/jsonld - no multiline
- https://github.com/youpy/go-jsonl - no multiline
- https://github.com/aaronland/go-jsonl - just a wrapper
