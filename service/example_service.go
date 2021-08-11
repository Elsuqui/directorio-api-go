package service

type ExampleServiceInterface interface {
	FindAll() //[]entity.Video
	//Save(entity.Video) //entity.Video
	Save()
}

type ExampleService struct {
	//videos []entity.Video
}

func (service *ExampleService) FindAll() {
	//return service.videos
}

func (service *ExampleService) Save() {
	//service.videos = append(service.videos, entidad)
	//return entidad
}
