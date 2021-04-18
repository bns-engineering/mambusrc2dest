package param

import (
	"context"

	"github.com/bns-engineering/mambusrc2des/tcommon/external"
	"github.com/bns-engineering/mambusrc2des/tcommon/logging"
	"github.com/bns-engineering/mambusrc2des/tcommon/lookup"
)

//GetCity - get list city based on province ID
func GetCity(partnerID, requestTimestamp, signature, provinceID string) (listcity []CityType, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"partnerID":        partnerID,
		"requestTimestamp": requestTimestamp,
		"signature":        signature,
		"provinceID":       provinceID,
	})

	listcity = []CityType{}
	responseCode = lookup.GENERALERROR

	response, err := external.ParamClient.City(context.Background(), partnerID, requestTimestamp, signature, provinceID)
	if err != nil {
		logger.Error(err, "error get city from grpc")
	}

	if response == nil {
		responseCode = lookup.SERVICEERROR
		return
	}

	responseCode = response.GetResponseCode()

	if responseCode == lookup.SUCCESS {
		if len(response.City) > 0 {
			for _, value := range response.City {
				listcity = append(listcity, CityType{
					CityID: value.CityID,
					City:   value.City,
				})
			}
		}
	}

	return
}

//GetDegree - get list all degree
func GetDegree(partnerID, requestTimestamp, signature string) (listDegreeType []DegreeType, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"partnerID":        partnerID,
		"requestTimestamp": requestTimestamp,
		"signature":        signature,
	})

	listDegreeType = []DegreeType{}
	responseCode = lookup.GENERALERROR

	response, err := external.ParamClient.Degree(context.Background(), partnerID, requestTimestamp, signature)
	if err != nil {
		logger.Error(err, "error get city from grpc")
	}

	if response == nil {
		responseCode = lookup.SERVICEERROR
		return
	}

	responseCode = response.GetResponseCode()

	if responseCode == lookup.SUCCESS {
		if len(response.Degree) > 0 {
			for _, value := range response.Degree {
				listDegreeType = append(listDegreeType, DegreeType{
					DegreeID: value.DegreeID,
					Degree:   value.Degree,
				})
			}
		}
	}

	return
}

//GetIncome - get list all income
func GetIncome(partnerID, requestTimestamp, signature string) (listIncomeType []IncomeType, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"partnerID":        partnerID,
		"requestTimestamp": requestTimestamp,
		"signature":        signature,
	})

	listIncomeType = []IncomeType{}
	responseCode = lookup.GENERALERROR

	response, err := external.ParamClient.Income(context.Background(), partnerID, requestTimestamp, signature)
	if err != nil {
		logger.Error(err, "error get income from grpc")
	}

	if response == nil {
		responseCode = lookup.SERVICEERROR
		return
	}

	responseCode = response.GetResponseCode()

	if responseCode == lookup.SUCCESS {
		if len(response.Income) > 0 {
			for _, value := range response.Income {
				listIncomeType = append(listIncomeType, IncomeType{
					IncomeID: value.IncomeID,
					Income:   value.Income,
				})
			}
		}
	}

	return
}

//GetIndustrialSector - get list all industrial sector
func GetIndustrialSector(partnerID, requestTimestamp, signature string) (listIndustrialSectorType []IndustrialSectorType, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"partnerID":        partnerID,
		"requestTimestamp": requestTimestamp,
		"signature":        signature,
	})

	listIndustrialSectorType = []IndustrialSectorType{}
	responseCode = lookup.GENERALERROR

	response, err := external.ParamClient.IndustrialSector(context.Background(), partnerID, requestTimestamp, signature)
	if err != nil {
		logger.Error(err, "error get industrial sector from grpc")
	}

	if response == nil {
		responseCode = lookup.SERVICEERROR
		return
	}

	responseCode = response.GetResponseCode()

	if responseCode == lookup.SUCCESS {
		if len(response.IndustrialSector) > 0 {
			for _, value := range response.IndustrialSector {
				listIndustrialSectorType = append(listIndustrialSectorType, IndustrialSectorType{
					IndustrialSectorID: value.IndustrialSectorID,
					IndustrialSector:   value.IndustrialSector,
				})
			}
		}
	}

	return
}

//GetJobTypes - get list all job types
func GetJobTypes(partnerID, requestTimestamp, signature string) (listJobTypesType []JobTypesType, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"partnerID":        partnerID,
		"requestTimestamp": requestTimestamp,
		"signature":        signature,
	})

	listJobTypesType = []JobTypesType{}
	responseCode = lookup.GENERALERROR

	response, err := external.ParamClient.JobTypes(context.Background(), partnerID, requestTimestamp, signature)
	if err != nil {
		logger.Error(err, "error get job types from grpc")
	}

	if response == nil {
		responseCode = lookup.SERVICEERROR
		return
	}

	responseCode = response.GetResponseCode()

	if responseCode == lookup.SUCCESS {
		if len(response.JobTypes) > 0 {
			for _, value := range response.JobTypes {
				listJobTypesType = append(listJobTypesType, JobTypesType{
					JobTypesID: value.JobTypeID,
					JobTypes:   value.JobType,
				})
			}
		}
	}

	return
}

//OpeningPurpose - get list all opening purpose
func OpeningPurpose(partnerID, requestTimestamp, signature string) (listOpeningPurposeType []OpeningPurposeType, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"partnerID":        partnerID,
		"requestTimestamp": requestTimestamp,
		"signature":        signature,
	})

	listOpeningPurposeType = []OpeningPurposeType{}
	responseCode = lookup.GENERALERROR

	response, err := external.ParamClient.OpeningPurpose(context.Background(), partnerID, requestTimestamp, signature)
	if err != nil {
		logger.Error(err, "error get opening purpose from grpc")
	}

	if response == nil {
		responseCode = lookup.SERVICEERROR
		return
	}

	responseCode = response.GetResponseCode()

	if responseCode == lookup.SUCCESS {
		if len(response.OpeningPurpose) > 0 {
			for _, value := range response.OpeningPurpose {
				listOpeningPurposeType = append(listOpeningPurposeType, OpeningPurposeType{
					OpeningPurposeID: value.OpeningPurposeID,
					OpeningPurpose:   value.OpeningPurpose,
				})
			}
		}
	}

	return
}

//GetProvince - get list all province
func GetProvince(partnerID, requestTimestamp, signature string) (listProvinceType []ProvinceType, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"partnerID":        partnerID,
		"requestTimestamp": requestTimestamp,
		"signature":        signature,
	})

	listProvinceType = []ProvinceType{}
	responseCode = lookup.GENERALERROR

	response, err := external.ParamClient.Province(context.Background(), partnerID, requestTimestamp, signature)
	if err != nil {
		logger.Error(err, "error get province from grpc")
	}

	if response == nil {
		responseCode = lookup.SERVICEERROR
		return
	}

	responseCode = response.GetResponseCode()

	if responseCode == lookup.SUCCESS {
		if len(response.Province) > 0 {
			for _, value := range response.Province {
				listProvinceType = append(listProvinceType, ProvinceType{
					ProvinceID: value.ProvinceID,
					Province:   value.Province,
				})
			}
		}
	}

	return
}

//GetSourceIncome - get list all source incomes
func GetSourceIncome(partnerID, requestTimestamp, signature string) (listSourceIncomeType []SourceIncomeType, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"partnerID":        partnerID,
		"requestTimestamp": requestTimestamp,
		"signature":        signature,
	})

	listSourceIncomeType = []SourceIncomeType{}
	responseCode = lookup.GENERALERROR

	response, err := external.ParamClient.SourceIncome(context.Background(), partnerID, requestTimestamp, signature)
	if err != nil {
		logger.Error(err, "error get source income from grpc")
	}

	if response == nil {
		responseCode = lookup.SERVICEERROR
		return
	}

	responseCode = response.GetResponseCode()

	if responseCode == lookup.SUCCESS {
		if len(response.SourceIncome) > 0 {
			for _, value := range response.SourceIncome {
				listSourceIncomeType = append(listSourceIncomeType, SourceIncomeType{
					SourceIncomeID: value.SourceIncomeID,
					SourceIncome:   value.SourceIncome,
				})
			}
		}
	}

	return
}
