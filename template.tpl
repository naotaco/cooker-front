<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ page.Title }}</title>
    <script language="JavaScript"><!--
				      function update(){
				      location.reload();
				      }
				      // -->
    </script>
</head>
<body onLoad="setInterval('update()', 5000)">
    <div id="content">
        <h1 class="title">{{ page.Title }}</h1>
        <div class="body">
		  <h2>Current status</h2>
		  <h3>Set point</h3>
		  <p>{{page.TargetTemperature}}</p>
		  <h3>Current Temperature</h3>
		  <p>{{page.CurrentTemperature}}</p>
		  <h3>Is heating</h3>
		  <p>{{page.IsHeating}}</p>

		  <h2>Setting</h2>
		  <p>
			<form method="post" action="/set" >
			  <input type="text" name="target" value=""/>
			  <input type="submit" name="Submit" value="Set target" />
			</form>
		  </p>

		  <hr />
		  <p>
			<form method="post" action="/force_on" >
			  <input type="submit" name="Submit" value="Force ON" />
			</form>
		  </p>
		  <p>
			<form method="post" action="/force_off" >
			  <input type="submit" name="Submit" value="Force OFF" />
			</form>
		  </p>
		  
        </div>
        <br />
    </div>
</body>
</html>
