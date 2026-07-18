import Alert from "/static/js/common/alert.js";
import { handleAjaxError } from "/static/js/common/helpers.js";

var dataObj = {};

const userUpdate = function () {
  const handleUpdateUser = function () {
    const form = document.getElementById("form_update_user");
    if (!form) return;

    form.addEventListener("submit", function (e) {
      e.preventDefault();

      const dataObj = {
        id: parseInt(form.elements["id"].value, 10),
        last_name: form.elements["last_name"].value,
        first_name: form.elements["first_name"].value,
        email: form.elements["email"].value,
        status: parseInt(form.elements["status"].value, 10),
      };

      $.ajax({
        type: "POST",
        url: "/api/admins/users/edit",
        data: JSON.stringify(dataObj),
        contentType: "application/json",
        success: function () {
          Alert.success("Update successfully");
          setTimeout(function () {
            window.location.href = "/admins/users/list";
          }, 1500);
        },
        error: function (xhr) {
          handleAjaxError(xhr);
        },
      });
    });
  };

  return {
    init: function () {
      handleUpdateUser();
    },
  };
};

document.addEventListener("DOMContentLoaded", function () {
  userUpdate().init();
});
