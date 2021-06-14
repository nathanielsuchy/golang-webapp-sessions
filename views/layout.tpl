<!DOCTYPE html>
<html lang="en">
    <head>
        <title>Web App Title</title>
        <meta charset="utf-8">
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
        <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
    </head>
  <body>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <a class="navbar-brand" href="/">Web App Title</a>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarText" aria-controls="navbarText" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="navbar-collapse" id="navbarText">
    
            <ul class="navbar-nav mr-auto">
            <li class="nav-item active">
                <a class="nav-link" href="/">Home <span class="sr-only">(current)</span></a>
            </li>
            </ul>
    
            <ul class="navbar-nav ml-auto">
                {{ if $.IsSignedIn }}
                    <li class="nav-item active">
                        <a class="nav-link" href="/sign-out">Sign out <span class="sr-only">(current)</span></a>
                    </li>
                {{ end }}
    
                {{ if $.IsNotSignedIn }}
                    <li class="nav-item active">
                        <a class="nav-link" href="/sign-in/">Sign in <span class="sr-only">(current)</span></a>
                    </li>
                    <li class="nav-item active">
                        <a class="nav-link" href="/sign-up/">Create Account <span class="sr-only">(current)</span></a>
                    </li>
                {{ end }}
            </ul>
        </div>
    </nav>

    {{ if and .flash.error }}
      <div class="alert alert-danger">
        {{.flash.error}}
      </div>
    {{ end }}

    {{ if and .flash.warning }}
      <div class="alert alert-warning">
        {{.flash.warning}}
      </div>
    {{ end }}

    {{ if and .flash.success }}
      <div class="alert alert-primary">
        {{.flash.success}}
      </div>
    {{ end }}

    {{ if and .flash.notice }}
      <div class="alert alert-secondary">
        {{.flash.notice}}
      </div>
    {{ end }}

    {{.LayoutContent}}
    
    <hr />
    <p style="text-align: center; font-style: italic;">Copyright 2021 Site Name</p>
    <script src="/static/js/reload.min.js"></script>
  </body>
</html>
