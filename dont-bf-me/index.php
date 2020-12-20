<?php
require_once "constant.php";
?>
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">
    <script src="https://www.google.com/recaptcha/api.js?render=<?php echo $PUB_KEY; ?>"></script>
</head>
<body>

<div class="container" style="margin-top: 100px; max-width: 500px;">
    <h1 style="font-weight: 800">Flag Panel</h1>
    <form class="form-group" method="GET" action="login.php">
        <input type="password" class="form-control" name="password" placeholder="Password">
        <br>
        <input type="hidden" id="g-recaptcha-response" name="g-recaptcha-response">
        <input type="hidden" name="action" value="validate_captcha">
        <button type="submit" style="margin-top: 10px; width: 100%" class="btn btn-primary">Login</button>
    </form>
</div>
</body>

<script>
    grecaptcha.ready(function() {
        grecaptcha.execute('<?php echo $PUB_KEY; ?>', {action:'validate_captcha'})
            .then(function(token) {
                document.getElementById('g-recaptcha-response').value = token;
            });
    });
</script>
</html>

