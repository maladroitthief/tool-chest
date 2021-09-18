package abstract_factory

import (
	"reflect"
	"testing"
)

func TestNewRoom(t *testing.T) {
	type args struct {
		f RoomFactory
	}
	tests := []struct {
		name string
		args args
		want *Room
	}{
		{
			name: "Boss Room Factory",
			args: args{
				f: &BossRoomFactory{},
			},
			want: &Room{
				Enemy:    "Dragon",
				Item:     "Heart container",
				Reward:   "Potions",
				Obstacle: "Rough terrain",
			},
		},
		{
			name: "Secret Room Factory",
			args: args{
				f: &SecretRoomFactory{},
			},
			want: &Room{
				Enemy:    "",
				Item:     "Coins",
				Reward:   "Skeleton key",
				Obstacle: "Hidden entrance",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRoom(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}
