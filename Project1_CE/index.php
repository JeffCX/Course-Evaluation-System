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
        margin-top:10vh;
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
    
    </style>
  </head>
  <body>
  <header class="container-fluid">
    <section class="row">
      <article class="col-sm-8">
        <img src="img/Tandon_logo_eng.jpg" alt="Tandon_logo">
      </article>

       <article class="col-sm-4 left_decerate">
          Welcome, <?php echo $_POST["NetID"]?>
          <button><a href="index.html">Log Out</a></button>
      </article>
    </section>
  </header>
  <form method="post" action="submit.php">
  <section class="container">
    <article class="row">
      <section class="col-sm-12 decorate">
          <h1>Say Feely</h1>
          
            <section class="container">
              <section class="row seperate">
                <article class="col-sm-3">
                  Question1:
                </article>
                  <input type="text" name="NetID" class="col-sm-9" />
              </section>
              <section class="row seperate">
                <article class="col-sm-3">
                  Question2
                </article>
                <input type="Password" name="password" class="col-sm-9"/>             
              </section>
              <section class="row seperate">
               <article class="col-sm-3">
                  Question3
                </article>
                <input type="Password" name="password" class="col-sm-9"/>             
                
               
              </section>
          </section>

      </section>
    </article>
  </section>

    <section class="container">
    <article class="row">
      <section class="col-sm-12 decorate">
          <h1>Say Feely</h1>
        
            <section class="container">
              <section class="row seperate">
                <article class="col-sm-3">
                  Question1:
                </article>
                  <input type="text" name="NetID" class="col-sm-9" />
              </section>
              <section class="row seperate">
                <article class="col-sm-3">
                  Question2
                </article>
                <input type="Password" name="password" class="col-sm-9"/>             
              </section>
              <section class="row seperate">
               <article class="col-sm-3">
                  Question3
                </article>
                <input type="Password" name="password" class="col-sm-9"/>             
                
               
              </section>
          </section>

      
      </section>
    </article>
  </section>

    <section class="container">
    <article class="row">
      <section class="col-sm-12 decorate">
          <h1>Say Feely</h1>
        
            <section class="container">
              <section class="row seperate">
                <article class="col-sm-3">
                  Question1:
                </article>
                  <input type="text" name="NetID" class="col-sm-9" />
              </section>
              <section class="row seperate">
                <article class="col-sm-3">
                  Question2
                </article>
                <input type="Password" name="password" class="col-sm-9"/>             
              </section>
              <section class="row seperate">
               <article class="col-sm-3">
                  Question3
                </article>
                <input type="Password" name="password" class="col-sm-9"/>             
                
               
              </section>

              <section class="row seperate_more">
               
                  <input type="Submit" class="col-sm-12" value = "Submit"/>
               
              </section>
          </section>

         
      </section>
    </article>
  </section>
            <form>



    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
  </body>
</html>