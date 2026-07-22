package mapping

import "app/modules/admins/users/transport/requests"

func MapperCondListDatatableUser(req *requests.ListDatatableUserReq) map[string]interface{} {
	data := map[string]interface{}{}
	if req.SearchName != "" {
		data["full_name ILIKE ?"] = "%" + req.SearchName + "%"
	}
	return data
}
