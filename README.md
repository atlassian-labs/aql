# AQL

[![Atlassian license](https://img.shields.io/badge/license-Apache%202.0-blue.svg?style=flat-square)](LICENSE) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](CONTRIBUTING.md)

AQL provides a helperful query language to write Atlassian Query Languages in
various formats for different products. We use this library to query Bitbucket,
Jira, Confluence and more.

[![Go Reference](https://pkg.go.dev/badge/github.com/atlassian-labs/aql.svg)](https://pkg.go.dev/github.com/atlassian-labs/aql)

## Usage

Writing raw queries and sending them to various Atlassian products can get
complex and harder to deal with the more you want to query or the more complex
you wanna query a product.

Using AQL you can write out various queries like

```go
jql := aql.NewStringQuery()

jql.Contains("comment", "Charlie").And().NotEqual("Reporter", "Charles Atlas").
    And().GtrEqualTo("created", "2023-01-01").
    And().LessEqualTo("created", "2023-03-01").
    And().NotEqual("resolution", aql.Literal("null"))
```

Which would translate too
```
comment ~ Charlie AND != Reporter "Charles Atlas" AND >= 2023-01-01 AND <= 2023-03-01
    AND != resolution == null
```

An example of using this via a HTTP endpoint would be something like
```go
jql := aql.NewStringQuery()

jql.Contains("comment", "Charlie").And().NotEqual("Reporter", "Charles Atlas").
    And().GtrEqualTo("created", "2023-01-01").
    And().LessEqualTo("created", "2023-03-01").
    And().NotEqual("resolution", aql.Literal("null"))

type Query struct {
    JQL string `json:"jql"`
}

var body = &bytes.Buffer{}

err := json.NewEncoder(body).Encode(&Query{
    JQL: jql.Encode(),
})

req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, "https://yourjira.atlassian.net/rest/api/3/search", body)
if err != nil {
    panic(err)
}

```


## Installation

`go get github.com/atlassian-labs/aql@latest`

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/atlassian-labs/aql.svg)](https://pkg.go.dev/github.com/atlassian-labs/aql)

## Tests

`make test`

## Contributions

Contributions to AQL are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

Copyright (c) 2023 Atlassian US., Inc.
Apache 2.0 licensed, see [LICENSE](LICENSE) file.

<br/>

[![With ❤️ from Atlassian](https://raw.githubusercontent.com/atlassian-internal/oss-assets/master/banner-with-thanks.png)](https://www.atlassian.com)
