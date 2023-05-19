package cmd

import (
	"fmt"

	"github.com/prae014/httpcli/pkg"
	"github.com/spf13/cobra"
)

// GET
var deleteCmd = &cobra.Command{
	Use:   "delete <URL>",
	Short: "delete sends a DELETE request to a given URL",
	Long: `delete sends a DELETE request to a given URL 

	Flags:
	-h, --help			
		help for get	
	-q, --query			
		return query parameters	
		User can specify the key they want.
		For example, "httpcli delete example.com --query key1=val1 --query key2=value2"
	-H, --header 	
		return headers
		User can specify the key they want.
		For example, "httpcli delete example.com --header key1=val1 --header key2=value2"	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		//for i := 0; i < len(args); i++ {
		//	fmt.Println(args[i])
		//}
		//TODO: make http print function (formatting) as a global function so we can print everthing by just calling it
		//FIXME: header value still has [], need to get rid of itA

		proto, status_code, header, body := pkg.Delete(args[0], query_flags)

		//if there is no header flag, we return all headers and other data. Otherwise return the specified flags
		if len(head_flags) == 0 {
			fmt.Printf("%v %v\n\n", proto, status_code)
			for key, val := range header {
				//fmt.Printf("%v: %v\n", key, val)

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

	//DELETE

	//getCmd.Flags().StringSliceVarP(&head_flags, "header", "H", []string{}, "return specified header")
	rootCmd.AddCommand(deleteCmd)
}
