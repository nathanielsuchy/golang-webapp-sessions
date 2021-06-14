<div class="jumbotron">
    <form action="/sign-in" method="post">
        <div class="form-group">
          <label for="exampleInputEmail1">Email Address</label>
          <input type="email" class="form-control" name="email" aria-describedby="emailHelp" placeholder="Enter email" required>
          <small id="emailHelp" class="form-text text-muted">We'll never share your email with anyone else.</small>
        </div>
        <div class="form-group">
          <label for="exampleInputPassword1">Password</label>
          <input type="password" class="form-control" name="password" placeholder="Password" required>
        </div>
        <button type="submit" class="btn btn-primary">Sign in</button>
      </form>
</div>