var user_dt;

const userList = function () {
  const initial = function () {
    user_dt = new DataTable("#user_table", {
      serverSide: true,
      processing: true,
      ajax: {
        url: "/api/admins/users",
        type: "GET",
      },
      columns: [
        // Cột 1: Checkbox
        {
          data: "id",
          orderable: false,
          searchable: false,
          render: function (data) {
            return `<div class="form-check">
                      <input class="form-check-input" type="checkbox" value="${data}">
                    </div>`;
          },
        },

        // Cột 2: Action Buttons
        {
          data: "id", // Sử dụng ID để làm data-id cho các nút Edit/Delete
          orderable: false,
          searchable: false,
          render: function (data) {
            return `<button class="btn btn-sm btn-soft-info edit-btn" data-id="${data}"><i class="ri-pencil-fill"></i></button>
                    <button class="btn btn-sm btn-soft-danger delete-btn" data-id="${data}"><i class="ri-delete-bin-fill"></i></button>`;
          },
        },

        // Cột 3: NO. (Số thứ tự tự động)
        {
          data: null, // Không lấy field nào từ API cả
          orderable: false,
          searchable: false,
          render: function (data, type, row, meta) {
            // meta.row là vị trí dòng hiện tại trên trang (0, 1, 2...)
            // meta.settings._iDisplayStart là vị trí bắt đầu của trang hiện tại (ví dụ đang ở trang 2 thì nó là 10)
            return meta.row + meta.settings._iDisplayStart + 1;
          },
        },

        // Cột 4: Avatar (Image)
        {
          data: "image", // Lấy đúng theo tag `json:"image"`
          orderable: false,
          searchable: false,
          render: function (data) {
            // Đề phòng user chưa có ảnh, dùng fallback avatar
            let imgSrc = data ? data : "/static/images/default-avatar.png";
            return `<img src="${imgSrc}" alt="avatar" class="avatar-xs rounded-circle" />`;
          },
        },

        // Cột 5: Full Name
        {
          data: "full_name", // Lấy đúng theo tag `json:"full_name"`
        },

        // Cột 6: Role
        {
          data: "role_id", // Tạm thời dùng role_id vì struct chưa có tên Role
          render: function (data) {
            // Sau này backend API trả về role_name thì em map trường data ở trên và bỏ dòng render này đi
            return data ? data : `<span class="text-muted">N/A</span>`;
          },
        },

        // Cột 7: Status
        {
          data: "status", // Kiểu số Int (ví dụ: 1 - Active, 0 - Inactive)
          render: function (data) {
            // Anh giả sử 1 là Active, em có thể sửa số này theo đúng quy ước backend của em
            let badgeClass = data === 1 ? "bg-success" : "bg-danger";
            let label = data === 1 ? "Active" : "Inactive";

            // Dùng dấu backtick ` ` (nằm cùng phím dấu ngã ~ góc trái bàn phím) để nội suy biến JS
            return `<span class="badge ${badgeClass}">${label}</span>`;
          },
        },

        // Cột 8: Audit Info (Thường là Created At lấy từ core.SQLModel)
        {
          data: "created_at", // Giả định SQLModel map ra json là created_at
          render: function (data) {
            if (!data) return "";
            // DataTables có thể tự render chuỗi ISO Time, nhưng nếu em muốn format đẹp thì có thể dùng moment.js hoặc hàm Date của JS
            let date = new Date(data);
            return (
              date.toLocaleDateString("vi-VN") +
              " " +
              date.toLocaleTimeString("vi-VN")
            );
          },
        },
      ],
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
