/* defines charts element*/
package render

var EchartsCommitActivityHTML = `
<div class="item" id="{{ .ChartID }}" style="width:{{ .Initialization.Width }};height:{{ .Initialization.Height }};min-height:15rem;"></div>
`
var EchartCommittersHTML = `
<div class="item" id="{{ .ChartID }}" style="width:{{ .Initialization.Width }};height:{{ .Initialization.Height }}; max-height:20rem;"></div>
`
var EchartLangHTML = `
<div class="item" id="{{ .ChartID }}" style="width:{{ .Initialization.Width }};height:{{ .Initialization.Height }}; max-height:25rem;"></div>
`
var EchartIssueHTML = `
<div class="item" id="{{ .ChartID }}" style="width:{{ .Initialization.Width }};height:{{ .Initialization.Height }}; max-height:25rem;"></div>
`

var EchartsCommonJS = `
			{{- range .JSAssets.Values }}
			<script src="{{ . }}"></script>
			{{- end }}
			<script type="text/javascript">
					"use strict";
					let goecharts_{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID | safeJS }}'), "{{ .Theme }}");
					let option_{{ .ChartID | safeJS }} = {{ .JSON }};
					goecharts_{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});
					{{- range .JSFunctions.Fns }}
					{{ . | safeJS }}
					{{- end }}
			</script>
`
