package entity

import "gioui.org/layout"

type Updateable interface {
	Update(layout.Context)
}
