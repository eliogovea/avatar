

$(document).ready(function () {

function renderUserPage(username) {
  $('#content').empty();

  let images = `
    <div id="images" class="row"></div>
  `

  $('#content').append(images);

  let approvedURL = `"/avatars/approved/` + username + `"`
  let approved = `
    <div class="card-profile col" style="width: 20rem;">
      <img class="card-avatar" src=` + approvedURL + ` alt="Card image cap">
      <div class="card-body">
        <p class="card-text">Aprobado</p>
      </div>
    </div>
  `

  $('#images').append(approved);

  let pendingURL = `"/avatars/pending/` + username + `"`
  let pending = `
   <div class="card-profile col" style="width: 20rem;">
      <img class="card-avatar" src=` + pendingURL + ` alt="Card image cap">
      <div class="card-body">
        <p class="card-text">Pendiente</p>
      </div>
    </div>
  `

  $('#images').append(pending);
  
  let resizerHtml = `
   <div class="container-fluid fileinput fileinput-new text-center" data-provides="fileinput">
      <div id="resizer"></div>
      <div>
        <span class="upload-btn-wrapper btn btn-raised btn-round btn-default btn-file">
          <span>Cambiar mi avatar</span>
            <input type="file" id="load" name="file" />
          </span>
          <button class="btn btn-round" id="send">Enviar</button>
      </div>
    </div>
  `;

//  $('#content').append(resizerHtml);
  $('#images').append(resizerHtml);

  $('#login-form').fadeOut();

  var resizer = $('#resizer').croppie({
    viewport: { width: 200, height: 200, type: "circle" },
    boundary: { width: 200, height: 200 },
    enableOrientation: true,
    mouseWheelZoom: 'ctrl',
    url: "web/static/img/default-avatar.jpeg" 
  });

  $('#send').on('click', function (event) {
    resizer.croppie('result', 'base64').then(result => {
      fetch('/upload', {
        method: 'POST',
        headers: {
          'authorization': "JWT " + (localStorage.getItem('accessToken') || ""),
          'content-type': 'application/json; charset=utf-8'
        },
        body: JSON.stringify({
          image: result
        })
      }).then(response => {
        if (!response.ok) {
          $('#alerts').prepend(createNotification('danger', 'Upload failed'));
          setTimeout(function () {
            $('#alerts').children('.alert:first-child').fadeOut();
          }, 1000);
        } else {
          $('#alerts').prepend(createNotification('success', 'Upload OK'));
          setTimeout(function () {
            $('#alerts').children('.alert:first-child').fadeOut();
          }, 1000);
        }
      }).catch(error => {
        alert("ERROR: ", error);
      });
    });
  });

  $('#load').on('change', function (event) {
    resizer.croppie('bind', {
      url: URL.createObjectURL(event.target.files[0])
    });
  });
}

function renderAboutPage() {
  $('#content').empty();
  // TODO
}

function renderAdminPage() {
  $('#content').empty();
  // TODO
}

function createNotification(type, msg) {
  return `
      <div class="float alert alert-` + type + `">
        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
          <i class="material-icons">close</i>
        </button>
        <span>
          <b> Success - </b>` + msg + `
        </span>
      </div>
    `
}

$('#login-button').on('click', function (event) {
  let username = document.getElementById('username').value;
  fetch('/login', {
    method: 'POST',
    headers: {
      'content-type': 'application/json; charset=utf-8'
    },
    body: JSON.stringify({
      username: username,
      password: document.getElementById('password').value
    })
  }).then(response => {
    return response.json();
  }).then(json => {
    localStorage.setItem('username', username);
    localStorage.setItem('accessToken', json.token);
    let html = createNotification('success', 'Login OK');
    $('#alerts').empty();
    $('#alerts').prepend(html).fadeIn();
    setTimeout(function () {
      $('#alerts').children('.alert:first-child').fadeOut();
    }, 1000);
    renderUserPage(username);
  }).catch(error => {
    html = createNotification('danger', 'Login failed');
    $('#alerts').prepend(html).fadeIn();
    setTimeout(function () {
      $('#alerts').children('.alert:first-child').fadeOut();
    }, 1000);
  });
});

$('#logout-button').on('click', function (event) {
  localStorage.clear();
  location.reload();
  // TODO
});

$('#profile-button').on('click', function (event) {
  renderUserPage(localStorage.getItem('username'));
});

$('#about-button').on('click', function (event) {
  renderAboutPage();
});

$('#admin-button').on('click', function (event) {
  renderAdminPage();
});

});
