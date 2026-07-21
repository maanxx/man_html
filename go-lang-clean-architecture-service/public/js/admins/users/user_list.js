import {
  checkBtnDatatable,
  statusTemplateWithDropdown,
  modalNotiUpdateStatus,
} from "/static/js/components/templates.js";
import Alert from "/static/js/common/alert.js";

import { handleAjaxError, getCookie } from "/static/js/common/helpers.js";

import { ischeckboxcheck } from "/static/js/common/helpers.js";

var user_dt;
var statusArray = JSON.parse($("#statusList").val());
var statusTemplates = {};
var statusColors = {
  active: "bg-success",
  inactive: "bg-primary",
  deleted: "bg-danger",
};

var tempUpdateData;

const userList = function () {
  const initial = function () {
    // window.user_dt = $("#user_table").DataTable({

    user_dt = $("#user_table").DataTable({
      processing: true,
      ajax: {
        url: "/api/admins/users/datatable",
        type: "POST",
      },
      dataSrc: function (response) {
        console.log(response);

        if (!response.data) {
          return [];
        }
        return response.data;
      },
      columns: [
        {
          render: function (data, type, row) {
            return checkBtnDatatable(row.id);
          },
        },
        {
          data: "custom",
          orderable: false,
          searchable: false,
        },
        {
          data: null,
          render: function (data, type, row, meta) {
            return meta.row + meta.settings._iDisplayStart + 1;
          },
        },

        {
          data: "avatar",
          orderable: false,
          searchable: false,
          render: function (data) {
            let imgSrc = data
              ? data
              : "/static/images/users/user-dummy-img.jpg";
            return `<img src="${imgSrc}" alt="avatar" class="avatar-xs rounded-circle" />`;
          },
        },
        {
          data: null,
          render: function (data, type, row) {
            let lastName = row.last_name || "";
            let firstName = row.first_name || "";
            return (lastName + " " + firstName).trim();
          },
        },
        {
          data: "status",
          render: function (data) {
            return statusTemplates[data] || "";
          },
        },
      ],
    });

    var checkAll = document.getElementById("checkAll");

    if (checkAll) {
      checkAll.onclick = function () {
        var checkboxes = document.querySelectorAll(
          '.form-check-all input[type="checkbox"]',
        );
        var checkedCount = document.querySelectorAll(
          '.form-check-all input[type="checkbox"]:checked',
        ).length;

        for (var i = 0; i < checkboxes.length; i++) {
          checkboxes[i].checked = this.checked;

          if (checkboxes[i].checked) {
            checkboxes[i].closest("tr").classList.add("table-active");
          } else {
            checkboxes[i].closest("tr").classList.remove("table-active");
          }
        }
        document.getElementById("remove-actions").style.display =
          checkedCount > 0 ? "none" : "block";
      };
    }

    user_dt.on("click", function (evt) {
      ischeckboxcheck();
    });

    user_dt.on("click", ".change_status", function (e) {
      e.preventDefault();

      let newStatusStr = $(this).data("status");

      let rowData = user_dt.row($(this).closest("tr")).index();

      processModalNotiUpdateStatus(newStatusStr, rowData);

      // let userId = rowData.id;
      // tempUpdateData = rowData;
      // $("#update_status_modal").remove();
      // let modalHtml = modalNotiUpdateStatus(newStatusStr, userId);
      // $("body").append(modalHtml);
      // $("#update_status_modal").modal("show");
    });

    $(document).on("click", "#update_status", function (e) {
      var $status = $(this).data("status");
      console.log("status: ", $status);

      var $rowIndex = $("#row_index").val();
      console.log("rowIndex: ", $rowIndex);

      var $rowData = user_dt.row($rowIndex).data();
      console.log("rowIndex: ", $rowData);

      if ($rowData.id == "") {
        $("#update_status_modal").modal("hide");
        Alert.error("Please select the row you want to change the status for");
        return;
      }
      var $statusInt = 1;
      if ($status === "inactive") $statusInt = 2;
      if ($status === "deleted") $statusInt = 3;

      $.ajax({
        url: "/api/admins/users/update-status",
        type: "PATCH",
        dataType: "json",
        data: JSON.stringify({ id: $rowData.id, status: $statusInt }),
        contentType: "application/json",
        headers: {
          "X-CSRF-Token": getCookie("csrf_"),
        },
        success: function (response) {
          Alert.success("Đổi trạng thái thành công");
          $("#user_table").DataTable().ajax.reload(null, false);
        },
        error: function (xhr) {
          handleAjaxError(xhr);
        },
      });
      $("#update_status_modal").modal("hide");
    });
  };
  return {
    init: function () {
      initStatusTemplates();
      initial();
    },
  };
};

document.addEventListener("DOMContentLoaded", function () {
  userList().init();
});

function initStatusTemplates() {
  statusArray.forEach((value, index) => {
    const otherStatuses = statusArray.filter((item) => item !== value);
    statusTemplates[value] = statusTemplateWithDropdown(
      otherStatuses,
      value,
      statusColors[value] || "badge-soft-secondary",
    );
  });
}

function processModalNotiUpdateStatus(statusName, rowIndex) {
  $("#update_status_modal").remove();
  var str = modalNotiUpdateStatus(statusName, rowIndex);
  $("body").append(str);
  $("#update_status_modal").modal("show");
}
