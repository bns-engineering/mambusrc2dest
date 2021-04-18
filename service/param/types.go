package param

//CityType - model city
type CityType struct {
	CityID string `json:"id" form:"id" query:"id"`
	City   string `json:"name" form:"name" query:"name"`
}

//DegreeType - model degree type
type DegreeType struct {
	DegreeID string `json:"id" form:"id" query:"id"`
	Degree   string `json:"name" form:"name" query:"name"`
}

//IncomeType - model income type
type IncomeType struct {
	IncomeID string `json:"id" form:"id" query:"id"`
	Income   string `json:"name" form:"name" query:"name"`
}

//IndustrialSectorType - model industrial sector type
type IndustrialSectorType struct {
	IndustrialSectorID string `json:"id" form:"id" query:"id"`
	IndustrialSector   string `json:"name" form:"name" query:"name"`
}

//JobTypesType - model job types
type JobTypesType struct {
	JobTypesID string `json:"id" form:"id" query:"id"`
	JobTypes   string `json:"name" form:"name" query:"name"`
}

//OpeningPurposeType - model opening purpose
type OpeningPurposeType struct {
	OpeningPurposeID string `json:"id" form:"id" query:"id"`
	OpeningPurpose   string `json:"name" form:"name" query:"name"`
}

//ProvinceType - model province type
type ProvinceType struct {
	ProvinceID string `json:"id" form:"id" query:"id"`
	Province   string `json:"name" form:"name" query:"name"`
}

//SourceIncomeType - model source income type
type SourceIncomeType struct {
	SourceIncomeID string `json:"id" form:"id" query:"id"`
	SourceIncome   string `json:"name" form:"name" query:"name"`
}
