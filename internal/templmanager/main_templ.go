package templmanager

const mainTmpl = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <link rel="icon" type="image/x-icon" href="/static/images/favicon.ico">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta http-equiv="Content-Type" content="type; charset= "/>
        <meta name="description" content="{{.Description}}"/>
        <meta name="keywords" content="{{.Keywords}}"/>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8"><meta charset="windows-1251">
        <title>{{.Title}}</title>
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
