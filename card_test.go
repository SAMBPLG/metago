package metago

import (
	"testing"
)

func TestCardF_String(t *testing.T) {
	tests := []struct {
		name string
		c    CardF
		want string
	}{
		{
			name: "ShouldReturnCorrectString",
			c:    CardFilterArchived,
			want: "archived",
		},
		{
			name: "ShouldReturnCorrectString1",
			c:    CardFilterAll,
			want: "all",
		},
		{
			name: "ShouldReturnCorrectString1",
			c:    CardFilterUsingModel,
			want: "using_model",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("CardF.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardF_EnumIndex(t *testing.T) {
	tests := []struct {
		name string
		c    CardF
		want uint8
	}{
		{
			name: "ShouldReturnCorrectIndex",
			c:    CardFilterAll,
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.EnumIndex(); got != tt.want {
				t.Errorf("CardF.EnumIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
