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
          data: "id",
          orderable: false,
          searchable: false,
          render: function (data) {
            return `<div class="form-check">
                      <input class="form-check-input user-checkbox" type="checkbox" value="${data}">
                    </div>`;
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
    
    const checkAll = document.getElementById("checkAll");
    const removeBtn = document.getElementById("remove-actions");
    if (removeBtn) removeBtn.style.display = "none";

    if (checkAll) {
      checkAll.addEventListener("change", function () {
        const checkboxes = document.querySelectorAll(".user-checkbox");
        checkboxes.forEach((cb) => {
          cb.checked = checkAll.checked;
          if (cb.checked) {
            cb.closest("tr").classList.add("table-active");
          } else {
            cb.closest("tr").classList.remove("table-active");
          }
        });
        toggleRemoveBtn();
      });
    }

    $(document).on("change", ".user-checkbox", function () {
      if (this.checked) {
        this.closest("tr").classList.add("table-active");
      } else {
        this.closest("tr").classList.remove("table-active");
        if (checkAll) checkAll.checked = false;
      }
      toggleRemoveBtn();
    });

    function toggleRemoveBtn() {
      const checkedCount = document.querySelectorAll(".user-checkbox:checked").length;
      if (removeBtn) {
        removeBtn.style.display = checkedCount > 0 ? "block" : "none";
      }
    }
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
