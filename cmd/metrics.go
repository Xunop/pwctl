package cmd

import (
	"example/pwctl/token"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

type metric struct {
	id          int    `json:"m_id"`
	sql         string `json:"m_sql"`
	name        string `json:"m_name"`
	sqlSu       string `json:"m_sql_su"`
	comment     string `json:"m_comment"`
	isActive    bool   `json:"m_is_active"`
	isHelper    bool   `json:"m_is_helper"`
	masterOnly  bool   `json:"m_master_only"`
	columnAttrs string `json:"m_column_attrs"`
	standbyOnly bool   `json:"m_standby_only"`
	pgVersion   int    `json:"m_pg_version_from"`
	lastMod     string `json:"m_last_modified_on"`
}

func init() {
	rootCmd.AddCommand(getMetricsCmd)

	deleteMetricsCmd.Flags().StringP("id", "i", "", "id of the metric to delete")
	rootCmd.AddCommand(deleteMetricsCmd)
}

var getMetricsCmd = &cobra.Command{
	Use:   "get-metrics",
	Short: "Get metrics",
	Long:  `Get metrics from pwatch3.`,
	Run:   getMetrics,
}

var deleteMetricsCmd = &cobra.Command{
	Use:   "delete-metrics",
	Short: "Delete metrics by id",
	Long:  `Delete metrics by id from pwatch3.`,
	Run:   deleteMetrics,
}

func deleteMetrics(cmd *cobra.Command, args []string) {
	method := "DELETE"
	client := &http.Client{}

	id, err := cmd.Flags().GetString("id")
	if err != nil {
		fmt.Println(err)
		return
	}
	params := url.Values{}
	params.Add("id", id)

	reqUrl := "http://" + PwatchURL + "/metric" + "?" + params.Encode()
	req, err := http.NewRequest(method, reqUrl, nil)
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
		fmt.Println("Delete metrics failed. Status code: ", res.StatusCode)
		return
	}
	fmt.Println("Metrics deleted.")
}

func getMetrics(cmd *cobra.Command, args []string) {
	reqUrl := "http://" + PwatchURL + "/metric"
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequest(method, reqUrl, nil)

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
		fmt.Println("Get metrics failed. Status code: ", res.StatusCode)
		return
	}
	content, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	metrics := []metric{}
	gjson.Parse(string(content)).ForEach(func(key, value gjson.Result) bool {
		m := metric{
			id:          int(value.Get("m_id").Int()),
			sql:         value.Get("m_sql").String(),
			name:        value.Get("m_name").String(),
			sqlSu:       value.Get("m_sql_su").String(),
			comment:     value.Get("m_comment").String(),
			isActive:    value.Get("m_is_active").Bool(),
			isHelper:    value.Get("m_is_helper").Bool(),
			masterOnly:  value.Get("m_master_only").Bool(),
			columnAttrs: value.Get("m_column_attrs").String(),
			standbyOnly: value.Get("m_standby_only").Bool(),
			pgVersion:   int(value.Get("m_pg_version_from").Int()),
			lastMod:     value.Get("m_last_modified_on").String(),
		}
		metrics = append(metrics, m)
		return true
	})

	// Sort the metrics by id
	sort.Slice(metrics, func(i, j int) bool {
		return metrics[i].id < metrics[j].id
	})

	// Use tablewriter to print the metrics in a table
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	fmt.Fprintln(w, "id\tname")
	for _, m := range metrics {
		fmt.Fprintf(w, "%d\t%s\n", m.id, m.name)
	}
	w.Flush()
}
