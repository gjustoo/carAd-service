package carAdModel

import (
	"time"
)

// CarAd structure for our blog
type CarAd struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Km       string    `json:"km"`
	Location string    `json:"location"`
	Date     time.Time `bson:"date,omitempty"`
	Price    string    `json:"price"`
	Url      string    `json:"url"`
	Img_data []byte    `json:"img_data"`
}
