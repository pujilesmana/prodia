$(document).ready(function () {
   $('#login').submit(function () {
      const email = $('[name="email"]').val();
      const password = $('[name="password"]').val();

      var settings = {
         "url": "http://localhost:8080/api/v1/users/login",
         "method": "POST",
         "timeout": 0,
         "headers": {
            "Content-Type": "application/json"
         },
         "data": JSON.stringify({
            "email": email,
            "password": password
         }),
      };

      $.ajax(settings).done(function (response) {
         localStorage.setItem('token', response.data.token);
         window.location.replace("http://localhost/frontend/dashboard.php");
      });
   })
})


