package profile

import (
	"sort"
)

func (p *Profile) Tidy() {
	sort.Slice(p.ApplicationVisibilities, func(i, j int) bool {
		return p.ApplicationVisibilities[i].Application.Text < p.ApplicationVisibilities[j].Application.Text
	})
	sort.Slice(p.ClassAccesses, func(i, j int) bool {
		return p.ClassAccesses[i].ApexClass.Text < p.ClassAccesses[j].ApexClass.Text
	})
	p.FieldPermissions.Tidy()
	p.ObjectPermissions.Tidy()
	sort.Slice(p.PageAccesses, func(i, j int) bool {
		return p.PageAccesses[i].ApexPage.Text < p.PageAccesses[j].ApexPage.Text
	})
	sort.Slice(p.TabVisibilities, func(i, j int) bool {
		return p.TabVisibilities[i].Tab < p.TabVisibilities[j].Tab
	})
	sort.Slice(p.RecordTypeVisibilities, func(i, j int) bool {
		return p.RecordTypeVisibilities[i].RecordType.Text < p.RecordTypeVisibilities[j].RecordType.Text
	})
	sort.Slice(p.FlowAccesses, func(i, j int) bool {
		return p.FlowAccesses[i].Flow.Text < p.FlowAccesses[j].Flow.Text
	})
	sort.Slice(p.UserPermissions, func(i, j int) bool {
		return p.UserPermissions[i].Name.Text < p.UserPermissions[j].Name.Text
	})
}

func (op ObjectPermissionsList) Tidy() {
	sort.Slice(op, func(i, j int) bool {
		return op[i].Object.Text < op[j].Object.Text
	})
}

func (fp FieldPermissionsList) Tidy() {
	sort.Slice(fp, func(i, j int) bool {
		return fp[i].Field.Text < fp[j].Field.Text
	})
}
