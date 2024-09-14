+++
title = "{{ .Title }}"
description = "{{ .Description }}"
tags = [{{- range $index, $tag := .Tags }}{{ if $index }}, {{ end }}"{{ $tag }}"{{- end }}]
categories = ["{{ .Categories }}"]
date = {{ .Date }}
draft = false
+++
