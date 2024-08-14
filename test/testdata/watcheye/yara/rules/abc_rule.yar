rule abcRule {
		meta: 
			author = "Wannes Vantorre"
		strings:
			$str = "abc"
		condition:
			$str
	}