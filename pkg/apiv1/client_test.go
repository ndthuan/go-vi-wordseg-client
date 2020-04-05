package apiv1_test

import (
	"github.com/ndthuan/go-vi-wordseg-client/pkg/apiv1"
	"os"
	"reflect"
	"testing"
)

func TestClient_Segment(t *testing.T) {
	c := apiv1.NewClient(validServiceHost())

	type args struct {
		text            string
		shouldSkipPunct bool
	}

	tests := []struct {
		host    string
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "SkipPunct=0",
			args:    args{text: "'Tổng tấn công' COVID!", shouldSkipPunct: false},
			want:    []string{"' Tổng_tấn_công ' COVID !"},
			wantErr: false,
		},
		{
			name:    "SkipPunct=1",
			args:    args{text: "Tổng tấn công COVID!", shouldSkipPunct: true},
			want:    []string{"Tổng_tấn_công COVID"},
			wantErr: false,
		},
		{
			name: "SkipPunct=1, Long text",
			args: args{text: `Lúc ấy, Tùng đã nghe nói đến dịch Covid-19 ở Trung Quốc, cướp đi sinh mạng của hàng nghìn bệnh nhân, có cả những người trẻ ban đầu không xuất hiện triệu chứng gì. Thế nhưng, nghĩ đi nghĩ lại, anh vẫn không hiểu được mình có thể bị nhiễm bệnh bằng cách nào. Bố mẹ từ Việt Nam nhiều lần gọi sang Hàn Quốc thúc giục con trai về nước. Tùng quyết định về, nhưng anh không nói đang triệu chứng rát họng cho bố mẹ biết, "sợ họ lo lắng quá".`, shouldSkipPunct: true},
			want: []string{
				"Lúc ấy Tùng đã nghe nói đến dịch Covid-19 ở Trung_Quốc cướp đi sinh_mạng của hàng nghìn bệnh_nhân có cả những người trẻ ban_đầu không xuất_hiện triệu_chứng gì",
				"Thế nhưng nghĩ_đi_nghĩ_lại anh vẫn không hiểu được mình có_thể bị nhiễm_bệnh bằng cách nào",
				"Bố_mẹ từ Việt_Nam nhiều lần gọi sang Hàn_Quốc thúc_giục con trai về nước",
				"Tùng quyết_định về nhưng anh không nói đang triệu_chứng rát họng cho bố_mẹ biết sợ họ lo_lắng quá",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Segment(tt.args.text, tt.args.shouldSkipPunct)
			if (err != nil) != tt.wantErr {
				t.Errorf("Segment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Sentences, tt.want) {
				t.Errorf("Segment() got = %v, want %v", got.Sentences, tt.want)
			}
		})
	}
}

func TestClient_Tag(t *testing.T) {
	c := apiv1.NewClient(validServiceHost())

	tests := []struct {
		name    string
		text    string
		want    *apiv1.TaggingResult
		wantErr bool
	}{
		{
			name: "",
			text: "Việt Nam tổng tấn công COVID!",
			want: &apiv1.TaggingResult{Sentences: [][]apiv1.Word{
				{
					apiv1.Word{
						Form: "Việt_Nam",
						Pos:  "Np",
						Ner:  "B-PER",
						Dep:  "sub",
					},
					apiv1.Word{
						Form: "tổng_tấn_công",
						Pos:  "V",
						Ner:  "O",
						Dep:  "root",
					},
					apiv1.Word{
						Form: "COVID",
						Pos:  "Ny",
						Ner:  "O",
						Dep:  "dob",
					},
					apiv1.Word{
						Form: "!",
						Pos:  "CH",
						Ner:  "O",
						Dep:  "punct",
					},
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Tag(tt.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tag() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func validServiceHost() string {
	h := os.Getenv("VALID_SERVICE_HOST")

	if len(h) < 1 {
		h = "http://segmenterv1:8080"
	}

	return h
}
