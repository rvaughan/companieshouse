/*
Golang Companies House REST service API
Copyright (C) 2016-2017, Balkan C & T OOD

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package testhandlers

var Companies = map[string]string{
	"12345678": `{
	  "company_status": "active",
	  "jurisdiction": "england-wales",
	  "type": "ltd",
	  "has_insolvency_history": false,
	  "accounts": {
	    "overdue": false,
	    "next_made_up_to": "2016-12-31",
	    "last_accounts": {
	      "type": "null"
	    },
	    "accounting_reference_date": {
	      "day": "31",
	      "month": "12"
	    },
	    "next_due": "2017-09-30"
	  },
	  "date_of_creation": "2016-01-01",
	  "annual_return": {},
	  "company_name": "TEST LTD",
	  "company_number": "12345678",
	  "undeliverable_registered_office_address": false,
	  "etag": "d97f19ae994f1a0609a0cc7d00b6e961aa7a0311",
	  "registered_office_address": {
	    "country": "United Kingdom",
	    "address_line_1": "12 Test Road",
	    "postal_code": "TS1 3ST",
	    "locality": "London",
	    "region": "Greater-London"
	  },
	  "has_charges": false,
	  "confirmation_statement": {
	    "next_made_up_to": "2017-02-10",
	    "next_due": "2017-02-24"
	  },
	  "links": {
	    "self": "/company/12345678",
	    "filing_history": "/company/12345678/filing-history",
	    "officers": "/company/12345678/officers"
	  },
	  "can_file": true
	}`,

	"OC123456": `{
	  "has_been_liquidated": false,
	  "jurisdiction": "england-wales",
	  "date_of_creation": "2014-04-01",
	  "company_number": "OC123456",
	  "type": "llp",
	  "company_name": "TESTING LLP",
	  "status": "active",
	  "registered_office_address": {
	    "address_line_2": "2 Test Gardens",
	    "address_line_1": "The Test building",
	    "region": "Surrey",
	    "postal_code": "TE3 4ST",
	    "locality": "Testington"
	  },
	  "undeliverable_registered_office_address": false,
	  "accounts": {
	    "accounting_reference_date": {
	      "month": "12",
	      "day": "31"
	    },
	    "overdue": false,
	    "last_accounts": {
	      "type": "total-exemption-small",
	      "made_up_to": "2015-12-31"
	    },
	    "next_due": "2017-09-30",
	    "next_made_up_to": "2016-12-31"
	  },
	  "annual_return": {
	    "overdue": false,
	    "last_made_up_to": "2016-04-04"
	  },
	  "etag": "5c8350cc1146a247d273da25215fc6854f86f212",
	  "company_status": "active",
	  "has_insolvency_history": false,
	  "has_charges": false,
	  "confirmation_statement": {
	    "overdue": false,
	    "next_due": "2017-04-18",
	    "next_made_up_to": "2017-04-04"
	  },
	  "links": {
	    "self": "/company/OC123456",
	    "filing_history": "/company/OC123456/filing-history",
	    "officers": "/company/OC123456/officers"
	  },
	  "can_file": false
	}`,

	"FC123456": `{
	  "has_been_liquidated": false,
	  "undeliverable_registered_office_address": false,
	  "type": "oversea-company",
	  "registered_office_address": {
	    "address_line_1": "Branch Registration",
	    "country": "Netherlands",
	    "address_line_2": "Refer To Parent Registry"
	  },
	  "date_of_creation": "1981-05-01",
	  "company_name": "TEST FC NV",
	  "company_number": "FC123456",
	  "jurisdiction": "united-kingdom",
	  "accounts": {
	    "next_made_up_to": "2016-12-31",
	    "last_accounts": {
	      "type": "full",
	      "made_up_to": "2015-12-31"
	    },
	    "overdue": false
	  },
	  "etag": "9681f0e5233256b3b6ca8170f444f044f1a22d13",
	  "has_insolvency_history": false,
	  "company_status": "active",
	  "has_charges": true,
	  "previous_company_names": [
	    {
	      "name": "DUTCH TEST COMPANY N.V.",
	      "effective_from": "1979-08-01",
	      "ceased_on": "1990-04-04"
	    }
	  ],
	  "links": {
	    "self": "/company/FC123456",
	    "filing_history": "/company/FC123456/filing-history",
	    "officers": "/company/FC123456/officers",
	    "charges": "/company/FC123456/charges",
	    "uk_establishments": "/company/FC123456/uk-establishments"
	  },
	  "foreign_company_details": {
	    "accounting_requirement": {
	      "terms_of_account_publication": "accounting-publication-date-does-not-need-to-be-supplied-by-company",
	      "foreign_account_type": "accounting-requirements-of-originating-country-apply"
	    },
	    "legal_form": "Public Limited Company",
	    "is_a_credit_financial_institution": false,
	    "registration_number": "31431",
	    "originating_registry": {
	      "name": "Commercial Register Amsterdam",
	      "country": "NETHERLANDS"
	    }
	  },
	  "can_file": false
	}`,
}

var Officers = map[string]string{
	"12345678": `{
		"etag": "b258bb7bbe145d43fd6a9c1800dffa32ba24a001",
		"links": {
			"self": "/company/12345678/officers"
		},
		"kind": "officer-list",
		"active_count": 3,
		"resigned_count": 0,
		"start_index": 0,
		"inactive_count": 0,
		"total_results": 3,
		"items_per_page": 35,
		"items": [{
			"appointed_on": "2016-02-11",
			"name": "PERSON, Test",
			"links": {
				"officer": {
					"appointments": "/officers/avms6b9Jv3F35fRQx-PFUiabl00/appointments"
				}
			},
			"officer_role": "secretary",
			"address": {
				"country": "United Kingdom",
				"address_line_1": "15 Test circus",
				"locality": "London",
				"premises": "The Test Tower",
				"postal_code": "TT1 3ST",
				"region": "Greater-London"
			}
		}, {
			"country_of_residence": "Bulgaria",
			"appointed_on": "2016-02-11",
			"links": {
				"officer": {
					"appointments": "/officers/m79T6LcXbM479wP98xFL2cvaV00/appointments"
				}
			},
			"name": "TEST, Person",
			"address": {
				"country": "United Kingdom",
				"address_line_1": "15 Test circus",
				"locality": "London",
				"premises": "The Test Tower",
				"postal_code": "TT1 3ST",
				"region": "Greater-London"
			},
			"date_of_birth": {
				"month": 1,
				"year": 1960
			},
			"nationality": "Dutch",
			"officer_role": "director",
			"occupation": "Director"
		}, {
			"appointed_on": "2016-02-11",
			"officer_role": "corporate-director",
			"links": {
				"officer": {
					"appointments": "/officers/vWMfukdhT9ziMM8P6aFLMwX_s00/appointments"
				}
			},
			"identification": {
				"identification_type": "eea",
				"registration_number": "1234",
				"place_registered": "LVA"
			},
			"name": "SIA TEST COMPANY",
			"address": {
				"premises": "1-2",
				"postal_code": "1000",
				"country": "Latvia",
				"address_line_1": "Test Ploshtad",
				"locality": "Riga"
			}
		}]
	}`,

	"OC123456": `{
	  "resigned_count": 0,
	  "total_results": 2,
	  "items": [
	    {
	      "name": "ANOTHER, Person",
	      "appointed_on": "2014-04-01",
	      "officer_role": "llp-designated-member",
	      "links": {
		"officer": {
		  "appointments": "/officers/j1BLaZi3gDiD1Tl2X-K3HIzGN00/appointments"
		}
	      },
	      "address": {
		"premises": "12",
		"address_line_1": "Avenue Test",
		"country": "Bulgaria",
		"postal_code": "1000",
		"locality": "Sofia"
	      },
	      "date_of_birth": {
		"month": 2,
		"year": 1960
	      },
	      "country_of_residence": "France"
	    },
	    {
	      "officer_role": "llp-designated-member",
	      "name": "TEST, Person",
	      "appointed_on": "2014-04-01",
	      "address": {
		"country": "Netherlands",
		"postal_code": "1234 TS",
		"address_line_1": "Teststraat",
		"premises": "1",
		"locality": "Rotterdam"
	      },
	      "links": {
		"officer": {
		  "appointments": "/officers/m79T6LcXbM479wP98xFL2cvaV00/appointments"
		}
	      },
	      "date_of_birth": {
		"month": 1,
		"year": 1970
	      },
	      "country_of_residence": "Netherlands"
	    }
	  ],
	  "links": {
	    "self": "/company/OC123456/officers"
	  },
	  "start_index": 0,
	  "items_per_page": 35,
	  "active_count": 2,
	  "etag": "bea16b744fb83eaad3f5d56cc0d447fb75edf002",
	  "inactive_count": 0,
	  "kind": "officer-list"
	}`,

	"FC123456": `{
	  "active_count": 7,
	  "inactive_count": 0,
	  "etag": "c9bfc69034c7e4ec0ce6c19a85dd125124c69d03",
	  "items": [
	    {
	      "address": {
		"country": "Netherlands",
		"locality": "Amsterdam",
		"region": "1100 TS",
		"premises": "Testplein 123"
	      },
	      "country_of_residence": "Amsterdam",
	      "name": "TESTING, Person",
	      "links": {
		"officer": {
		  "appointments": "/officers/fHZJQUiIR5B_ALAv9j1nOaR6y00/appointments"
		}
	      },
	      "date_of_birth": {
		"year": 1953,
		"month": 1
	      },
	      "nationality": "Canadian",
	      "officer_role": "director",
	      "occupation": "Company Director",
	      "appointed_on": "2014-10-01"
	    },
	    {
	      "address": {
		"premises": "Testplein 123",
		"country": "Netherlands",
		"locality": "1100 TS Amsterdam"
	      },
	      "country_of_residence": "The Netherlands",
	      "nationality": "American",
	      "date_of_birth": {
		"year": 1948,
		"month": 9
	      },
	      "links": {
		"officer": {
		  "appointments": "/officers/k1tnTzFfKO6r9wGDaKkC1Uwyc00/appointments"
		}
	      },
	      "name": "TESTER, Tinus",
	      "occupation": "Company Director",
	      "appointed_on": "2011-01-01",
	      "officer_role": "director"
	    },
	    {
	      "officer_role": "director",
	      "appointed_on": "2009-04-12",
	      "occupation": "Company Director",
	      "date_of_birth": {
		"year": 1950,
		"month": 2
	      },
	      "nationality": "Irish",
	      "name": "O'SELSKI, Marc",
	      "links": {
		"officer": {
		  "appointments": "/officers/yTKlCi_i3rOXZo6EyU8XIozZK00/appointments"
		}
	      },
	      "country_of_residence": "Netherlands",
	      "address": {
		"locality": "1000 TS Amsterdam",
		"country": "Netherlands",
		"address_line_1": "Testplein 123"
	      }
	    },
	    {
	      "appointed_on": "2000-03-31",
	      "officer_role": "secretary",
	      "links": {
		"officer": {
		  "appointments": "/officers/AvJBZx9hkTzdAd6C2rT9-DRU700/appointments"
		}
	      },
	      "name": "TEST, Secretary",
	      "address": {
		"postal_code": "TS4R 3XT",
		"address_line_1": "12 Test Lane",
		"locality": "London",
		"premises": "Testing House"
	      }
	    },
	    {
	      "nationality": "Belgian",
	      "date_of_birth": {
		"year": 1962,
		"month": 5
	      },
	      "name": "DE LA TEST, Person",
	      "links": {
		"officer": {
		  "appointments": "/officers/dXJ_BszeBJFx-1snbFmpQPgMB00/appointments"
		}
	      },
	      "officer_role": "director",
	      "appointed_on": "2006-02-07",
	      "occupation": "Company Director",
	      "address": {
		"country": "Netherlands",
		"address_line_1": "Testplein 123",
		"locality": "1100 TS Amsterdam"
	      },
	      "resigned_on": "2011-10-01",
	      "country_of_residence": "Belgium"
	    }
	  ],
	  "start_index": 0,
	  "total_results": 35,
	  "links": {
	    "self": "/company/FC123456/officers"
	  },
	  "resigned_count": 28,
	  "items_per_page": 35,
	  "kind": "officer-list"
	}`,
}

var Insolvencies = map[string]string{
	"12345678": `{
	  "cases": [
	    {
	      "type": "compulsory-liquidation",
	      "practitioners": [
		{
		  "name": "Test Practitioner",
		  "address": {
		    "address_line_1": "1 Test Court Oxford Street",
		    "postal_code": "TS12 3ST",
		    "region": "Greater London",
		    "locality": "London"
		  },
		  "role": "practitioner"
		},
		{
		  "name": "Another Practitioner",
		  "role": "practitioner",
		  "address": {
		    "address_line_2": "Test Grove",
		    "locality": "Surrey",
		    "postal_code": "TE3 5ST",
		    "address_line_1": "1st Floor Test House"
		  }
		}
	      ],
	      "dates": [
		{
		  "type": "wound-up-on",
		  "date": "2014-04-28"
		},
		{
		  "date": "2014-02-28",
		  "type": "petitioned-on"
		}
	      ],
	      "number": "1"
	    }
	  ],
	  "etag": "1234",
	  "status": [
	    "liquidation"
	  ]
	}`,

	"OC123456": `{
	  "cases": [
	    {
	      "type": "compulsory-liquidation",
	      "practitioners": [
		{
		  "name": "Test Practitioner",
		  "address": {
		    "address_line_1": "1 Test Court Oxford Street",
		    "postal_code": "TS12 3ST",
		    "region": "Greater London",
		    "locality": "London"
		  },
		  "role": "practitioner"
		},
		{
		  "name": "Another Practitioner",
		  "role": "practitioner",
		  "address": {
		    "address_line_2": "Test Grove",
		    "locality": "Surrey",
		    "postal_code": "TE3 5ST",
		    "address_line_1": "1st Floor Test House"
		  }
		}
	      ],
	      "dates": [
		{
		  "type": "wound-up-on",
		  "date": "2014-04-28"
		},
		{
		  "date": "2014-02-28",
		  "type": "petitioned-on"
		}
	      ],
	      "number": "1"
	    }
	  ],
	  "etag": "1234",
	  "status": [
	    "liquidation"
	  ]
	}`,

	"FC123456": `{
	  "cases": [
	    {
	      "type": "compulsory-liquidation",
	      "practitioners": [
		{
		  "name": "Test Practitioner",
		  "address": {
		    "address_line_1": "1 Test Court Oxford Street",
		    "postal_code": "TS12 3ST",
		    "region": "Greater London",
		    "locality": "London"
		  },
		  "role": "practitioner"
		},
		{
		  "name": "Another Practitioner",
		  "role": "practitioner",
		  "address": {
		    "address_line_2": "Test Grove",
		    "locality": "Surrey",
		    "postal_code": "TE3 5ST",
		    "address_line_1": "1st Floor Test House"
		  }
		}
	      ],
	      "dates": [
		{
		  "type": "wound-up-on",
		  "date": "2014-04-28"
		},
		{
		  "date": "2014-02-28",
		  "type": "petitioned-on"
		}
	      ],
	      "number": "1"
	    }
	  ],
	  "etag": "1234",
	  "status": [
	    "liquidation"
	  ]
	}`,
}

var Charges = map[string]string{
	"12345678": `{
	  "items": [
	    {
	      "created_on": "1991-04-30",
	      "transactions": [
		{
		  "delivered_on": "1991-05-08",
		  "filing_type": "create-charge"
		},
		{
		  "delivered_on": "2016-02-05",
		  "filing_type": "charge-satisfaction",
		  "links": {
		    "filing": "/company/12345678/filing-history/MzE0MTMyNzA0NmFkaXF6a2N4"
		  }
		}
	      ],
	      "etag": "51943ec0077900cfaa9bb159e73cae043e937114",
	      "classification": {
		"description": "Mortgage debenture",
		"type": "charge-description"
	      },
	      "particulars": {
		"description": "A specific equitable charge over all freehold and leasehold properties and/or the proceeds of sale thereof fixed and floating charges over undertaking and all property and assets present and future including goodwill bookdebts and the benefits of any licences.",
		"type": "short-particulars"
	      },
	      "links": {
		"self": "/company/12345678/charges/CkDWpZcM4AKGK7oqRpmdEtliG_I"
	      },
	      "satisfied_on": "2016-02-05",
	      "delivered_on": "1991-05-08",
	      "persons_entitled": [
		{
		  "name": "Test Bank PLC"
		}
	      ],
	      "secured_details": {
		"description": "All monies due or to become due from the company to the chargee on any account whatsoever",
		"type": "amount-secured"
	      },
	      "charge_number": 5,
	      "status": "fully-satisfied"
	    },
	    {
	      "etag": "0fae715a940e532378831206561b84e504f5de29",
	      "classification": {
		"description": "Mortgage",
		"type": "charge-description"
	      },
	      "secured_details": {
		"type": "amount-secured",
		"description": "All monies due or to become due from the company to the chargee on any account whatsoever"
	      },
	      "status": "fully-satisfied",
	      "delivered_on": "1979-03-06",
	      "links": {
		"self": "/company/12345678/charges/YWHV8QNkOm7XxV7z_f2aKJzKPOE"
	      },
	      "particulars": {
		"type": "short-particulars",
		"description": "Floating charge over all moveable test machinery implements utensils furniture and equipment."
	      },
	      "satisfied_on": "2015-12-15",
	      "transactions": [
		{
		  "delivered_on": "1979-03-06",
		  "filing_type": "create-charge"
		},
		{
		  "links": {
		    "filing": "/company/12345678/filing-history/MzEzNzU5NDgzNWFkaXF6a2N4"
		  },
		  "filing_type": "charge-satisfaction",
		  "delivered_on": "2015-12-15"
		}
	      ],
	      "created_on": "1979-02-15",
	      "charge_number": 4,
	      "persons_entitled": [
		{
		  "name": "Test Bank PLC"
		}
	      ]
	    },
	    {
	      "persons_entitled": [
		{
		  "name": "Test company LTD"
		}
	      ],
	      "charge_number": 3,
	      "particulars": {
		"type": "short-particulars",
		"description": "Test equipment"
	      },
	      "transactions": [
		{
		  "delivered_on": "1974-05-07",
		  "filing_type": "create-charge"
		}
	      ],
	      "links": {
		"self": "/company/12345678/charges/1bbP7hnOm03q6wz8OOctdy2ZaBc"
	      },
	      "created_on": "1974-04-22",
	      "secured_details": {
		"type": "amount-secured",
		"description": "£8,500 & all other monies due or to become due from the company to the chargee on any account whatsoever."
	      },
	      "status": "outstanding",
	      "classification": {
		"type": "charge-description",
		"description": "Further charge"
	      },
	      "etag": "175a21610fc96ff5d7f62e6db6049bda682996c4",
	      "delivered_on": "1974-05-07"
	    },
	    {
	      "status": "outstanding",
	      "classification": {
		"description": "Legal mortgage",
		"type": "charge-description"
	      },
	      "particulars": {
		"description": "Floating charge over all movable test machinery implements utensils furniture and equipment.",
		"type": "short-particulars"
	      },
	      "links": {
		"self": "/company/12345678/charges/B9kYrc0F7P6C0n5nELcgXBaLqoY"
	      },
	      "transactions": [
		{
		  "delivered_on": "1973-07-06",
		  "filing_type": "create-charge"
		}
	      ],
	      "created_on": "1973-07-02",
	      "secured_details": {
		"type": "amount-secured",
		"description": "All monies due or to become due from the company to the chargee on any account whatsoever"
	      },
	      "charge_number": 2,
	      "persons_entitled": [
		{
		  "name": "Test Bank PLC"
		}
	      ],
	      "delivered_on": "1973-07-06",
	      "etag": "1ac1a6ee4ca12e54dc5a38d97f3af3fa29d99f5d"
	    },
	    {
	      "etag": "68a69353ae4157992017fd7bd0def92c6171524b",
	      "delivered_on": "1969-07-17",
	      "links": {
		"self": "/company/12345678/charges/tqM_Pztu3Edr88WdEiSjw_e7Ub0"
	      },
	      "transactions": [
		{
		  "filing_type": "create-charge",
		  "delivered_on": "1969-07-17"
		}
	      ],
	      "created_on": "1969-07-15",
	      "secured_details": {
		"type": "amount-secured",
		"description": "For securing £4,000 and all other monies due or to become due from thecompany to the chargee on any account whatsoever and any other monies that may become owing by the company to the chargee on account current or stated for goodssupplied or otherwise."
	      },
	      "particulars": {
		"type": "short-particulars",
		"description": "The beneficial interest and goodwill of the business carried on thereat (fixed & floating charge see doc 11. for further details). Undertaking and all property and assets present and future including uncalled capital."
	      },
	      "charge_number": 1,
	      "persons_entitled": [
		{
		  "name": "Another Test LTD."
		}
	      ],
	      "classification": {
		"description": "Mortgage",
		"type": "charge-description"
	      },
	      "status": "outstanding"
	    }
	  ],
	  "part_satisfied_count": 0,
	  "unfiltered_count": 5,
	  "satisfied_count": 2,
	  "total_count": 5
	}`,

	"OC123456": `{
	  "items": [
	    {
	      "created_on": "1991-04-30",
	      "transactions": [
		{
		  "delivered_on": "1991-05-08",
		  "filing_type": "create-charge"
		},
		{
		  "delivered_on": "2016-02-05",
		  "filing_type": "charge-satisfaction",
		  "links": {
		    "filing": "/company/12345678/filing-history/MzE0MTMyNzA0NmFkaXF6a2N4"
		  }
		}
	      ],
	      "etag": "51943ec0077900cfaa9bb159e73cae043e937114",
	      "classification": {
		"description": "Mortgage debenture",
		"type": "charge-description"
	      },
	      "particulars": {
		"description": "A specific equitable charge over all freehold and leasehold properties and/or the proceeds of sale thereof fixed and floating charges over undertaking and all property and assets present and future including goodwill bookdebts and the benefits of any licences.",
		"type": "short-particulars"
	      },
	      "links": {
		"self": "/company/12345678/charges/CkDWpZcM4AKGK7oqRpmdEtliG_I"
	      },
	      "satisfied_on": "2016-02-05",
	      "delivered_on": "1991-05-08",
	      "persons_entitled": [
		{
		  "name": "Test Bank PLC"
		}
	      ],
	      "secured_details": {
		"description": "All monies due or to become due from the company to the chargee on any account whatsoever",
		"type": "amount-secured"
	      },
	      "charge_number": 5,
	      "status": "fully-satisfied"
	    },
	    {
	      "etag": "0fae715a940e532378831206561b84e504f5de29",
	      "classification": {
		"description": "Mortgage",
		"type": "charge-description"
	      },
	      "secured_details": {
		"type": "amount-secured",
		"description": "All monies due or to become due from the company to the chargee on any account whatsoever"
	      },
	      "status": "fully-satisfied",
	      "delivered_on": "1979-03-06",
	      "links": {
		"self": "/company/12345678/charges/YWHV8QNkOm7XxV7z_f2aKJzKPOE"
	      },
	      "particulars": {
		"type": "short-particulars",
		"description": "Floating charge over all moveable test machinery implements utensils furniture and equipment."
	      },
	      "satisfied_on": "2015-12-15",
	      "transactions": [
		{
		  "delivered_on": "1979-03-06",
		  "filing_type": "create-charge"
		},
		{
		  "links": {
		    "filing": "/company/12345678/filing-history/MzEzNzU5NDgzNWFkaXF6a2N4"
		  },
		  "filing_type": "charge-satisfaction",
		  "delivered_on": "2015-12-15"
		}
	      ],
	      "created_on": "1979-02-15",
	      "charge_number": 4,
	      "persons_entitled": [
		{
		  "name": "Test Bank PLC"
		}
	      ]
	    },
	    {
	      "persons_entitled": [
		{
		  "name": "Test company LTD"
		}
	      ],
	      "charge_number": 3,
	      "particulars": {
		"type": "short-particulars",
		"description": "Test equipment"
	      },
	      "transactions": [
		{
		  "delivered_on": "1974-05-07",
		  "filing_type": "create-charge"
		}
	      ],
	      "links": {
		"self": "/company/12345678/charges/1bbP7hnOm03q6wz8OOctdy2ZaBc"
	      },
	      "created_on": "1974-04-22",
	      "secured_details": {
		"type": "amount-secured",
		"description": "£8,500 & all other monies due or to become due from the company to the chargee on any account whatsoever."
	      },
	      "status": "outstanding",
	      "classification": {
		"type": "charge-description",
		"description": "Further charge"
	      },
	      "etag": "175a21610fc96ff5d7f62e6db6049bda682996c4",
	      "delivered_on": "1974-05-07"
	    },
	    {
	      "status": "outstanding",
	      "classification": {
		"description": "Legal mortgage",
		"type": "charge-description"
	      },
	      "particulars": {
		"description": "Floating charge over all movable test machinery implements utensils furniture and equipment.",
		"type": "short-particulars"
	      },
	      "links": {
		"self": "/company/12345678/charges/B9kYrc0F7P6C0n5nELcgXBaLqoY"
	      },
	      "transactions": [
		{
		  "delivered_on": "1973-07-06",
		  "filing_type": "create-charge"
		}
	      ],
	      "created_on": "1973-07-02",
	      "secured_details": {
		"type": "amount-secured",
		"description": "All monies due or to become due from the company to the chargee on any account whatsoever"
	      },
	      "charge_number": 2,
	      "persons_entitled": [
		{
		  "name": "Test Bank PLC"
		}
	      ],
	      "delivered_on": "1973-07-06",
	      "etag": "1ac1a6ee4ca12e54dc5a38d97f3af3fa29d99f5d"
	    },
	    {
	      "etag": "68a69353ae4157992017fd7bd0def92c6171524b",
	      "delivered_on": "1969-07-17",
	      "links": {
		"self": "/company/12345678/charges/tqM_Pztu3Edr88WdEiSjw_e7Ub0"
	      },
	      "transactions": [
		{
		  "filing_type": "create-charge",
		  "delivered_on": "1969-07-17"
		}
	      ],
	      "created_on": "1969-07-15",
	      "secured_details": {
		"type": "amount-secured",
		"description": "For securing £4,000 and all other monies due or to become due from thecompany to the chargee on any account whatsoever and any other monies that may become owing by the company to the chargee on account current or stated for goodssupplied or otherwise."
	      },
	      "particulars": {
		"type": "short-particulars",
		"description": "The beneficial interest and goodwill of the business carried on thereat (fixed & floating charge see doc 11. for further details). Undertaking and all property and assets present and future including uncalled capital."
	      },
	      "charge_number": 1,
	      "persons_entitled": [
		{
		  "name": "Another Test LTD."
		}
	      ],
	      "classification": {
		"description": "Mortgage",
		"type": "charge-description"
	      },
	      "status": "outstanding"
	    }
	  ],
	  "part_satisfied_count": 0,
	  "unfiltered_count": 5,
	  "satisfied_count": 2,
	  "total_count": 5
	}`,

	"FC123456": `{
	  "items": [
	    {
	      "created_on": "1991-04-30",
	      "transactions": [
		{
		  "delivered_on": "1991-05-08",
		  "filing_type": "create-charge"
		},
		{
		  "delivered_on": "2016-02-05",
		  "filing_type": "charge-satisfaction",
		  "links": {
		    "filing": "/company/12345678/filing-history/MzE0MTMyNzA0NmFkaXF6a2N4"
		  }
		}
	      ],
	      "etag": "51943ec0077900cfaa9bb159e73cae043e937114",
	      "classification": {
		"description": "Mortgage debenture",
		"type": "charge-description"
	      },
	      "particulars": {
		"description": "A specific equitable charge over all freehold and leasehold properties and/or the proceeds of sale thereof fixed and floating charges over undertaking and all property and assets present and future including goodwill bookdebts and the benefits of any licences.",
		"type": "short-particulars"
	      },
	      "links": {
		"self": "/company/12345678/charges/CkDWpZcM4AKGK7oqRpmdEtliG_I"
	      },
	      "satisfied_on": "2016-02-05",
	      "delivered_on": "1991-05-08",
	      "persons_entitled": [
		{
		  "name": "Test Bank PLC"
		}
	      ],
	      "secured_details": {
		"description": "All monies due or to become due from the company to the chargee on any account whatsoever",
		"type": "amount-secured"
	      },
	      "charge_number": 5,
	      "status": "fully-satisfied"
	    },
	    {
	      "etag": "0fae715a940e532378831206561b84e504f5de29",
	      "classification": {
		"description": "Mortgage",
		"type": "charge-description"
	      },
	      "secured_details": {
		"type": "amount-secured",
		"description": "All monies due or to become due from the company to the chargee on any account whatsoever"
	      },
	      "status": "fully-satisfied",
	      "delivered_on": "1979-03-06",
	      "links": {
		"self": "/company/12345678/charges/YWHV8QNkOm7XxV7z_f2aKJzKPOE"
	      },
	      "particulars": {
		"type": "short-particulars",
		"description": "Floating charge over all moveable test machinery implements utensils furniture and equipment."
	      },
	      "satisfied_on": "2015-12-15",
	      "transactions": [
		{
		  "delivered_on": "1979-03-06",
		  "filing_type": "create-charge"
		},
		{
		  "links": {
		    "filing": "/company/12345678/filing-history/MzEzNzU5NDgzNWFkaXF6a2N4"
		  },
		  "filing_type": "charge-satisfaction",
		  "delivered_on": "2015-12-15"
		}
	      ],
	      "created_on": "1979-02-15",
	      "charge_number": 4,
	      "persons_entitled": [
		{
		  "name": "Test Bank PLC"
		}
	      ]
	    },
	    {
	      "persons_entitled": [
		{
		  "name": "Test company LTD"
		}
	      ],
	      "charge_number": 3,
	      "particulars": {
		"type": "short-particulars",
		"description": "Test equipment"
	      },
	      "transactions": [
		{
		  "delivered_on": "1974-05-07",
		  "filing_type": "create-charge"
		}
	      ],
	      "links": {
		"self": "/company/12345678/charges/1bbP7hnOm03q6wz8OOctdy2ZaBc"
	      },
	      "created_on": "1974-04-22",
	      "secured_details": {
		"type": "amount-secured",
		"description": "£8,500 & all other monies due or to become due from the company to the chargee on any account whatsoever."
	      },
	      "status": "outstanding",
	      "classification": {
		"type": "charge-description",
		"description": "Further charge"
	      },
	      "etag": "175a21610fc96ff5d7f62e6db6049bda682996c4",
	      "delivered_on": "1974-05-07"
	    },
	    {
	      "status": "outstanding",
	      "classification": {
		"description": "Legal mortgage",
		"type": "charge-description"
	      },
	      "particulars": {
		"description": "Floating charge over all movable test machinery implements utensils furniture and equipment.",
		"type": "short-particulars"
	      },
	      "links": {
		"self": "/company/12345678/charges/B9kYrc0F7P6C0n5nELcgXBaLqoY"
	      },
	      "transactions": [
		{
		  "delivered_on": "1973-07-06",
		  "filing_type": "create-charge"
		}
	      ],
	      "created_on": "1973-07-02",
	      "secured_details": {
		"type": "amount-secured",
		"description": "All monies due or to become due from the company to the chargee on any account whatsoever"
	      },
	      "charge_number": 2,
	      "persons_entitled": [
		{
		  "name": "Test Bank PLC"
		}
	      ],
	      "delivered_on": "1973-07-06",
	      "etag": "1ac1a6ee4ca12e54dc5a38d97f3af3fa29d99f5d"
	    },
	    {
	      "etag": "68a69353ae4157992017fd7bd0def92c6171524b",
	      "delivered_on": "1969-07-17",
	      "links": {
		"self": "/company/12345678/charges/tqM_Pztu3Edr88WdEiSjw_e7Ub0"
	      },
	      "transactions": [
		{
		  "filing_type": "create-charge",
		  "delivered_on": "1969-07-17"
		}
	      ],
	      "created_on": "1969-07-15",
	      "secured_details": {
		"type": "amount-secured",
		"description": "For securing £4,000 and all other monies due or to become due from thecompany to the chargee on any account whatsoever and any other monies that may become owing by the company to the chargee on account current or stated for goodssupplied or otherwise."
	      },
	      "particulars": {
		"type": "short-particulars",
		"description": "The beneficial interest and goodwill of the business carried on thereat (fixed & floating charge see doc 11. for further details). Undertaking and all property and assets present and future including uncalled capital."
	      },
	      "charge_number": 1,
	      "persons_entitled": [
		{
		  "name": "Another Test LTD."
		}
	      ],
	      "classification": {
		"description": "Mortgage",
		"type": "charge-description"
	      },
	      "status": "outstanding"
	    }
	  ],
	  "part_satisfied_count": 0,
	  "unfiltered_count": 5,
	  "satisfied_count": 2,
	  "total_count": 5
	}`,
}

var Filings = map[string]string{
	"12345678": `{
	  "items": [
	    {
	      "date": "2016-02-11",
	      "links": {
		"self": "/company/12345678/filing-history/MzE0MTcyMzAyMmFkaXF6a200",
		"document_metadata": "https://document-api.companieshouse.gov.uk/document/CPuWIu7JpIOSz5A040Tg8lX48fRMSsSzO9OamjxaP00"
	      },
	      "type": "AA01",
	      "description": "change-account-reference-date-company-current-shortened",
	      "category": "accounts",
	      "description_values": {
		"made_up_date": "2017-02-28",
		"new_date": "2016-12-31"
	      },
	      "action_date": "2016-12-31",
	      "pages": 1,
	      "barcode": "12312345",
	      "transaction_id": "MzE0MTcyMzAyMmFkaXF6a200"
	    },
	    {
	      "associated_filings": [
		{
		  "action_date": 1455148800000,
		  "category": "capital",
		  "date": "2016-02-11",
		  "description": "statement-of-capital",
		  "description_values": {
		    "capital": [
		      {
			"currency": "EUR",
			"figure": "1"
		      }
		    ],
		    "date": "2016-02-11"
		  },
		  "original_description": "11/02/16 Statement of Capital;EUR 1",
		  "type": "SH01"
		}
	      ],
	      "links": {
		"self": "/company/12345678/filing-history/MzE0MTcxMTU5OWFkaXF6a200",
		"document_metadata": "https://document-api.companieshouse.gov.uk/document/2VQ2bkhYRwzK3RXigUcwQZHdzeZSHJXoyLXVcFOLj00"
	      },
	      "description": "incorporation-company",
	      "type": "NEWINC",
	      "category": "incorporation",
	      "date": "2016-02-01",
	      "pages": 28,
	      "barcode": "123123457",
	      "transaction_id": "MzE0MTcxMTU5OWFkaXF6a200"
	    }
	  ],
	  "total_count": 2,
	  "items_per_page": 25,
	  "start_index": 0,
	  "filing_history_status": "filing-history-available"
	}`,

	"OC123456": `{
	  "start_index": 0,
	  "filing_history_status": "filing-history-available",
	  "total_count": 11,
	  "items_per_page": 25,
	  "items": [
	    {
	      "date": "2016-09-22",
	      "links": {
		"self": "/company/OC123456/filing-history/MzE1Nzc0OTk1N2FkaXF6a200",
		"document_metadata": "https://document-api.companieshouse.gov.uk/document/XDSn_uEk_qOjKbEpGOxNgRyEL3-bVMAStk8tLj1gh00"
	      },
	      "action_date": "2015-12-31",
	      "category": "accounts",
	      "description": "accounts-with-accounts-type-total-exemption-small",
	      "description_values": {
		"made_up_date": "2015-12-31"
	      },
	      "paper_filed": true,
	      "type": "AA",
	      "pages": 1,
	      "barcode": "BBFJYSEE",
	      "transaction_id": "MzE1Nzc0OTk1N2FkaXF6a200"
	    },
	    {
	      "links": {
		"self": "/company/OC123456/filing-history/MzE1NjE2MzQ3NWFkaXF6a200",
		"document_metadata": "https://document-api.companieshouse.gov.uk/document/cmJcP_nDl_pa_IQfBLXgPJpAmUxEFp1GJbHMpI4Ej00"
	      },
	      "type": "DISS40",
	      "date": "2016-08-31",
	      "category": "gazette",
	      "description": "gazette-filings-brought-up-to-date",
	      "paper_filed": true,
	      "pages": 1,
	      "barcode": null,
	      "transaction_id": "MzE1NjE2MzQ3NWFkaXF6a200"
	    },
	    {
	      "action_date": "2016-04-04",
	      "date": "2016-08-30",
	      "description_values": {
		"made_up_date": "2016-04-04"
	      },
	      "category": "annual-return",
	      "description": "annual-return-limited-liability-partnership-with-made-up-date",
	      "links": {
		"self": "/company/OC123456/filing-history/MzE1NjE2MzQ0MmFkaXF6a200",
		"document_metadata": "https://document-api.companieshouse.gov.uk/document/IZs1b8vGcstEWOa3GzShNtJIsF0ob2vVYztCoowGd00"
	      },
	      "type": "LLAR01",
	      "pages": 3,
	      "barcode": "YYEHOGEE",
	      "transaction_id": "MzE1NjE2MzQ0MmFkaXF6a200"
	    },
	    {
	      "type": "GAZ1",
	      "links": {
		"self": "/company/OC123456/filing-history/MzE1MTI4NTI2MWFkaXF6a200",
		"document_metadata": "https://document-api.companieshouse.gov.uk/document/2I5d9rCC6bvyeuQBFI1P4Y8MxV_DmU7x1kDuw_j1b00"
	      },
	      "category": "gazette",
	      "date": "2016-07-05",
	      "description": "gazette-notice-compulsory",
	      "paper_filed": true,
	      "pages": 1,
	      "barcode": null,
	      "transaction_id": "MzE1MTI4NTI2MWFkaXF6a200"
	    },
	    {
	      "links": {
		"self": "/company/OC123456/filing-history/MzEzMDk0Mzc4OGFkaXF6a200",
		"document_metadata": "https://document-api.companieshouse.gov.uk/document/IR2r-A5RJgWGKx7RbD3exbLXXOgTjMFGGpJi63u6r00"
	      },
	      "paper_filed": true,
	      "date": "2015-09-18",
	      "type": "AA",
	      "action_date": "2014-12-31",
	      "description_values": {
		"made_up_date": "2014-12-31"
	      },
	      "description": "accounts-with-accounts-type-total-exemption-small",
	      "category": "accounts",
	      "pages": 1,
	      "barcode": "A4FIAI5C",
	      "transaction_id": "MzEzMDk0Mzc4OGFkaXF6a200"
	    },
	    {
	      "description": "gazette-filings-brought-up-to-date",
	      "category": "gazette",
	      "date": "2015-08-11",
	      "type": "DISS40",
	      "links": {
		"self": "/company/OC123456/filing-history/MzEyODQzOTg0MmFkaXF6a200",
		"document_metadata": "https://document-api.companieshouse.gov.uk/document/yzfK_oFekTlFDoftuci26jI_sLNPcBdJwv2ochRYr00"
	      },
	      "paper_filed": true,
	      "pages": 1,
	      "barcode": null,
	      "transaction_id": "MzEyODQzOTg0MmFkaXF6a200"
	    },
	    {
	      "description_values": {
		"made_up_date": "2015-04-04"
	      },
	      "date": "2015-08-05",
	      "action_date": "2015-04-04",
	      "description": "annual-return-limited-liability-partnership-with-made-up-date",
	      "category": "annual-return",
	      "type": "LLAR01",
	      "links": {
		"self": "/company/OC123456/filing-history/MzEyODM4MDA1NmFkaXF6a200",
		"document_metadata": "https://document-api.companieshouse.gov.uk/document/Of5Pdh2cFO6Prj5fksi4lTIIhv0rPoEcq-tYu8OD-00"
	      },
	      "pages": 3,
	      "barcode": "X4D22NIQ",
	      "transaction_id": "MzEyODM4MDA1NmFkaXF6a200"
	    },
	    {
	      "description": "change-registered-office-address-limited-liability-partnership-with-date-old-address-new-address",
	      "category": "address",
	      "type": "LLAD01",
	      "links": {
		"self": "/company/OC123456/filing-history/MzEyODM3OTg5N2FkaXF6a200",
		"document_metadata": "https://document-api.companieshouse.gov.uk/document/JGs1zqOKPoHQ7px3d6f8ujqmphbXgq3Ce3DZUeOZi00"
	      },
	      "action_date": "2015-08-04",
	      "date": "2015-08-04",
	      "description_values": {
		"new_address": "12 Test Road Testington Surrey TT5 3XT United Kingdom",
		"change_date": "2015-08-04",
		"old_address": "15 Test Lane London W1 12T United Kingdom"
	      },
	      "pages": 1,
	      "barcode": "12312388",
	      "transaction_id": "MzEyODM3OTg5N2FkaXF6a200"
	    },
	    {
	      "paper_filed": true,
	      "category": "gazette",
	      "description": "gazette-notice-compulsory",
	      "links": {
		"self": "/company/OC123456/filing-history/MzEyNzg2NjEzMmFkaXF6a200",
		"document_metadata": "https://document-api.companieshouse.gov.uk/document/tEuoTM3sDUolvVTChLF-KN4NfMbuwbI-Idp6qNw_p00"
	      },
	      "type": "GAZ1",
	      "date": "2015-08-04",
	      "pages": 1,
	      "barcode": null,
	      "transaction_id": "MzEyNzg2NjEzMmFkaXF6a200"
	    },
	    {
	      "description_values": {
		"made_up_date": "2015-04-30",
		"new_date": "2014-12-31"
	      },
	      "type": "LLAA01",
	      "action_date": "2014-12-31",
	      "category": "accounts",
	      "date": "2014-04-04",
	      "description": "change-account-reference-date-limited-liability-partnership-current-shortened",
	      "links": {
		"document_metadata": "https://document-api.companieshouse.gov.uk/document/vOusUXBePaOW-J-uuR0gHB7c5F3cw5E8-A2IvTfEc00",
		"self": "/company/OC123456/filing-history/MzA5NzY1ODgxNmFkaXF6a200"
	      },
	      "pages": 1,
	      "barcode": "X3546HSW",
	      "transaction_id": "MzA5NzY1ODgxNmFkaXF6a2N4"
	    },
	    {
	      "description": "incorporation-limited-liability-partnership",
	      "date": "2014-04-04",
	      "category": "incorporation",
	      "type": "LLIN01",
	      "links": {
		"document_metadata": "https://document-api.companieshouse.gov.uk/document/Msl5LzDF6B-OZe3AUtU3r9DVfpBNgyujFUf4-x_E800",
		"self": "/company/OC123456/filing-history/MzA5NzY0MDQ5MGFkaXF6a200"
	      },
	      "pages": 5,
	      "paper_filed": true,
	      "barcode": "X3544Y4R",
	      "transaction_id": "MzA5NzY0MDQ5MGFkaXF6a200"
	    }
	  ]
	}`,

	"FC123456": `{
		"items": [{
			"date": "2016-02-11",
			"links": {
				"self": "/company/FC123456/filing-history/MzE0MTcyMzAyMmFkaXF6a210",
				"document_metadata": "https://document-api.companieshouse.gov.uk/document/CPuWIu7JpIOSz5A040Tg8lX48fRMSsSzO9OamjxaP10"
			},
			"type": "AA01",
			"description": "change-account-reference-date-company-current-shortened",
			"category": "accounts",
			"description_values": {
				"made_up_date": "2017-02-28",
				"new_date": "2016-12-31"
			},
			"action_date": "2016-12-31",
			"pages": 1,
			"barcode": "12312345",
			"transaction_id": "MzE0MTcyMzAyMmFkaXF6a210"
		}, {
			"associated_filings": [{
				"action_date": 1455148810100,
				"category": "capital",
				"date": "2016-02-11",
				"description": "statement-of-capital",
				"description_values": {
					"capital": [{
						"currency": "EUR",
						"figure": "1"
					}],
					"date": "2016-02-11"
				},
				"original_description": "11/02/16 Statement of Capital;EUR 1",
				"type": "SH01"
			}],
			"links": {
				"self": "/company/FC123456/filing-history/MzE0MTcxMTU5OWFkaXF6a210",
				"document_metadata": "https://document-api.companieshouse.gov.uk/document/2VQ2bkhYRwzK3RXigUcwQZHdzeZSHJXoyLXVcFOLj10"
			},
			"description": "incorporation-company",
			"type": "NEWINC",
			"category": "incorporation",
			"date": "2016-02-01",
			"pages": 28,
			"barcode": "123123457",
			"transaction_id": "MzE0MTcxMTU5OWFkaXF6a210"
		}],
		"total_count": 2,
		"items_per_page": 25,
		"start_index": 0,
		"filing_history_status": "filing-history-available"
	}`,
}

//Todo: testdata.go:21:5: exported var Companies should have comment or be unexported
//Todo: testdata.go:171:5: exported var Officers should have comment or be unexported
//Todo: testdata.go:432:5: exported var Insolvencies should have comment or be unexported
//Todo: testdata.go:569:5: exported var Charges should have comment or be unexported
//Todo: testdata.go:1141:5: exported var Filings should have comment or be unexported
