import { checkBtnDatatable } from "/static/js/components/templates.js";

import { ischeckboxcheck } from "/static/js/common/helpers.js";

var user_dt;

const userList = function () {
  const initial = function () {
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
            let imgSrc = data ? data : "static/images/users/user-dummy-img.jpg";
            return `<img src="${imgSrc}" alt="avatar" class="avatar-xs rounded-circle" />`;
          },
        },
        {
          data: "name",
        },
        {
          data: "status",
          render: function (data) {
            let badgeClass = data === "Active" ? "bg-success" : "bg-danger";
            let label = data === "Active" ? "Active" : "Inactive";
            return `<span class="badge ${badgeClass}">${label}</span>`;
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

    $("#user_table").on("click", function (evt) {
      ischeckboxcheck();
    });

    
  };
  return {
    init: function () {
      initial();
    },
  };
};

document.addEventListener("DOMContentLoaded", function () {
  userList().init();
});
