import Alert from "/static/js/common/alert.js";
import { handleAjaxError } from "/static/js/common/helpers.js";

var formData = {};
var dataObj = {};

const userCreate = function () {
  const handleCreateUser = function () {
    const $form = $("#form_create_user");

    if (!$form.length) return;

    $form.on("submit", function (e) {
      e.preventDefault();

      $.each($form.serializeArray(), function (_, kv) {
        dataObj[kv.name] = kv.value;
      });

      // dataObj.status = 3
      if (dataObj.status) {
        dataObj.status = parseInt(dataObj.status, 10);
      }
      
      $.ajax({
        type: "POST",
        url: "/api/admins/users/create",
        data: JSON.stringify(dataObj),
        contentType: "application/json",
        success: function () {
          Alert.success("Thêm mới thành công!");
          $form[0].reset();

          $("#create_user_modal").modal("hide");

          if ($.fn.DataTable.isDataTable("#user_table")) {
            $("#user_table").DataTable().ajax.reload(null, false);
          }
        },
        error: function (xhr) {
          handleAjaxError(xhr);
        },
      });
    });
  };
  return {
    init: function () {
      handleCreateUser();
    },
  };
};

document.addEventListener("DOMContentLoaded", function () {
  userCreate().init();
});
