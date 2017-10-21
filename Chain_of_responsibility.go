package main

import "fmt"

type Handler interface{
handleRequest(request string)
SetSuccessor(next Handler)
}

type ConcreteHandler1 struct{
	Successor Handler
}

func (concrete1 ConcreteHandler1) handleRequest(request string) {
 	fmt.Printf("R1 R1 got the request...") 
	if ( request == "R1" ){
            fmt.Printf(" => This one is mine!")
        }else{
	     fmt.Printf(" => This one is not mine!")
            	if ( concrete1.Successor != nil ){
			fmt.Printf("Test!")
			concrete1.Successor.handleRequest(request);
		}	
                
        }
}

func (concrete1 *ConcreteHandler1) SetSuccessor(next Handler) {
        concrete1.Successor = next 
	}

type ConcreteHandler2 struct{
	Successor Handler
}

func (concrete2 ConcreteHandler2) handleRequest(request string) {
 	fmt.Printf("R2 R2 got the request...") 
	if ( request == "R2" ){
            fmt.Printf(" => This one is mine!")
        }else{
	     fmt.Printf(" => This one is not mine!")
            	if ( concrete2.Successor != nil ){
			fmt.Printf("Test!")
			concrete2.Successor.handleRequest(request);
		}	
                
        }
}

func (concrete2 *ConcreteHandler2) SetSuccessor(next Handler) {
        concrete2.Successor = next 
	}


type ConcreteHandler3 struct{
	Successor Handler
}

func (concrete3 ConcreteHandler3) handleRequest(request string) {
 	fmt.Printf("R3 R3 got the request...") 
	if ( request == "R3" ){
            fmt.Printf(" => This one is mine!")
        }else{
	     fmt.Printf(" => This one is not mine!")
            	if ( concrete3.Successor != nil ){
			fmt.Printf("Test!")
			concrete3.Successor.handleRequest(request);
		}	
                
        }
}

func (concrete3 *ConcreteHandler3) SetSuccessor(next Handler) {
        concrete3.Successor = next 
	}


func runTest(){
	h1 := new(ConcreteHandler1)
	h2 := new(ConcreteHandler2)
	h3 := new(ConcreteHandler3)
	
	h1.SetSuccessor(h2)
	h2.SetSuccessor(h3)
	
	fmt.Printf("Sending R2...")
	h1.handleRequest("R2")
	fmt.Printf("Sending R3...")
	h1.handleRequest("R3")
	fmt.Printf("Sending R1...")
	h1.handleRequest("R1")
	fmt.Printf("Sending RX...")
	h1.handleRequest("RX")
}

func main(){
	runTest()
}



    
   
