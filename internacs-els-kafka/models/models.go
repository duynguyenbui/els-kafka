package models

type DebeziumEvent struct {
	Before interface{} `json:"before"`
	After  struct {
		Address                       string      `json:"address"`
		AmenityGroups                 string      `json:"amenity_groups"`
		CheckInTime                   string      `json:"check_in_time"`
		CheckOutTime                  string      `json:"check_out_time"`
		DescriptionStruct             string      `json:"description_struct"`
		ID                            string      `json:"id"`
		Images                        string      `json:"images"`
		Kind                          string      `json:"kind"`
		Latitude                      float64     `json:"latitude"`
		Longitude                     float64     `json:"longitude"`
		Name                          string      `json:"name"`
		Phone                         string      `json:"phone"`
		PolicyStruct                  string      `json:"policy_struct"`
		PostalCode                    string      `json:"postal_code"`
		RoomGroups                    string      `json:"room_groups"`
		Region                        string      `json:"region"`
		StarRating                    int         `json:"star_rating"`
		Email                         interface{} `json:"email"`
		SerpFilters                   string      `json:"serp_filters"`
		IsClosed                      bool        `json:"is_closed"`
		IsGenderSpecificationRequired bool        `json:"is_gender_specification_required"`
		MetapolicyStruct              string      `json:"metapolicy_struct"`
		MetapolicyExtraInfo           string      `json:"metapolicy_extra_info"`
		StarCertificate               interface{} `json:"star_certificate"`
		Facts                         string      `json:"facts"`
		PaymentMethods                string      `json:"payment_methods"`
		HotelChain                    string      `json:"hotel_chain"`
		FrontDeskTimeStart            interface{} `json:"front_desk_time_start"`
		FrontDeskTimeEnd              interface{} `json:"front_desk_time_end"`
		SemanticVersion               int         `json:"semantic_version"`
	} `json:"after"`
	Source struct {
		Version   string      `json:"version"`
		Connector string      `json:"connector"`
		Name      string      `json:"name"`
		TsMs      int64       `json:"ts_ms"`
		Snapshot  string      `json:"snapshot"`
		Db        string      `json:"db"`
		Sequence  string      `json:"sequence"`
		Schema    string      `json:"schema"`
		Table     string      `json:"table"`
		TxID      int         `json:"txId"`
		Lsn       int         `json:"lsn"`
		Xmin      interface{} `json:"xmin"`
	} `json:"source"`
	Op          string      `json:"op"`
	TsMs        int64       `json:"ts_ms"`
	Transaction interface{} `json:"transaction"`
}

type Hotel struct {
	Address       string `json:"address"`
	AmenityGroups []struct {
		Amenities []string `json:"amenities"`
		GroupName string   `json:"group_name"`
	} `json:"amenity_groups"`
	CheckInTime       string `json:"check_in_time"`
	CheckOutTime      string `json:"check_out_time"`
	DescriptionStruct []struct {
		Paragraphs []string `json:"paragraphs"`
		Title      string   `json:"title"`
	} `json:"description_struct"`
	ID           string        `json:"id"`
	Images       []interface{} `json:"images"`
	Kind         string        `json:"kind"`
	Latitude     float64       `json:"latitude"`
	Longitude    float64       `json:"longitude"`
	Name         string        `json:"name"`
	Phone        string        `json:"phone"`
	PolicyStruct []interface{} `json:"policy_struct"`
	PostalCode   string        `json:"postal_code"`
	RoomGroups   []struct {
		RoomGroupID   int           `json:"room_group_id"`
		Images        []interface{} `json:"images"`
		Name          string        `json:"name"`
		RoomAmenities []string      `json:"room_amenities"`
		RgExt         struct {
			Class    int `json:"class"`
			Quality  int `json:"quality"`
			Sex      int `json:"sex"`
			Bathroom int `json:"bathroom"`
			Bedding  int `json:"bedding"`
			Family   int `json:"family"`
			Capacity int `json:"capacity"`
			Club     int `json:"club"`
			Bedrooms int `json:"bedrooms"`
			Balcony  int `json:"balcony"`
			Floor    int `json:"floor"`
			View     int `json:"view"`
		} `json:"rg_ext"`
		NameStruct struct {
			Bathroom    interface{} `json:"bathroom"`
			BeddingType interface{} `json:"bedding_type"`
			MainName    string      `json:"main_name"`
		} `json:"name_struct"`
	} `json:"room_groups"`
	Region struct {
		ID          int         `json:"id"`
		CountryCode string      `json:"country_code"`
		Iata        interface{} `json:"iata"`
		Name        string      `json:"name"`
		Type        string      `json:"type"`
	} `json:"region"`
	StarRating                    int         `json:"star_rating"`
	Email                         interface{} `json:"email"`
	SerpFilters                   []string    `json:"serp_filters"`
	IsClosed                      bool        `json:"is_closed"`
	IsGenderSpecificationRequired bool        `json:"is_gender_specification_required"`
	MetapolicyStruct              struct {
		Internet     []interface{} `json:"internet"`
		Meal         []interface{} `json:"meal"`
		ChildrenMeal []interface{} `json:"children_meal"`
		ExtraBed     []interface{} `json:"extra_bed"`
		Cot          []interface{} `json:"cot"`
		Pets         []interface{} `json:"pets"`
		Shuttle      []interface{} `json:"shuttle"`
		Parking      []interface{} `json:"parking"`
		Children     []interface{} `json:"children"`
		Visa         struct {
			VisaSupport string `json:"visa_support"`
		} `json:"visa"`
		Deposit []interface{} `json:"deposit"`
		NoShow  struct {
			Availability string      `json:"availability"`
			Time         interface{} `json:"time"`
			DayPeriod    string      `json:"day_period"`
		} `json:"no_show"`
		AddFee          []interface{} `json:"add_fee"`
		CheckInCheckOut []interface{} `json:"check_in_check_out"`
	} `json:"metapolicy_struct"`
	MetapolicyExtraInfo string      `json:"metapolicy_extra_info"`
	StarCertificate     interface{} `json:"star_certificate"`
	Facts               struct {
		FloorsNumber  interface{} `json:"floors_number"`
		RoomsNumber   interface{} `json:"rooms_number"`
		YearBuilt     interface{} `json:"year_built"`
		YearRenovated interface{} `json:"year_renovated"`
		Electricity   struct {
			Frequency []int    `json:"frequency"`
			Voltage   []int    `json:"voltage"`
			Sockets   []string `json:"sockets"`
		} `json:"electricity"`
	} `json:"facts"`
	PaymentMethods     []interface{} `json:"payment_methods"`
	HotelChain         string        `json:"hotel_chain"`
	FrontDeskTimeStart interface{}   `json:"front_desk_time_start"`
	FrontDeskTimeEnd   interface{}   `json:"front_desk_time_end"`
	SemanticVersion    int           `json:"semantic_version"`
}