//LookupAddr performs a reverse lookup for the given address, returning a list of names mapping to that address.
//Here is a go lang example that shows how to get a host name for a given ip address.

package main
 
import (
  "fmt"
  "net"
)
 
func main() {

    // ip address

   host, err := net.LookupAddr("")
   if err == nil {
      fmt.Println(host)
   }
}