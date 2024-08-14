rule abcRule {
	meta: 
		author = "Wannes Vantorre"
	strings:
		$str = "abc"
	condition:
		$str
}
rule xyzRule {
	meta: 
		author = "Wannes Vantorre"
	strings:
		$str = "xyz"
	condition:
		$str
}