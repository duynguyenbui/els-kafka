package models

type Hotels struct {
	Address           string `json:"address"`
	AmenityGroups     []any  `json:"amenity_groups"`
	CheckInTime       string `json:"check_in_time"`
	CheckOutTime      string `json:"check_out_time"`
	DescriptionStruct []struct {
		Paragraphs []string `json:"paragraphs"`
		Title      string   `json:"title"`
	} `json:"description_struct"`
	ID           string  `json:"id"`
	Images       []any   `json:"images"`
	Kind         string  `json:"kind"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Name         string  `json:"name"`
	Phone        any     `json:"phone"`
	PolicyStruct []any   `json:"policy_struct"`
	PostalCode   string  `json:"postal_code"`
	RoomGroups   []any   `json:"room_groups"`
	Region       struct {
		ID          int    `json:"id"`
		CountryCode string `json:"country_code"`
		Iata        string `json:"iata"`
		Name        string `json:"name"`
		Type        string `json:"type"`
	} `json:"region"`
	StarRating                    int   `json:"star_rating"`
	Email                         any   `json:"email"`
	SerpFilters                   []any `json:"serp_filters"`
	IsClosed                      bool  `json:"is_closed"`
	IsGenderSpecificationRequired bool  `json:"is_gender_specification_required"`
	MetapolicyStruct              struct {
		Internet     []any `json:"internet"`
		Meal         []any `json:"meal"`
		ChildrenMeal []any `json:"children_meal"`
		ExtraBed     []any `json:"extra_bed"`
		Cot          []any `json:"cot"`
		Pets         []any `json:"pets"`
		Shuttle      []any `json:"shuttle"`
		Parking      []any `json:"parking"`
		Children     []any `json:"children"`
		Visa         struct {
			VisaSupport string `json:"visa_support"`
		} `json:"visa"`
		Deposit []any `json:"deposit"`
		NoShow  struct {
			Availability string `json:"availability"`
			Time         any    `json:"time"`
			DayPeriod    string `json:"day_period"`
		} `json:"no_show"`
		AddFee          []any `json:"add_fee"`
		CheckInCheckOut []any `json:"check_in_check_out"`
	} `json:"metapolicy_struct"`
	MetapolicyExtraInfo any `json:"metapolicy_extra_info"`
	StarCertificate     any `json:"star_certificate"`
	Facts               struct {
		FloorsNumber  any `json:"floors_number"`
		RoomsNumber   int `json:"rooms_number"`
		YearBuilt     any `json:"year_built"`
		YearRenovated any `json:"year_renovated"`
		Electricity   struct {
			Frequency []int    `json:"frequency"`
			Voltage   []int    `json:"voltage"`
			Sockets   []string `json:"sockets"`
		} `json:"electricity"`
	} `json:"facts"`
	PaymentMethods     []any  `json:"payment_methods"`
	HotelChain         string `json:"hotel_chain"`
	FrontDeskTimeStart any    `json:"front_desk_time_start"`
	FrontDeskTimeEnd   any    `json:"front_desk_time_end"`
	SemanticVersion    int    `json:"semantic_version"`
}

type Regions struct {
	CountryName struct {
		Ar   string `json:"ar"`
		Bg   string `json:"bg"`
		Cs   string `json:"cs"`
		De   string `json:"de"`
		El   string `json:"el"`
		En   string `json:"en"`
		Es   string `json:"es"`
		Fr   string `json:"fr"`
		He   string `json:"he"`
		Hu   string `json:"hu"`
		It   string `json:"it"`
		Ja   string `json:"ja"`
		Kk   string `json:"kk"`
		Ko   string `json:"ko"`
		Nl   string `json:"nl"`
		Pl   string `json:"pl"`
		Pt   string `json:"pt"`
		PtPT string `json:"pt_PT"`
		Ro   string `json:"ro"`
		Ru   string `json:"ru"`
		Sq   string `json:"sq"`
		Sr   string `json:"sr"`
		Th   string `json:"th"`
		Tr   string `json:"tr"`
		Vi   string `json:"vi"`
		ZhCN string `json:"zh_CN"`
	} `json:"country_name"`
	CountryCode string `json:"country_code"`
	Center      struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	} `json:"center"`
	Hotels []string `json:"hotels"`
	Iata   string   `json:"iata"`
	ID     int      `json:"id"`
	Type   string   `json:"type"`
	Name   struct {
		Ar   string `json:"ar"`
		Bg   string `json:"bg"`
		Cs   string `json:"cs"`
		De   string `json:"de"`
		El   string `json:"el"`
		En   string `json:"en"`
		Es   string `json:"es"`
		Fr   string `json:"fr"`
		He   string `json:"he"`
		Hu   string `json:"hu"`
		It   string `json:"it"`
		Ja   string `json:"ja"`
		Kk   string `json:"kk"`
		Ko   string `json:"ko"`
		Nl   string `json:"nl"`
		Pl   string `json:"pl"`
		Pt   string `json:"pt"`
		PtPT string `json:"pt_PT"`
		Ro   string `json:"ro"`
		Ru   string `json:"ru"`
		Sq   string `json:"sq"`
		Sr   string `json:"sr"`
		Th   string `json:"th"`
		Tr   string `json:"tr"`
		Vi   string `json:"vi"`
		ZhCN string `json:"zh_CN"`
	} `json:"name"`
}
