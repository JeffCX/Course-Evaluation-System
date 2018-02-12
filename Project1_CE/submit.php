<!DOCTYPE html>
<html>
<head>
	<title></title>
</head>
<body>
<!doctype html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
    <title>NYU Course Evaluation Login</title>
    <style type="text/css">
      body{
         background: linear-gradient(white,#f2edf1,white);
      }
      img{
        width:50%;
        height:100%;
        border:1vh solid purple;
      }

      .left_decerate{
        border-left:2vh solid purple;
        border-bottom:2vh solid purple;
        border-bottom-left-radius: 20vh;
        text-align:center;

      }


      .decorate{
        margin-top:3vh;
        border-left:10px double purple;
        border-right:10px double purple;
        border-radius:20px;
        box-shadow:2px 1px 2px black;
        padding:15%;
      }

      button{
        margin:auto;
        display: block;

      }

       .seperate_more{
        margin:2vh;
      }

      .form-check-input{
        margin-left:3vh;
      }

      label{
        margin-left:6vh;
      }

      .responsive{
        width:85%;
      }

      .responsive_button{
        margin-left:2.5%;
        width:10%;
      }


      .responsive_button_sub{
        width:100%;
      }

      label{
        text-align: left;
      }

      .head_doc{
        background: purple;
        width:100%;
        height:10vh;
        line-height: 10vh;
        color:white;
      
      }
    
      h1{
        display: inline-block;
        vertical-align: middle;
        
      }

      button{
        display: inline-block;
      }

    
      a:hover{
      	background:rgb(128, 128, 128);
      	color:purple;
      }

      iframe{
     	width:100%;
      	height:80vh;
      	border:1vh solid purple;
      	border-radius: 1vh;
      }

      
    </style>
  </head>
  <body>
  <header class="container-fluid head_doc">
    <section class="row">
      <article class="col-sm-4 ">
          
          <h1>NYU Course Evaluation</h1>
      </article >
      <article class="col-sm-2 offset-sm-6">
          Welcome, <?php echo $_POST["NetID"]?> |
          <button class="btn btn-light"><a href="index.html">Log Out</a></button>
      </article>
    </section>
  </header>

  <section>
	 <ul class="nav nav-pills nav-fill">
	  <li class="nav-item">
	    <a class="nav-link" href="https://www.nyu.edu/life/information-technology/instructional-technology-support/instructional-technology-tools-and-services/nyu-classes.html">NYU Classes</a>
	  </li>
	  <li class="nav-item">
	    <a class="nav-link" href="http://albert.nyu.edu/albert_index.html">NYU Ablert</a>
	  </li>
	  <li class="nav-item">
	    <a class="nav-link" href="https://www.nyu.edu/life/information-technology/communication-and-conferencing/nyu-email.html">NYU Gmail</a>
	  </li>
	  <li class="nav-item">
	    <a class="nav-link" href="https://www.google.com">Google</a>
	  </li>
	</ul>
</section>

 <section>
	  <section class="row">
      <article class="col-sm-10 offset-sm-1">
          
         <iframe src="Calendar_Template/fullcalendar.html"></iframe>
      </article>
    </section>
</section>



 
    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
  </body>
</html>
</body>
</html>