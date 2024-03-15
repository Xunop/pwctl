package cmd

import (
	"example/pwctl/token"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	loginCmd.Flags().StringP("user", "u", "", "Username")
	loginCmd.Flags().StringP("password", "p", "", "Password")
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to pwatch3",
	Long:  `Login to pwatch3 and get a token to use for authentication.`,
	Run:   login,
}

func login(cmd *cobra.Command, args []string) {
	fmt.Println("Login to pwatch3")
	baseURL := "http://" + PwatchURL
	fmt.Println("Base URL: ", baseURL)
	method := "POST"
	username, _ := cmd.Flags().GetString("user")
	password, _ := cmd.Flags().GetString("password")
	payload := strings.NewReader(`{
        "user": "` + username + `",
        "password": "` + password + `"
    }`)

	client := &http.Client{}
	url := baseURL + "/login"
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res.StatusCode != 200 {
		fmt.Println("Login failed. Status code: ", res.StatusCode)
		fmt.Println(string(body))
		return
	}
	if err := token.StoreToken(string(body)); err != nil {
		fmt.Println(err)
		return
	}
}
