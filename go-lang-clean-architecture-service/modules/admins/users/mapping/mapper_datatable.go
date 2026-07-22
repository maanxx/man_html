package mapping

import (
	"app/modules/admins/common/consts"
	"app/modules/admins/users/entity"
	"app/modules/admins/users/transport/respones"
	"fmt"
)

func MapperDatatable(users []*entity.User) *[]respones.DataTableUserResp {
	var result []respones.DataTableUserResp

	if users == nil {
		return nil
	}

	for _, v := range users {
		user := respones.DataTableUserResp{
			ID:        v.ID,
			Name:      v.FullName,
			LastName:  v.LastName,
			FirstName: v.FirstName,
			Avatar:    v.Image,
			Email:     v.Email,
			Custom: fmt.Sprintf(`<a href="/admins/users/edit/%d" class="btn btn-sm btn-soft-info edit-btn"><i class="ri-pencil-fill"></i></a>
                    <button  id="delete-item-btn" class="btn btn-sm btn-soft-danger delete-btn" data-id="%d"><i class="ri-delete-bin-fill"></i></button>`, v.ID, v.ID),
		}

		// switch v.Status {
		// case 1:
		// 	user.Status = "active"
		// case 2:
		// 	user.Status = "inactive"
		// case 3:
		// 	user.Status = "deleted"
		// }

		user.Status = consts.MapStatusInt[v.Status]

		result = append(result, user)
	}

	return &result
}
