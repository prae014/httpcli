package cmd

import (
	"fmt"

	"github.com/prae014/httpcli/pkg"
	"github.com/spf13/cobra"
)

var json_flags string
var postCmd = &cobra.Command{
	Use:   "post <URL>",
	Short: "post sends a POST request to a given URL",
	Long: `post sends a POST request to a given URL 

	Flags:
	-h, --help			
		help for post	
	-q, --query			
		return query parameters	
		User can specify the key they want.
		For example, "httpcli post example.com --query key1=val1 --query key2=value2"
	-H, --header 	
		return headers
		User can specify the key they want.
		For example, "httpcli post example.com --header key1=val1 --header key2=value2"	
	-j, --json
		Construct JSON body of a request.
		This command also validates the JSON input
		For example, "httpcli post example.com --json '{"key": "value"}'
	`,
	Run: func(cmd *cobra.Command, args []string) {

		proto, status_code, header, body := pkg.Post(args[0], query_flags, json_flags)

		//if there is no header flag, we return all headers and other data. Otherwise return the specified flags
		//fmt.Printf("args[1] is: %v\n", json_flags)
		if len(head_flags) == 0 {
			fmt.Printf("%v %v\n\n", proto, status_code)
			for key, val := range header {
				fmt.Printf("%v: ", key)
				for _, each_val := range val {
					fmt.Printf("%v", each_val)
				}
				fmt.Printf("\n")
			}
			fmt.Printf("\n%v\n", string(body))
		} else {
			//flags (testing)
			for i, head := range head_flags {
				//head flags contain all specified flags requested
				requested_head, ok := header[head]
				if !ok {
					fmt.Printf("%v: This header does not exist\n", head_flags[i])
				} else {
					fmt.Printf("%v: %v\n", head_flags[i], requested_head)
				}
			}
		}
	},
}

func init() {

	//POST
	postCmd.Flags().StringVarP(&json_flags, "json", "j", "", "construct json body of the POST request")
	rootCmd.AddCommand(postCmd)
}
