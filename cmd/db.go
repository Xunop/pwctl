package cmd

import (
	"bytes"
	"encoding/json"
	"example/pwctl/token"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

/*
*
  - {
    "md_name": "test1",
    "md_group": "default",
    "md_config": null,
    "md_dbtype": "postgres",
    "md_connstr": "postgresql://pgwatch3:pgwatch3admin@postgres/pgwatch3",
    "md_encryption": "plain-text",
    "md_is_enabled": true,
    "md_custom_tags": null,
    "md_host_config": null,
    "md_is_superuser": false,
    "md_config_standby": null,
    "md_only_if_master": false,
    "md_exclude_pattern": null,
    "md_include_pattern": null,
    "md_last_modified_on": "2024-03-09T09:10:05.779019+00:00",
    "md_preset_config_name": "exhaustive",
    "md_preset_config_name_standby": null
    }

*
*/
type db struct {
	name       string `json:"md_name"`
	connString string `json:"md_connstr"`
}

func init() {
	rootCmd.AddCommand(getMonitoredDbCmd)

	addMonitoredDbCmd.Flags().StringP("name", "n", "", "Database name")
	addMonitoredDbCmd.Flags().StringP("connstr", "d", "", "Database connection string")
	addMonitoredDbCmd.Flags().StringP("config", "c", "", "Database config")
	addMonitoredDbCmd.Flags().StringP("preset", "p", "exhaustive", "Preset config name")
	addMonitoredDbCmd.Flags().BoolP("superuser", "u", false, "Is superuser")
	addMonitoredDbCmd.Flags().BoolP("enable", "", false, "Is enabled")
	rootCmd.AddCommand(addMonitoredDbCmd)

	delMonitoredDbCmd.Flags().StringP("name", "n", "", "Database name")
	rootCmd.AddCommand(delMonitoredDbCmd)

	enableMonitoredDbCmd.Flags().StringP("name", "n", "", "Database name")
	enableMonitoredDbCmd.Flags().BoolP("enable", "", false, "Is enabled")
	rootCmd.AddCommand(enableMonitoredDbCmd)
}

var getMonitoredDbCmd = &cobra.Command{
	Use:   "get-monitored-db",
	Short: "Get monitored databases",
	Long:  `Get monitored databases from pwatch3.`,
	Run:   getMonitoredDb,
}

var addMonitoredDbCmd = &cobra.Command{
	Use:   "add-monitored-db",
	Short: "Add monitored database",
	Long:  `Add monitored database to pwatch3.`,
	Run:   addMonitoredDb,
}

var delMonitoredDbCmd = &cobra.Command{
	Use:   "del-monitored-db",
	Short: "Delete monitored database",
	Long:  `Delete monitored database from pwatch3.`,
	Run:   delMonitoredDb,
}

var enableMonitoredDbCmd = &cobra.Command{
	Use:   "enable-monitored-db",
	Short: "Enable or disable monitored database",
	Long:  `Enable or disable monitored database in pwatch3.`,
	Run:   enableMonitoredDb,
}

func getMonitoredDb(cmd *cobra.Command, args []string) {
	dbURL := "http://" + PwatchURL + "/db"
	method := "GET"
	client := &http.Client{}
	// println(dbURL)
	// println(token.GetToken())
	req, err := http.NewRequest(method, dbURL, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", token.GetToken())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("Get monitored databases failed. Status code: ", res.StatusCode)
		return
	}
	content, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("Get monitored databases success")
	dbs := []db{}
	gjson.Parse(string(content)).ForEach(func(key, value gjson.Result) bool {
		// fmt.Println("DB Name: ", value.Get("md_name").String())
		// fmt.Println("DB Connection String: ", value.Get("md_connstr").String())
		dbs = append(dbs, db{
			name:       value.Get("md_name").String(),
			connString: value.Get("md_connstr").String(),
		})
		return true
	})

	// Use tabwriter
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	fmt.Fprintln(w, "DB Name\tConnection String")
	for _, db := range dbs {
		fmt.Fprintf(w, "%s\t%s\n", db.name, db.connString)
	}

	w.Flush()
}

func addMonitoredDb(cmd *cobra.Command, args []string) {
	dbURL := "http://" + PwatchURL + "/db"
	method := "POST"
	client := &http.Client{}

	mdName, _ := cmd.Flags().GetString("name")
	mdConnStr, _ := cmd.Flags().GetString("connstr")
	mdConfig, _ := cmd.Flags().GetString("config")
	mdPresetConfigName, _ := cmd.Flags().GetString("preset")
	mdIsSuperuser, _ := cmd.Flags().GetBool("superuser")
	mdIsEnabled, _ := cmd.Flags().GetBool("enable")
	if !isValidPresetConfigName(mdPresetConfigName) {
		fmt.Println("Invalid preset config name")
		return
	}
	data := map[string]interface{}{
		"md_name":               mdName,
		"md_connstr":            mdConnStr,
		"md_config":             nilIfEmpty(mdConfig),
		"md_preset_config_name": mdPresetConfigName,
		"md_is_superuser":       mdIsSuperuser,
		"md_is_enabled":         mdIsEnabled,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest(method, dbURL, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", token.GetToken())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("Add monitored database failed. Status code: ", res.StatusCode)
		// Print the response body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))
		return
	}

	fmt.Println("Add monitored database success")
}

func delMonitoredDb(cmd *cobra.Command, args []string) {
	method := "DELETE"

	client := &http.Client{}
	params := url.Values{}
	dbName, err := cmd.Flags().GetString("name")
	if err != nil {
		fmt.Println(err)
		return
	}
	params.Add("id", dbName)
	dbURL := "http://" + PwatchURL + "/db" + "?" + params.Encode()
	req, err := http.NewRequest(method, dbURL, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", token.GetToken())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("Delete monitored database failed. Status code: ", res.StatusCode)
		return
	}
	fmt.Println("Delete monitored database success")
}

func enableMonitoredDb(cmd *cobra.Command, args []string) {
	dbUrl := "http://" + PwatchURL + "/db"
	method := "PATCH"
	client := &http.Client{}
	params := url.Values{}
	dbName, err := cmd.Flags().GetString("name")
	if err != nil {
		fmt.Println(err)
		return
	}
	params.Add("id", dbName)
	dbURL := dbUrl + "?" + params.Encode()
	isEnabled, _ := cmd.Flags().GetBool("enable")
	data := map[string]interface{}{
		"md_is_enabled": isEnabled,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest(method, dbURL, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", token.GetToken())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("Enable monitored database failed. Status code: ", res.StatusCode)
		// Print the response body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))
		return
	}
	fmt.Println("Enable monitored database success")
}

// Pause all monitored databases
func pauseMonitoredDb(cmd *cobra.Command, args []string) {
    
}

func nilIfEmpty(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}

var validPresetConfigNames = map[string]bool{
	"minimal":             true,
	"basic":               true,
	"standard":            true,
	"pgbouncer":           true,
	"pgpool":              true,
	"exhaustive":          true,
	"full":                true,
	"full_influx":         true,
	"unprivileged":        true,
	"aurora":              true,
	"azure":               true,
	"rds":                 true,
	"gce":                 true,
	"prometheus":          true,
	"prometheus-async":    true,
	"superuser_no_python": true,
}

func isValidPresetConfigName(name string) bool {
	_, ok := validPresetConfigNames[name]
	return ok
}
