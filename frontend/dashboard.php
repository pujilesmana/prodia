<?php

$curl = curl_init();

curl_setopt_array($curl, array(
   CURLOPT_URL => "https://api.rajaongkir.com/starter/province",
   CURLOPT_RETURNTRANSFER => true,
   CURLOPT_ENCODING => "",
   CURLOPT_MAXREDIRS => 10,
   CURLOPT_TIMEOUT => 30,
   CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
   CURLOPT_CUSTOMREQUEST => "GET",
   CURLOPT_HTTPHEADER => array(
      "key: 91f364d3a44051e30c2bae62c1f3090b"
   ),
));

$response = curl_exec($curl);
$err = curl_error($curl);

curl_close($curl);

$kota_array   = json_decode($response, true);

if ($kota_array['rajaongkir']['status']['code'] == 200) :
   $kota_result  = $kota_array['rajaongkir']['results'];
else :
   die('This key has reached the daily limit.');
endif;
?>

<!DOCTYPE html>
<html lang="en">

<head>
   <meta charset="UTF-8">
   <meta http-equiv="X-UA-Compatible" content="IE=edge">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <title>Dashboard Data</title>

   <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

   <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.6.1/css/all.css" integrity="sha384-gfdkjb5BdAXd+lj+gudLWI+BXq4IuLW5IT+brZEZsLFm++aCMlF1V92rMkPaX4PP" crossorigin="anonymous">
   <link rel="stylesheet" href="https://cdn.datatables.net/1.11.3/css/jquery.dataTables.min.css">

</head>

<body>
   <table id="example" class="display" style="width:100%">
      <thead>
         <tr>
            <th width="50">No</th>
            <th>Nama Kota</th>
         </tr>
      </thead>
      <tbody>
         <?php
         $no = 0;
         foreach ($kota_result as $value) :
            $nama = $value["province"];
            $no++;
         ?>
            <tr>
               <td><?= $no ?></td>
               <td><?= $nama ?></td>
            </tr>
         <?php endforeach; ?>
      </tbody>
      <tfoot>
         <tr>
            <th>No</th>
            <th>Nama Kota</th>
         </tr>
      </tfoot>
   </table>
</body>

<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
<script src="https://cdn.datatables.net/1.11.3/js/jquery.dataTables.min.js"></script>
<script>
   $(document).ready(function() {
      $('#example').DataTable();
   });
</script>

</html>