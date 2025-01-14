{{/* vim: set filetype=mustache: */}}

{{/*
**********************
kubeforge - Naming 
**********************
*/}}

{{- define "kubeforge.name" -}}
    {{- default .Chart.Name .Values.kubeforge.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "kubeforge.fullname" -}}
    {{- include "library.componentFullname" (dict
        "componentName" "kubeforge"
        "componentValues" .Values.kubeforge
        "context" $
    ) -}}
{{- end -}}

{{/*
**********************
kubeforge - Labels
**********************
*/}}

{{/*
Defines extra labels for optimize.
*/}}
{{- define "kubeforge.extraLables" -}}
app.kubernetes.io/component: kubeforge 
{{- end -}}

{{/*
Define common labels, combining the match labels and transient labels, which might change on updating
(version depending). These labels should not be used on matchLabels selector, since the selectors are immutable.
*/}}
{{- define "kubeforge.labels" -}}
{{- template "library.labels" . }}
{{ template "kubeforge.extraLables" . }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "kubeforge.matchLabels" -}}
{{- template "library.matchLabels" . }}
app.kubernetes.io/component: kubeforge 
{{- end -}}

