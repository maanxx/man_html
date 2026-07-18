import Alert from "/static/js/common/alert.js";
import { handleAjaxError } from "/static/js/common/helpers.js";

const deleteUser = function () {
  const handleDeleteUser = function () {
    $(document).on("click", ".delete-btn", function (e) {
      e.preventDefault();
      const id = $(this).data("id");
      $("#delete_user_id").val(id); 
      $("#delete_modal").modal("show");
    });

    $(document).on("click", "#remove-actions", function (e) {
      e.preventDefault();
      $("#delete_user_id").val(""); 
    });

    const $form = $("#form_delete_user");
    if (!$form.length) return;

    $form.on("submit", function (e) {
      e.preventDefault();
      
      let idsToDelete = [];
      const singleId = $("#delete_user_id").val();

      if (singleId !== "") {
        idsToDelete.push(parseInt(singleId, 10));
      } else {
        const checkedBoxes = document.querySelectorAll(".user-checkbox:checked");
        checkedBoxes.forEach(function (cb) {
          idsToDelete.push(parseInt(cb.value, 10));
        });
      }

      if (idsToDelete.length === 0) {
        $("#delete_modal").modal("hide");
        return;
      }

      $.ajax({
        type: "POST",
        url: "/api/admins/users/delete",
        data: JSON.stringify({ ids: idsToDelete }), 
        contentType: "application/json",
        success: function () {
          Alert.success("Xóa người dùng thành công!");
          $("#delete_modal").modal("hide");
          
          if ($.fn.DataTable.isDataTable("#user_table")) {
            $("#user_table").DataTable().ajax.reload(null, false);
          }
          
          const checkAll = document.getElementById("checkAll");
          if (checkAll) checkAll.checked = false;
          const removeBtn = document.getElementById("remove-actions");
          if (removeBtn) removeBtn.style.display = "none";
        },
        error: function (xhr) {
          handleAjaxError(xhr);
        },
      });
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
