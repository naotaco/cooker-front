<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ page.Title }}</title>
    <link rel="stylesheet" href="/assets/css/style.css" media="all">
</head>
<body>
    <div id="content">
        <h1 class="title">{{ page.Title }}</h1>
        <div class="body">
		  <h3>Set point</h3>
		  <p>{{page.TargetTemperature}}</p>
		  <h3>Current Temperature</h3>
		  <p>{{page.CurrentTemperature}}</p>
		  <h3>Is heating</h3>
		  <p>{{page.IsHeating}}</p>
        </div>
        <br />
    </div>
</body>
</html>
