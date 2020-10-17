package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func check(k *string, l int){
	i := 0
	for true{
		if len(*k) < l {
			key := *k
			*k += string(key[i])
			i ++
		}else{
			break
		}
	}
}

func encrypt(text,k string) (string, string){ //here text and k are formal parameter
	//The parameters used in the header of function definition are called formal parameters of the function.
	check( &k, len(text))
	cipher := ""
	c := ""
	for j,i := range(text){
		if i == 32{
			c = " "
		}else{
			if i > 96{
				c = string(((int(i) + int(k[j]) - 194) % 26) + 97)
			}else{
				c = string(((int(i) - 65) + (int(k[j]) - 97) % 26) + 65)
			}
		}
		cipher += c
	}
	return cipher, k
}

func decrypt(cipher, k string) string{
	check(&k, len(cipher))
	plain := ""
	c := ""
	for j,i := range(cipher){
		if i == 32{
			c = " "
		}else{
			if i > 96{
				c = string(((((int(i) - 97) - (int(k[j]) - 97)) + 26) % 26) + 97)
			}else{
				c = string((((int(i) - 65) - (int(k[j]) - 97) + 26)% 26) + 65)
			}
		}
		plain += c
	}
	return plain
}

func main(){
	fmt.Print("Enter the string :")
	var text,key,text1,key1 string
	s := bufio.NewReader(os.Stdin)
 	text, _ = s.ReadString('\n')
 	text = strings.ReplaceAll(text, "\n", "")
 	fmt.Print("Enter key :")
	fmt.Scan(&key)
    k := strings.ToLower(key)
    cipher, k := encrypt(text, k) //here text and k are actual parameter
    //The parameters used in the function call are called actual parameters.
    fmt.Println("Cipher :" ,cipher)
    fmt.Print("Enter Cipher :")
    text1, _ = s.ReadString('\n')
 	text1 = strings.ReplaceAll(text1, "\n", "")
 	fmt.Print("Enter key : ")
 	fmt.Scan(&key1)
    plain := decrypt(text1, key1)
    fmt.Println("Message :", plain)
}