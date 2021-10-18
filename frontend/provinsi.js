$(document).ready(function () {

   var myHeaders = new Headers();
   myHeaders.append("key", "91f364d3a44051e30c2bae62c1f3090b");

   var raw = "";

   var requestOptions = {
      method: 'GET',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
   };

   fetch("https://api.rajaongkir.com/starter/province", requestOptions)
      .then(response => response.text())
      .then(result => console.log(result))
      .catch(error => console.log('error', error));
})
