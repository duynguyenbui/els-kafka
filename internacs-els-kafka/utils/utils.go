package utils

import (
	"encoding/json"
	"fmt"
	"internacs-els-kafka/models"
)

func ConvertAfterToHotel(event models.DebeziumEvent) (*models.Hotel, error) {
	after := event.After

	var hotel models.Hotel

	if err := json.Unmarshal([]byte(after.AmenityGroups), &hotel.AmenityGroups); err != nil {
		return nil, fmt.Errorf("error unmarshalling amenity_groups: %w", err)
	}

	if err := json.Unmarshal([]byte(after.DescriptionStruct), &hotel.DescriptionStruct); err != nil {
		return nil, fmt.Errorf("error unmarshalling description_struct: %w", err)
	}

	if err := json.Unmarshal([]byte(after.Images), &hotel.Images); err != nil {
		return nil, fmt.Errorf("error unmarshalling images: %w", err)
	}

	if err := json.Unmarshal([]byte(after.PolicyStruct), &hotel.PolicyStruct); err != nil {
		return nil, fmt.Errorf("error unmarshalling policy_struct: %w", err)
	}

	if err := json.Unmarshal([]byte(after.RoomGroups), &hotel.RoomGroups); err != nil {
		return nil, fmt.Errorf("error unmarshalling room_groups: %w", err)
	}

	if err := json.Unmarshal([]byte(after.Region), &hotel.Region); err != nil {
		return nil, fmt.Errorf("error unmarshalling region: %w", err)
	}

	if err := json.Unmarshal([]byte(after.SerpFilters), &hotel.SerpFilters); err != nil {
		return nil, fmt.Errorf("error unmarshalling serp_filters: %w", err)
	}

	if err := json.Unmarshal([]byte(after.MetapolicyStruct), &hotel.MetapolicyStruct); err != nil {
		return nil, fmt.Errorf("error unmarshalling metapolicy_struct: %w", err)
	}

	if err := json.Unmarshal([]byte(after.Facts), &hotel.Facts); err != nil {
		return nil, fmt.Errorf("error unmarshalling facts: %w", err)
	}

	if err := json.Unmarshal([]byte(after.PaymentMethods), &hotel.PaymentMethods); err != nil {
		return nil, fmt.Errorf("error unmarshalling payment_methods: %w", err)
	}

	hotel.Address = after.Address
	hotel.CheckInTime = after.CheckInTime
	hotel.CheckOutTime = after.CheckOutTime
	hotel.ID = after.ID
	hotel.Kind = after.Kind
	hotel.Latitude = after.Latitude
	hotel.Longitude = after.Longitude
	hotel.Name = after.Name
	hotel.Phone = after.Phone
	hotel.PostalCode = after.PostalCode
	hotel.StarRating = after.StarRating
	hotel.Email = after.Email
	hotel.IsClosed = after.IsClosed
	hotel.IsGenderSpecificationRequired = after.IsGenderSpecificationRequired
	hotel.MetapolicyExtraInfo = after.MetapolicyExtraInfo
	hotel.StarCertificate = after.StarCertificate
	hotel.HotelChain = after.HotelChain
	hotel.FrontDeskTimeStart = after.FrontDeskTimeStart
	hotel.FrontDeskTimeEnd = after.FrontDeskTimeEnd
	hotel.SemanticVersion = after.SemanticVersion

	return &hotel, nil
}
