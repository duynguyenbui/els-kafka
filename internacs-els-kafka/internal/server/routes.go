package server

import (
	"context"
	"encoding/json"
	"fmt"
	"internacs-els-kafka/global"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var index = os.Getenv("ELASTIC_INDEX")

func (s *Server) RegisterRoutes() http.Handler {
	var r *gin.Engine

	if os.Getenv("APP_ENV") == "local" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"alive": "heaththy",
		})
	})

	r.POST("/els/search", s.Search)

	r.POST("/db/create", s.Create)

	return r
}

func (s *Server) Search(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	defer c.Request.Body.Close()

	var customQuery map[string]interface{}
	if err := json.Unmarshal(body, &customQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	client := global.Els

	queryBody, err := json.Marshal(customQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error marshalling query: %s", err)})
		return
	}

	queryString := string(queryBody)

	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(index),
		client.Search.WithBody(strings.NewReader(queryString)),
		client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting response: %s", err)})
		return
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error parsing the response body: %s", err)})
		return
	}

	c.JSON(http.StatusOK, result)
}

func randomAddress() string {
	addresses := []string{
		"41/53 Mechnikov Street, Rostov-on-Don",
		"123 Main Street, New York",
		"456 Elm Street, Chicago",
		"789 Maple Avenue, Los Angeles",
		"1010 Oak Lane, San Francisco",
	}
	rand.Seed(time.Now().UnixNano())
	return addresses[rand.Intn(len(addresses))]
}

func (s *Server) Create(c *gin.Context) {
	address := randomAddress()
	insertQuery := `
		INSERT INTO hotels (
			address,
			amenity_groups,
			check_in_time,
			check_out_time,
			description_struct,
			id,
			images,
			kind,
			latitude,
			longitude,
			name,
			phone,
			policy_struct,
			postal_code,
			room_groups,
			region,
			star_rating,
			email,
			serp_filters,
			is_closed,
			is_gender_specification_required,
			metapolicy_struct,
			metapolicy_extra_info,
			star_certificate,
			facts,
			payment_methods,
			hotel_chain,
			front_desk_time_start,
			front_desk_time_end,
			semantic_version
		)
		VALUES (
			$1, -- Random address
			'[{"amenities":["ATM","Shopping on site","Elevator/lift"],"group_name":"General"},{"amenities":["Cafe"],"group_name":"Meals"},{"amenities":["Wi-Fi"],"group_name":"Internet"},{"amenities":["Parking"],"group_name":"Parking"},{"amenities":["Children''s playground","Kids'' TV Networks"],"group_name":"Kids"}]',
			'15:00:00',
			'11:00:00',
			'[{"paragraphs":["The best kind of vacation is when you come somewhere new and it feels like home: apartment «Apartments on 41/53 Mechnikova Street» is located in Rostov-on-Don. This apartment is located 2 km from the city center. You can take a walk and explore the neighbourhood area of the apartment. Places nearby: Rostov Zoo, Rostov Gorky Park and Don River embankment."],"title":"Location"},{"paragraphs":["Have a cup of coffee in the cafe and, who knows, maybe it’s going to be the best one in the city. Wi-Fi on the territory will help you stay on-line. If you travel by car, you can park in a parking zone. Accessible for guests with disabilities: the elevator helps them to go to the highest floors.","There are other services available for the guests of the apartment. For example, an ATM."],"title":"At the apartment"}]',
			'apartments_on_4153_mechnikova_street_3',
			'[]',
			'Apartment',
			47.23981475830078,
			39.68932342529297,
			'Apartments on 41/53 Mechnikova Street',
			'+79675554112',
			'[]',
			'344012',
			'[{"room_group_id":18395703,"images":[],"name":"1 Bedroom Standard Apartment","room_amenities":["private-bathroom"],"rg_ext":{"class":6,"quality":2,"sex":0,"bathroom":2,"bedding":0,"family":0,"capacity":0,"club":0,"bedrooms":1,"balcony":0,"floor":0,"view":0},"name_struct":{"bathroom":null,"bedding_type":null,"main_name":"1 Bedroom Standard Apartment"}},{"room_group_id":18395751,"images":[],"name":"1 Bedroom Superior Apartment","room_amenities":["private-bathroom"],"rg_ext":{"class":6,"quality":5,"sex":0,"bathroom":2,"bedding":0,"family":0,"capacity":0,"club":0,"bedrooms":1,"balcony":0,"floor":0,"view":0},"name_struct":{"bathroom":null,"bedding_type":null,"main_name":"1 Bedroom Superior Apartment"}},{"room_group_id":26,"images":[],"name":"Studio","room_amenities":["air-conditioning","kitchen","private-bathroom","shower","wi-fi"],"rg_ext":{"class":7,"quality":0,"sex":0,"bathroom":2,"bedding":0,"family":0,"capacity":0,"club":0,"bedrooms":0,"balcony":0,"floor":0,"view":0},"name_struct":{"bathroom":null,"bedding_type":null,"main_name":"Studio"}},{"room_group_id":18395672,"images":[],"name":"1 Bedroom Studio","room_amenities":["air-conditioning","kitchen","private-bathroom","shower","wi-fi"],"rg_ext":{"class":7,"quality":0,"sex":0,"bathroom":2,"bedding":0,"family":0,"capacity":0,"club":0,"bedrooms":1,"balcony":0,"floor":0,"view":0},"name_struct":{"bathroom":null,"bedding_type":null,"main_name":"1 Bedroom Studio"}}]',
			'{"id":3028,"country_code":"RU","iata":"ROV","name":"Rostov-on-Don","type":"City"}',
			0,
			NULL,
			'["has_internet","has_parking","has_kids","has_meal"]',
			FALSE,
			FALSE,
			'{"internet":[],"meal":[],"children_meal":[],"extra_bed":[],"cot":[],"pets":[],"shuttle":[],"parking":[],"children":[],"visa":{"visa_support":"unspecified"},"deposit":[],"no_show":{"availability":"unspecified","time":null,"day_period":"unspecified"},"add_fee":[],"check_in_check_out":[]}',
			'Russian citizens must have an original Russian passport upon arrival.',
			NULL,
			'{"floors_number":null,"rooms_number":null,"year_built":null,"year_renovated":null,"electricity":{"frequency":[50],"voltage":[220],"sockets":["c","f"]}}',
			'[]',
			'No chain',
			NULL,
			NULL,
			0
		)
	`

	_, err := global.Pdb.Exec(insertQuery, address)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to execute query"})
		return
	}

	c.JSON(200, gin.H{"message": "Query executed successfully"})
}
