package hangmanweb

import "fmt"

func PrintHangman(wrongGuesses int) {
	hangmanStates := []string{
		`
		 ._________.
		 |/        
		 |         
		 |         
		 |         
		 |         
	         |         
             ____|____
         
	            	`,
		`
		 ._________.
		 |/   |    
		 |         
		 |         
		 |         
		 |         
	    	 |         
	     ____|____

	`,
		`
		 ._________.
		 |/   |    
		 |    O    
		 |         
		 |         
		 |         
	    	 |         
	     ____|____

	`,
		`
		 ._________.
		 |/   |    
		 |    O    
		 |    |    
		 |   
		 |         
	    	 |         
	     ____|____

	`,
		`
		 ._________.
		 |/   |    
		 |    O    
		 |    |    
		 |    |    
		 |         
	    	 |         
	     ____|____

	`,
		`
		 ._________.
		 |/   |    
		 |    O    
		 |   /|    
		 |    |    
		 |         
	    	 |         
	     ____|____

	`,
		`
		 ._________.
		 |/   |    
		 |    O    
		 |   /|\   
		 |    |    
		 |         
	    	 |         
	     ____|____

	`,
		`
		 ._________.
		 |/   |    
		 |    O    
		 |   /|\   
		 |    |    
		 |   /     
	    	 |         
    	     ____|____

	`,
		`
		 ._________.
		 |/   |    
		 |    O    
		 |   /|\   
		 |    |    
		 |   / \   
	    	 |         
	     ____|____

	`,
		`
		 ._________.
		 |/   |    
		 |    O    
		 |   /|\   
		 |  ° |    
		 |   / \   
	         |        
	     ____|____

	`,
		`
		 ._________.
		 |/   |    
		 |    O    
		 |   /|\   
		 |  ° | °   
		 |   / \   
	    	 |       
	     ____|____

	`,
	}

	if wrongGuesses < 0 {
		wrongGuesses = 0
	} else if wrongGuesses >= len(hangmanStates) {
		wrongGuesses = len(hangmanStates) - 1
	}

	fmt.Println(hangmanStates[wrongGuesses])
}
