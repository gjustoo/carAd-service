package carAdService

import (
	carAdModel "fm-scrapper-go/app/model/carAd"
	carAdRepo "fm-scrapper-go/app/repo/carAd"
	"log"
)

func GetAllCarAds() []carAdModel.CarAd {

	posts, err := carAdRepo.GetAllCarAds()
	if err != nil {
		log.Println("Error retrieving all carAds ", err)
		return nil
	}
	return posts
}

func GetCarAd(id uint64) carAdModel.CarAd {

	post, err := carAdRepo.GetCarAd(id)
	if err != nil {
		log.Println("Error retrieving carAd : ", id, err)
		return carAdModel.CarAd{}
	}
	return post
}

func CreateCarAd(post carAdModel.CarAd) {

	err := carAdRepo.CreateCarAd(post)
	if err != nil {
		log.Println("Error creating carAd : ", post, err)
		return
	}

	log.Print("CarAd created : ", post)

}

func DeleteCarAd(id uint64) {

	err := carAdRepo.DeleteCarAd(id)
	if err != nil {
		log.Println("Error deleting carAd : ", id, err)
		return
	}

	log.Print("CarAd deleted : ", id)
}
