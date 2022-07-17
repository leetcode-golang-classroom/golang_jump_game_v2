package sol

import "testing"

func BenchmarkTestDP(b *testing.B) {
	nums := []int{2, 3, 0, 1, 4}
	for idx := 0; idx < b.N; idx++ {
		jump_dp(nums)
	}
}
func Test_jump_dp(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [2,3,1,1,4]",
			args: args{nums: []int{2, 3, 1, 1, 4}},
			want: 2,
		},
		{
			name: "nums = [2,3,0,1,4]",
			args: args{nums: []int{2, 3, 0, 1, 4}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump_dp(tt.args.nums); got != tt.want {
				t.Errorf("jump_dp() = %v, want %v", got, tt.want)
			}
		})
	}
}
