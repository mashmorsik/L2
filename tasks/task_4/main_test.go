package main

import (
	"reflect"
	"testing"
)

func TestAnagramDict(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "test_anagrams_dict_1",
			args: args{list: []string{"горбик", "Гробик", "керамит", "собака", "Метрика", "Материк", "Грибок"}},
			want: map[string][]string{"горбик": {"грибок", "гробик"}, "керамит": {"материк", "метрика"}},
		},
		{
			name: "test_anagrams_dict_2",
			args: args{list: []string{"Кластер", "нирвана", "Стрелка", "равнина", "Рванина", "Уголь", "сталкер"}},
			want: map[string][]string{"кластер": {"сталкер", "стрелка"}, "нирвана": {"равнина", "рванина"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnagramDict(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnagramDict() = %v, want %v", got, tt.want)
			}
		})
	}
}
