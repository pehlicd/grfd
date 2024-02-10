# {{ .Date }}

## Daily
Last {{ .LastXDays }} day you have worked on {{ .IssueCount }} different issues.

{{range $val := .Issues}}
- [{{ $val.Title }}]({{ $val.URL }})
  - Description: {{ $val.Description }}
  - Status: {{ $val.Status }}
  - Last Updated: {{ $val.LastUpdatedAt }}
{{end}}