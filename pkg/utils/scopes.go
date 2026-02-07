package utils

// func CurOrganization(r *http.Request) func(db *gorm.DB) *gorm.DB {
//   return func (db *gorm.DB) *gorm.DB {
//     org := r.Query("org")

//     if org != "" {
//       var organization Organization
//       if db.Session(&Session{}).First(&organization, "name = ?", org).Error == nil {
//         return db.Where("org_id = ?", organization.ID)
//       }
//     }

//     db.AddError("invalid organization")
//     return db
//   }
// }

//TODO: create DB scope for checking if the current user can update / delete the cart
// func CurrentUser(userID uint) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		var
// 	}
// }
