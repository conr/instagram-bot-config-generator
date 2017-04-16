<html>
  <head>
    <title>Create InstaBot Config File</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
    <style>
      .input-group {
        padding: 25px;
      }
      .btn-lg {
        margin-bottom: 20px;
      }
    </style>
  </head>
  <body>
    <form action="/create" method="post">
      <div class="col-md-12">
        <h1 class="text-center">Instabot Configuration Generator :)</h1>
        <div class="container">
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="text" class="form-control" placeholder="Username" name="username">
            <span class="input-group-addon" id="basic-addon1">#</span>
            <input type="password" class="form-control" placeholder="Password" name="password">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="number" class="form-control" placeholder="Likes per day" name="likesPerDay">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="number" class="form-control" placeholder="Media Max likes" name="mediaMaxLikes">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="number" class="form-control" placeholder="Media Min Likes" name="mediaMinLikes">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="number" class="form-control" placeholder="Follows per day" name="followsPerDay">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="number" class="form-control" placeholder="Follow time (seconds)" name="followTime">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="number" class="form-control" placeholder="Unfollows per day" name="unfollowsPerDay">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="number" class="form-control" placeholder="Comments per day" name="commentsPerDay">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="text" class="form-control" placeholder="Target hashtags (separate by comma)" name="hashtags">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="text" class="form-control" placeholder="Blacklist hashtags (separate by comma)" name="hashtagsBlacklist">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="text" class="form-control" placeholder="Blacklist users (separate by comma)" name="usersBlacklist">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="number" class="form-control" placeholder="Max likes per hashtag" name="maxLikesPerHashtag">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="number" class="form-control" placeholder="Unfollow break minimum" name="unfollowBreakMin">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="number" class="form-control" placeholder="Unfollow break maximum" name="unfollowBreakMax">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="number" class="form-control" placeholder="Log mod" name="logMod">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="text" class="form-control" placeholder="Proxy" name="proxy">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="text" class="form-control" placeholder="Unwanted Username List" name="unwantedUsernameList">
          </div>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">@</span>
            <input type="text" class="form-control" placeholder="Unfollow Whitelist" name="unfollowWhitelist">
          </div>

          <input type="submit" value="Create Config" class="col-md-12 btn-lg">
        </div>
      </div>

    </form>
  </body>
</html>
