<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>TravelSphere - Discover Your Next Destination</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
    {{template "partials/header.tpl" .}}
    
    <main class="content-wrapper">
        {{.LayoutContent}}
    </main>

    {{template "partials/footer.tpl" .}}
    <script src="/static/js/autocomplete.js"></script>
</body>
</html>