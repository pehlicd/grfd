# {{ .Date }}

## Issues
Last {{ .LastXDays }} day you have worked on {{ .IssueCount }} different issues.
{{range $val := .Issues}}
- [{{ $val.Title }}]({{ $val.URL }})
  - Description: {{ $val.Description }}
  - Status: {{ $val.Status }}
  - Last Updated: {{ $val.LastUpdatedAt }}
{{end}}

## Approved Merge Requests
Last {{ .LastXDays }} day {{ .GitlabUsername }} has approved {{ .ApprovedMRsCount }} different merge requests.
{{range $val := .ApprovedMRs}}
- [{{ $val.Title }}]({{ $val.URL }})
  - Description: {{ $val.Description }}
  - Status: {{ $val.Status }}
  - Last Updated: {{ $val.UpdatedAt }}
{{end}}

## Merged Merge Requests
Last {{ .LastXDays }} day you have merged {{ .MergedMRsCount }} different merge requests merged by {{ .GitlabUsername }}.
{{range $val := .MergedMRs}}
- [{{ $val.Title }}]({{ $val.URL }})
  - Description: {{ $val.Description }}
  - Status: {{ $val.Status }}
  - Merged At: {{ $val.MergedAt }}
{{end}}