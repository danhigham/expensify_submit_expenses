package main

import (
	"log"
	"time"
	// "fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigyaml"
)

type Config struct {
	API struct {
		Endpoint string `yaml:"endpoint"`
		Key      string `yaml:"key"`
		Secret   string `yaml:"secret"`
	} `yaml:"api"`
	Email  string `yaml:"email"`
	Report struct {
		Name     string `yaml:"name"`
		PolicyID string `yaml:"policy_id"`
	} `yaml:"report"`
	Expenses []ReportExpense `yaml:"expenses"`
}

type ReportCredentials struct {
	PartnerUserID     string `json:"partnerUserID"`
	PartnerUserSecret string `json:"partnerUserSecret"`
}

type ReportExpense struct {
	Date     string `json:"date", yaml:"date"`
	Currency string `json:"currency", yaml:"currency"`
	Merchant string `json:"merchant", yaml:"merchant"`
	Amount   int    `json:"amount", yaml:"amount"`
	Category string `json:"category", yaml:"category"`
	Comment string `json:"comment", yaml:"comment"`
	Department string `json:"tag", yaml:"department"`
}

type ReportInfo struct {
	Title  string `json:"title"`
}

type ReportInputSettings struct {
	Type     string `json:"type"`
	PolicyID string `json:"policyID"`
	Report  ReportInfo `json:"report"`
	EmployeeEmail string `json:"employeeEmail"`
	Expenses      []ReportExpense `json:"expenses"`
}

type ReportRequest struct {
	Type        string `json:"type"`
	Credentials ReportCredentials `json:"credentials"`
	InputSettings ReportInputSettings `json:"inputSettings"`
}

func postToExpensify(jsonJobDescription, endpoint string) error {
	resp, err := http.PostForm(endpoint, url.Values{"requestJobDescription": {jsonJobDescription}})
	// read response body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(resp.StatusCode)
	log.Println(string(body))

	return err
}

func main() {
	var cfg Config
	
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		SkipDefaults: true,
		SkipFlags:    true,
		EnvPrefix:       "EXPENSE",
		Files:           []string{"config.yaml"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": aconfigyaml.New(),
		},
	})

	if err := loader.Load(); err != nil {
		panic(err)
	}

	currentTime := time.Now()

	// if any expenses doen't have a date, use today's date as yyyy-mm-dd
	for i := range cfg.Expenses {
		if cfg.Expenses[i].Date == "" {
			// get today's date as yyyy-mm-dd


			// cfg.Expenses[i].Date = getTodaysDate()
			cfg.Expenses[i].Date = currentTime.Format("2006-01-02")
			// cfg.Expenses[i].Date = fmt.Sprintln("YYYY-MM-DD : ", currentTime.Format("2017-09-07"))
		}
	}

	reportRequest := ReportRequest{
		Type: "create",
		Credentials: ReportCredentials {
			PartnerUserID:     cfg.API.Key,
			PartnerUserSecret: cfg.API.Secret,
		},
		InputSettings: ReportInputSettings {
			Type:     "report",
			PolicyID: cfg.Report.PolicyID,
			Report:   ReportInfo {
				Title: cfg.Report.Name,
			},
			EmployeeEmail: cfg.Email,
			Expenses:      cfg.Expenses,	
		},
	}

	// marshal to json
	json, err := json.MarshalIndent(reportRequest, "", "  ")
	if err != nil {
		panic(err)
	}

	log.Println(string(json))
	postToExpensify(string(json), cfg.API.Endpoint)
}