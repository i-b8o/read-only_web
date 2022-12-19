package templmanager

const mainTmpl = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <link rel="icon" type="image/x-icon" href="/static/images/favicon.ico">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta charset="UTF-8">
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
