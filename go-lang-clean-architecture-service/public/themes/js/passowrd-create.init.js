import {validatePasswordAuth} from '/static/themes/js/common/auth_inits.js';

// password addon
Array.from(document.querySelectorAll("form .auth-pass-inputgroup")).forEach(function (item) {
    Array.from(item.querySelectorAll(".password-addon")).forEach(function (subitem) {
            subitem.addEventListener("click", function (event) {
                var passwordInput = item.querySelector(".password-input");
                if (passwordInput.type === "password") {
                    passwordInput.type = "text";
                } else {
                    passwordInput.type = "password";
                }
            });
        });
    });

// passowrd match
var password = document.getElementById("password-input"),
    confirm_password = document.getElementById("confirm-password-input");

function validatePassword() {
    if (password.value != confirm_password.value) {
        confirm_password.setCustomValidity("Passwords Don't Match");
    } else {
        confirm_password.setCustomValidity("");
    }
}

//Password validation
password.onchange = validatePassword;

validatePasswordAuth();