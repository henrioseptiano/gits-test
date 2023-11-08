package repository

var maem = &MainAppEntityMock{}
var mainAppEntity = NewAppRepository(maem)

//var mainAppEntity = &AppRepository{DB: maem}
