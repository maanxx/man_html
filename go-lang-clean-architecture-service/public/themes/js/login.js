import { accessToken,saveAuth,redirectURL } from '/static/themes/js/common/inits.js';
import { alertErrorTopRight,alertSuccessTopRight } from '/static/themes/js/common/alerts.js';

//Password Without Spaces
$(".password-input").on("input", function () {
    var sanitizedPassword = $(this).val().replace(/\s/g, "");
    $(this).val(sanitizedPassword);
})

const loginAuth = (function() {
    const urlLogin = $("#frm_login").attr('action');

    $(document).on("submit","#frm_login",function(e){
        e.preventDefault();
        var email = $("#username").val();
        var pass = $("#password-input").val();
        if (email == "" || pass == "" || urlLogin == ""){
            alertErrorTopRight("Email and password are incorrect!");
        }
        var formData = new FormData($(this)[0]);
        login(urlLogin,{"email":email,"password":pass});
    });    
})();

const signupAuth =  (function() {
    const urlSignup = $("#frm_signup").attr('action');
    $(document).on("submit","#frm_signup",function(e){
        e.preventDefault();
        var email = $("#email-input").val();
        var pass = $("#password-input").val();
        var lastname = $("#lastname-input").val();
        var firstname = $("#firstname-input").val();
        if (email == "" || pass == "" || urlSignup == "" || lastname =="" || firstname == ""){
            errorAlert("Form data has not been entered yet!")
            return;
        }
        var data = {"email":email,"password":pass,"lastname":lastname, "firstname":firstname};
        signup(urlSignup,data);
    });
})();


function login(url,data){
    axios.post(url,data,{
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(function (response) {   
        saveAuth(response.data.data); 
        alertSuccessTopRight("Logged in successfully!")    
        redirectURL();
    }).catch(error => {
        if (error.response && error.response.data && error.response.data.details && error.response.data.details.msg) {
          const errorMessage = error.response.data.details.msg;        
          alertErrorTopRight(errorMessage);
        } else {
          console.error('An error occurred:', error);
        }
    });    
}

function signup(url,data){
    axios.post(url,data,{
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(function (response) {
        successAlert("Registered successfully!");        
        setTimeout(function () {
            window.location.href = "/login";
        }, 1200);
    }).catch(function (error) {
        if (error.response && error.response.data && error.response.data.details && error.response.data.details.msg) {
            const errorMessage = error.response.data.details.msg;        
            alertErrorTopRight(errorMessage);
        } else {
            console.error('An error occurred:', error);
        }
    });
}

function errorAlert(msg) {
    Swal.fire({
      position: "top-end", 
      title: "Error",
      text: msg,
      icon: "error",
      showConfirmButton: false,
      timer: 2500,
      showCloseButton: true
    });
}

function successAlert(msg) {
    Swal.fire({
      position: "top-end",
      icon: "success",
      title: msg,
      showConfirmButton: false,
      timer: 1500,
      showCloseButton: true
    });
}

// Save the data to local store using Axios
function saveToLocalStore(data) {
    // Assuming you're using the 'localStorage' object to store data
    localStorage.setItem('access_token', data.access_token.token);
    localStorage.setItem('expire_in', data.access_token.expire_in);
    localStorage.setItem('first_name', data.user.first_name);
    localStorage.setItem('last_name', data.user.last_name);
}