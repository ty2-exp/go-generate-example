// DO NOT EDIT. CODE GENERATED AUTOMATICALLY
package main

import "fmt"

func main(){
{{- range $index, $fn := .}}
    {{ $fn.Name }}()
{{- end }}
}
{{ range $index, $fn := . }}
func {{ $fn.Name }}(){
    fmt.Println("{{ $fn.Name }} value =", "{{ $fn.Value }}")
}
{{ end }}

