import Alert from "/static/js/common/alert.js";
import {
  handleAjaxError,
  loader,
  getCookie,
} from "/static/js/common/helpers.js";

var $listID = {};
var $idsArr = [];

const deleteUser = function () {
  const handleDeleteUser = function () {
    var isDeleteUser = 0;
    $(document).on("click", ".delete-btn", function (evt) {
      isDeleteUser = 0;
      // var tr = $(this).closest("tr");
      // var $rowData = user_dt.row(tr).data();
      // var $dataID = $rowData._id;

      var $dataID = parseInt($(this).attr("data-id"));

      $idsArr = [$dataID];
      $listID = {}; // Clear existing items
      $listID[$dataID] = 1;

      $("#delete_modal").modal("show");
    });

    $("#delete_modal").on("shown.bs.modal", function (e) {
      isDeleteUser = 0;
    });

    // Delete Roles
    $(document).on("click", "#btn-submit-delete", function (evt) {
      evt.preventDefault();
      if (isDeleteUser == 1) {
        return;
      }

      isDeleteUser = 1;

      var checked = document.querySelectorAll(
        '.form-check-all input[type="checkbox"]:checked',
      );
      if (checked.length > 0) {
        $idsArr.length = 0;
        checked.forEach(function (check) {
          var checkVal = parseInt(check.value);
          $idsArr.push(checkVal);
          $listID[checkVal] = 1;
        });
      }
      loader();
      if ($idsArr.length > 0) {
        $.ajax({
          url: $("#remove-actions").data("url"),
          method: "DELETE",
          dataType: "json",
          contentType: "application/json",
          data: JSON.stringify({ ids: $idsArr }),
          enctype: $(this).attr("enctype") || "multipart/form-data",
          headers: {
            "X-CSRF-Token": getCookie("csrf_"),
          },
          success: function (res) {
            Alert.success(res.message);
            $("#delete_modal").modal("hide");
            $("#user_table").DataTable().ajax.reload(null, false);
            $("#remove-actions").hide();
            checkAll.checked = false;

            $idsArr.forEach(function ($row) {
              var $detailID = $("#info_status").data("id");
              if ($detailID == $row) {
                var data = $("#user_table").DataTable().row(0).data();
                detailUser(data);
              }
            });
          },
          error: function (xhr) {
            handleAjaxError(xhr);
            $("#delete_modal").modal("hide");
          },
        });
      }
      loader(false);
      isDeleteUser = 1;
    });
  };
  return {
    init: function () {
      handleDeleteUser();
    },
  };
};

document.addEventListener("DOMContentLoaded", function () {
  deleteUser().init();
});
