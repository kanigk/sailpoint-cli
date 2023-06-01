package config

const OAuthSuccessPage = `
<!doctype html>
<meta charset="utf-8">
<title>Success: SailPoint CLI</title>
<style type="text/css">
body {
  color: #1B1F23;
  background: #F6F8FA;
  font-size: 14px;
  font-family: -apple-system, "Segoe UI", Helvetica, Arial, sans-serif;
  line-height: 1.5;
  max-width: 620px;
  margin: 28px auto;
  text-align: center;
}
h1 {
  font-size: 24px;
  margin-bottom: 0;
}
p {
  margin-top: 0;
}
.box {
  border: 1px solid #E1E4E8;
  background: white;
  padding: 24px;
  margin: 28px;
}

html, body {
  height: 100%;
  overflow: hidden;
}

body {
  background-image:
			radial-gradient(at 0% 0%, rgba(0 113 206 / 0.33) 0px, transparent 50%),
			radial-gradient(at 98% 1%, rgba(204 39 176 / 0.33) 0px, transparent 50%);
}

</style>
<body>
  <div class="box">
    <h1>Successfully authenticated with the SailPoint CLI</h1>
    <p>You may now close this tab and return to the terminal.</p>
  </div>
</body>
`
