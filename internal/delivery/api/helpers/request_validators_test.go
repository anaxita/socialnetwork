package helpers

import (
	"synergycommunity/internal/domain"
	"testing"
)

// func TestFilterAllowedFields(t *testing.T) {
//	type args struct {
//		filters []restmodel.OptionsFilter
//		allowed map[string]struct{}
//	}
//	tests := []struct {
//		name string
//		args args
//		want []restmodel.OptionsFilter
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(
//			tt.name, func(t *testing.T) {
//				if got := FilterAllowedFields(
//					tt.args.filters, tt.args.allowed,
//				); !reflect.DeepEqual(got, tt.want) {
//					t.Errorf("FilterAllowedFields() = %v, want %v", got, tt.want)
//				}
//			},
//		)
//	}
// }

func TestHasAccess(t *testing.T) { // nolint:funlen
	t.Parallel()

	type args struct {
		got        []domain.Perm
		allowed    []domain.Perm
		disallowed []domain.Perm
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "all empty",
			args: args{
				got:        nil,
				allowed:    nil,
				disallowed: nil,
			},
			want: true,
		},
		{
			name: "allow edit positive",
			args: args{
				got:        []domain.Perm{domain.PermDelete, domain.PermEdit},
				allowed:    []domain.Perm{domain.PermEdit},
				disallowed: nil,
			},
			want: true,
		},
		{
			name: "allow edit negative",
			args: args{
				got:        []domain.Perm{domain.PermDelete, domain.PermAdministrate},
				allowed:    []domain.Perm{domain.PermEdit, domain.PermDelete, domain.PermWrite},
				disallowed: nil,
			},
			want: false,
		},
		{
			name: "allow edit with readonly negative",
			args: args{
				got: []domain.Perm{
					domain.PermDelete, domain.PermAdministrate, domain.PermNoWriting,
				},
				allowed:    []domain.Perm{domain.PermEdit, domain.PermDelete, domain.PermWrite},
				disallowed: []domain.Perm{domain.PermNoWriting},
			},
			want: false,
		},
		{
			name: "allow everyone except readonly positive",
			args: args{
				got: []domain.Perm{
					domain.PermDelete, domain.PermAdministrate,
				},
				allowed:    nil,
				disallowed: []domain.Perm{domain.PermNoWriting},
			},
			want: true,
		},
		{
			name: "allow everyone except readonly negative",
			args: args{
				got: []domain.Perm{
					domain.PermDelete, domain.PermNoWriting, domain.PermAdministrate,
				},
				allowed:    nil,
				disallowed: []domain.Perm{domain.PermNoWriting},
			},
			want: false,
		},
		{
			name: "allow with multiple permissions negative",
			args: args{
				got: []domain.Perm{
					domain.PermDelete, domain.PermNoWriting, domain.PermAdministrate,
				},
				allowed: []domain.Perm{
					domain.PermEdit,
					domain.PermDelete,
					domain.PermAdministrate,
					domain.PermWrite,
				},
				disallowed: nil,
			},
			want: false,
		},
		{
			name: "allow with multiple permissions positive",
			args: args{
				got: []domain.Perm{
					domain.PermDelete, domain.PermEdit, domain.PermAdministrate, domain.PermWrite,
				},
				allowed: []domain.Perm{
					domain.PermEdit,
					domain.PermDelete,
					domain.PermAdministrate,
					domain.PermWrite,
				},
				disallowed: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()
				if got := HasAccess(
					tt.args.got, tt.args.allowed, tt.args.disallowed,
				); got != tt.want {
					t.Errorf("HasAccess() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// func TestIsAllowAccess(t *testing.T) {
//	type args struct {
//		got     []entity.Perm
//		allowed []entity.Perm
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(
//			tt.name, func(t *testing.T) {
//				if got := IsAccessAllowed(tt.args.got, tt.args.allowed); got != tt.want {
//					t.Errorf("IsAccessAllowed() = %v, want %v", got, tt.want)
//				}
//			},
//		)
//	}
// }
//
// func TestIsDisallowAccess(t *testing.T) {
//	type args struct {
//		got        []entity.Perm
//		disallowed []entity.Perm
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(
//			tt.name, func(t *testing.T) {
//				if got := IsDisallowAccess(tt.args.got, tt.args.disallowed); got != tt.want {
//					t.Errorf("IsDisallowAccess() = %v, want %v", got, tt.want)
//				}
//			},
//		)
//	}
// }
//
// func TestValidateOptions(t *testing.T) {
//	type args struct {
//		o      *restmodel.Options
//		fields []string
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(
//			tt.name, func(t *testing.T) {
//				ValidateOptions(tt.args.o, tt.args.fields...)
//			},
//		)
//	}
// }
