package templmanager

const mainTmpl = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <link rel="icon" type="image/x-icon" href="/static/images/favicon.ico">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta charset="UTF-8">
        <title>{{if .Abbreviation}}{{.Abbreviation}}.{{else}}read-only.ru{{ end }} {{if .Paragraphs}}{{if .Num}}{{.Num}}.  {{.Name}} {{else}}{{.Name}}{{ end }}{{end}}</title>
        <link rel="stylesheet" href="/static/css/styles.css">
        <style>
            {{template "css" .}}
        </style>
    </head>
    <body>
       {{template "body" .}}
     </body>
</html>
`
