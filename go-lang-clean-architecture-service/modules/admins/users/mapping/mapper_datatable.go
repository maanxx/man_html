package mapping

import (
	"app/modules/admins/users/entity"
	"app/modules/admins/users/transport/respones"
	"fmt"
)

func MapperDatatable(users *[]entity.User) *[]respones.DataTableUserResp {
	var result []respones.DataTableUserResp

	if users == nil {
		return nil
	}

	for _, v := range *users {
		user := respones.DataTableUserResp{
			ID:     v.ID,
			Name:   v.FullName,
			Avatar: v.Image,
			Custom: fmt.Sprintf(`<a id="delete-item-btn" href="/admins/users/edit/%d" class="btn btn-sm btn-soft-info edit-btn"><i class="ri-pencil-fill"></i></a>
                    <button class="btn btn-sm btn-soft-danger delete-btn" data-id="%d"><i class="ri-delete-bin-fill"></i></button>`, v.ID, v.ID),
		}

		if v.Status == 1 {
			user.Status = "Active"
		} else if v.Status == 2 {
			user.Status = "Inactive"
		}

		result = append(result, user)
	}

	return &result
}
