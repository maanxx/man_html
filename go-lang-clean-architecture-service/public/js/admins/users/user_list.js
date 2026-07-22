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
var listData = {};

const userList = function () {
  const initial = function () {
    user_dt = $("#user_table").DataTable({
      serverSide: true,
      processing: true,
      ajax: {
        url: "/api/admins/users/datatable",
        type: "POST",
        dataType: "json",
        data: function (d) {
          d = { ...d, ...listData };
          return d;
        },
        dataSrc: function (response) {
          if (!response.data) {
            return [];
          }
          return response.data;
        },
      },
      columns: [
        {
          render: function (data, type, row) {
            return checkBtnDatatable(row.id);
          },
        },
        {
          data: "status",
          visible: false,
          searchable: true,
        },
        {
          data: "custom",
          orderable: false,
          searchable: false,
        },
        {
          data: null,
          render: function (data, type, row, meta) {
            return meta.row + 1;
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

    $(document).on("input", "#inputSearch", function () {
      user_dt.search($("#inputSearch").val()).draw();
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
  const filterStatus = function () {
    $("#filter-data").on("change", function () {
      let statusVal = $(this).val();
      if (statusVal === "") {
        user_dt.column(1).search("").draw();
      } else {
        user_dt
          .column(1)
          .search("^" + statusVal + "$", true, false)
          .draw();
      }
    });
  };

  return {
    init: function () {
      initStatusTemplates();
      initial();
      filterStatus();
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
