package cmd

import (
	"fmt"

	"github.com/prae014/httpcli/pkg"
	"github.com/spf13/cobra"
)

// GET
// var head_flags []string
// var query_flags []string
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
		User can specify the key-value pair they want.
		For example, "httpcli post example.com --header key1=val1 --header key2=value2"	
	-j, --json
		Construct JSON body of a request.
		This command also validates the JSON input
		For example, "httpcli post exmaple.com --json "{'key': 'value'}"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		//for i := 0; i < len(args); i++ {
		//	fmt.Println(args[i])
		//}
		//TODO: make http print function (formatting) as a global function so we can print everthing by just calling it
		//FIXME: header value still has [], need to get rid of itA

		proto, status_code, header, body := pkg.Post(args[0], query_flags)

		//if there is no header flag, we return all headers and other data. Otherwise return the specified flags
		if len(head_flags) == 0 {
			fmt.Printf("%v %v\n\n", proto, status_code)
			for key, val := range header {
				fmt.Printf("%v: %v\n", key, val)
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

	//GET

	//getCmd.Flags().StringSliceVarP(&head_flags, "header", "H", []string{}, "return specified header")
	postCmd.Flags().StringVarP(&json_flags, "json", "j", "", "construct json body of the POST request")
	rootCmd.AddCommand(postCmd)
}

// modify URL if queries exist
//func URLmod(req *http.Request, q_flags []string) (new_url string) {
//	query := req.URL.Query()
//	for _, flag := range q_flags {
//		//split flags to ["key", "val"}, using = as a delimiter
//		split_str := strings.Split(flag, "=")
//		query.Add(split_str[0], split_str[1])
//	}
//	req.URL.RawQuery = query.Encode() // RawQuery (from URL struct) encoded query values, without '?'
//
//	return req.URL.String()
//}